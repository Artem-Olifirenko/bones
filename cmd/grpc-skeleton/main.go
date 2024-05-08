package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jessevdk/go-flags"
	"go.citilink.cloud/citizap"
	citizap_factory_ctx "go.citilink.cloud/citizap/factory/ctx"
	article_grpc "go.citilink.cloud/grpc-skeleton/internal/api/grpc"
	articlev1 "go.citilink.cloud/grpc-skeleton/internal/api/grpc/gen/citilink/blog/article/v1"
	"go.citilink.cloud/grpc-skeleton/internal/api/http/api"
	"go.citilink.cloud/grpc-skeleton/internal/api/http/service"
	"go.citilink.cloud/grpc-skeleton/internal/inmemory"
	"go.citilink.cloud/grpc-skeleton/internal/metrics"
	storev1 "go.citilink.cloud/grpc-skeleton/internal/specs/grpcclient/gen/citilink/store/store/v1"
	"go.citilink.cloud/organism/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	goa_grpc_middleware "goa.design/goa/v3/grpc/middleware"
	goa_middleware "goa.design/goa/v3/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

var (
	// AppName название приложения
	AppName = "{BS_REPO_NAME}"
	// Version версия приложения
	Version = "{{variable.version}}"
	// Commit хеш коммита git
	Commit = "{{variable.commit}}"
)

func main() {
	// чтобы линтер не ругался на неиспользуемые переменные.
	_ = AppName
	_ = Commit

	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(fatalJsonLog("Failed to parse config.", err))
	}

	logger, err := initLogger(cfg.LogLevel, cfg.LogJSON)
	if err != nil {
		log.Fatal(fatalJsonLog("Failed to init logger.", err))
	}
	loggerFactory := citizap_factory_ctx.New(logger, citizap_factory_ctx.WithRequestIdKey(goa_middleware.RequestIDKey))

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	configsRenewer := NewEmptyConfigsRenewer(loggerFactory)
	initSignalHandler(logger, ctx, cancelFunc, configsRenewer)

	defer func() {
		if msg := recover(); msg != nil {
			err := fmt.Errorf("%s", msg)
			logger.Error("recovered from panic, but application will be terminated", zap.Error(err))
		}
	}()

	appMetrics := metrics.New()
	// TODO: Удалите этот блок из реального проекта
	// вызов метода здесь оставлен чтобы не возникала ошибка, что переменная appMetrics не используется
	appMetrics.IncPrometheusTestCounter()
	// ----------------------------------------------------------------------------

	storage := inmemory.NewArticleStorage()

	// Пример создания grpc-клиента
	storeClient := storev1.NewStoreAPIClient(dialGrpcClient(ctx, logger, cfg.StoreEndpoint))

	orgnism := organism.NewOrganism()

	wg := sync.WaitGroup{}
	wg.Add(1)
	grpcServiceLimb := orgnism.GrowLimb("grpc_server")
	go func() {
		defer wg.Done()
		defer grpcServiceLimb.Die()
		err := startGRPCServer(ctx, cfg.GrpcListen, storage, storeClient, grpcServiceLimb, loggerFactory)
		if err != nil {
			logger.Error("can't start gRPC server or server return error while working", zap.Error(err))
		}
	}()

	wg.Add(1)
	serviceServerLimb := orgnism.GrowLimb("http_service_server")
	go func() {
		defer wg.Done()
		defer serviceServerLimb.Die()
		logger.Info("Starting private HTTP server", citizap.Listen(cfg.HttpPrivateListen))
		server := service.NewServer(orgnism, loggerFactory, configsRenewer.RenewConfig)
		finishErrChan := server.ListenAndServe(ctx, cfg.HttpPrivateListen, cfg.EnablePprof)
		serviceServerLimb.Ready()
		err := <-finishErrChan
		cancelFunc() // завершаем работу приложения, если по какой-то причине завершилась работа http сервера
		if err != nil {
			logger.Error("error on listen and serve api HTTP server", zap.Error(err))
		}
	}()

	wg.Add(1)
	apiServerLimb := orgnism.GrowLimb("http_api_server")
	go func() {
		defer wg.Done()
		defer apiServerLimb.Die()
		logger.Info("Starting public HTTP server", citizap.Listen(cfg.HttpPublicListen))
		server := api.NewServer(loggerFactory)
		if err != nil {
			logger.Error("Can't create api HTTP server", zap.Error(err))
			return
		}
		finishErrChan := server.ListenAndServe(ctx, cfg.HttpPublicListen)
		apiServerLimb.Ready()
		err := <-finishErrChan
		cancelFunc() // завершаем работу приложения, если по какой-то причине завершилась работа http сервера
		if err != nil {
			logger.Error("Error on listen and serve api HTTP server", zap.Error(err))
		}
	}()

	go func() {
		<-ctx.Done()
		orgnism.Die()
	}()
	orgnism.Ready()
	wg.Wait()
	logger.Info("Application has been shutdown gracefully")
}

func fatalJsonLog(msg string, err error) string {
	escape := func(s string) string {
		return strings.ReplaceAll(strings.ReplaceAll(s, `\`, `\\`), `"`, `\"`)
	}
	errString := ""
	if err != nil {
		errString = err.Error()
	}
	return fmt.Sprintf(
		`{"level":"fatal","ts":"%s","msg":"%s","error":"%s"}`,
		time.Now().Format(time.RFC3339),
		escape(msg),
		escape(errString),
	)
}

// startGRPCServer запускает gRPC сервер
func startGRPCServer(
	ctx context.Context,
	listen string,
	storage *inmemory.ArticleStorage,
	storeClient storev1.StoreAPIClient,
	limb *organism.Limb,
	loggerFactory citizap_factory_ctx.Factory,
) error {
	logger := loggerFactory.Create(ctx)

	logger.Info("gRPC started", citizap.Listen(listen))
	lis, err := net.Listen("tcp", listen)
	if err != nil {
		return fmt.Errorf("failed to listen GRPC server: %w", err)
	}

	grpc_prometheus.EnableClientHandlingTimeHistogram()

	recoverFromPanicHandler := func(p interface{}) error {
		err := fmt.Errorf("recovered from panic: %s", p)
		logger.Error("recovered from panic", zap.Error(err))

		return err
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoverFromPanicHandler),
	}

	grpc_prometheus.EnableHandlingTimeHistogram()
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
			grpc_zap.UnaryServerInterceptor(logger, grpc_zap.WithLevels(grpcCodeToZapLevel)),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_validator.UnaryServerInterceptor(),
			goa_grpc_middleware.UnaryRequestID(
				goa_grpc_middleware.UseXRequestIDMetadataOption(true),
				goa_grpc_middleware.XRequestMetadataLimitOption(128),
			),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
			grpc_zap.StreamServerInterceptor(logger, grpc_zap.WithLevels(grpcCodeToZapLevel)),
			grpc_prometheus.StreamServerInterceptor,
			grpc_validator.StreamServerInterceptor(),
			goa_grpc_middleware.StreamRequestID(
				goa_grpc_middleware.UseXRequestIDMetadataOption(true),
				goa_grpc_middleware.XRequestMetadataLimitOption(128),
			),
		),
	)
	articlev1.RegisterArticleAPIServer(s, article_grpc.NewArticleServer(storage, storeClient, loggerFactory))
	reflection.Register(s)

	limb.Ready()
	go func() {
		<-ctx.Done()
		s.GracefulStop()
	}()
	return s.Serve(lis)
}

// initLogger создает и настраивает новый экземпляр логгера
func initLogger(logLevel string, isLogJson bool) (*zap.Logger, error) {
	lvl := zap.InfoLevel
	err := lvl.UnmarshalText([]byte(logLevel))
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal log-level: %w", err)
	}
	opts := zap.NewProductionConfig()
	opts.Level = zap.NewAtomicLevelAt(lvl)
	opts.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if opts.InitialFields == nil {
		opts.InitialFields = map[string]interface{}{}
	}
	opts.InitialFields["version"] = Version
	if !isLogJson {
		opts.Encoding = "console"
		opts.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	return opts.Build(zap.WrapCore(citizap.WrapCore))
}

func grpcCodeToZapLevel(code codes.Code) zapcore.Level {
	switch code {
	case codes.OK:
		return zapcore.DebugLevel
	default:
		return grpc_zap.DefaultCodeToLevel(code)
	}
}

// initSignalHandler обрабатывает системные сигналы
func initSignalHandler(
	logger *zap.Logger,
	ctx context.Context,
	cancelFunc context.CancelFunc,
	configsRenewer ConfigsRenewer,
) {
	osSigCh := make(chan os.Signal, 1)

	signal.Notify(
		osSigCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		s := <-osSigCh
		switch s {
		case syscall.SIGHUP:
			logger.Info("Received signal SIGHUP! Renew configs")
			err := configsRenewer.RenewConfig(ctx)
			if err != nil {
				logger.Fatal("can't init renew configs", zap.Error(err))
			}
		case syscall.SIGINT:
			logger.Info("Received signal SIGINT! Process exited")
			cancelFunc()
		case syscall.SIGTERM:
			logger.Info("Received signal SIGTERM! Process exited")
			cancelFunc()
		case syscall.SIGQUIT:
			logger.Info("Received signal SIGQUIT! Process exited")
			cancelFunc()
		}
	}()
}

func dialGrpcClient(ctx context.Context, logger *zap.Logger, endpoint string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, endpoint, grpcClientDialOpts...)
	if err != nil {
		logger.Fatal("failed to connect to grpc client at "+endpoint, zap.Error(err))
	}

	return conn
}
