package service

import (
	"context"
	"go.citilink.cloud/citizap"
	"go.citilink.cloud/organism/v2"
	"net/http/pprof"

	"github.com/buaazp/fasthttprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	citizap_factory "go.citilink.cloud/citizap/factory"
	"go.citilink.cloud/grpc-skeleton/internal/api/http"
	"go.uber.org/zap"
)

const pprofUrlPrefix = "/debug/pprof"

func NewServer(
	organism *organism.Organism,
	loggerFactory citizap_factory.Factory,
	configsRenewer func(ctx context.Context) error,
) *Server {
	return &Server{
		loggerFactory:  loggerFactory,
		organism:       organism,
		httpServer:     http.NewFastHttpServer(),
		configsRenewer: configsRenewer,
	}
}

type Server struct {
	loggerFactory  citizap_factory.Factory
	organism       *organism.Organism
	httpServer     *http.FastHttpServer
	configsRenewer func(ctx context.Context) error
}

func (s *Server) ListenAndServe(ctx context.Context, listen string, enablePprof bool) <-chan error {
	logger := s.loggerFactory.Create(ctx)
	router := fasthttprouter.New()

	router.GET("/healthz", s.HandleLivenessProbe)
	router.GET("/healthy", s.HandleReadinessProbe)
	router.GET("/metrics", s.HandleMetrics)
	router.GET("/reload-configs", s.ReloadConfigs)
	// обязательный путь для проверки валидной работы самого HTTP-сервера
	router.GET("/ping", s.HandlePing)

	if enablePprof {
		for _, path := range []string{"/", "/allocs", "/block", "/goroutine", "/heap", "/mutex", "/threadcreate"} {
			router.GET(pprofUrlPrefix+path, fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
		}
		router.GET(pprofUrlPrefix+"/cmdline", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline))
		router.GET(pprofUrlPrefix+"/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
		router.GET(pprofUrlPrefix+"/symbol", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol))
		router.GET(pprofUrlPrefix+"/trace", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace))
		logger.Warn("pprof routes registered", citizap.Name(pprofUrlPrefix))
	}

	return s.httpServer.Run(ctx, router.Handler, listen, logger)
}

func (s *Server) HandleLivenessProbe(ctx *fasthttp.RequestCtx) {
	if s.organism.IsAlive() {
		ctx.SuccessString("text/plain; charset=utf-8", "OK")
		return
	}
	ctx.Error("Service Unavailable", fasthttp.StatusServiceUnavailable)

	s.loggerFactory.Create(ctx).
		Error("some limbs of the organism died", zap.Strings("limbs", limbNames(s.organism.DeadLimbs())))
}

func (s *Server) HandleReadinessProbe(ctx *fasthttp.RequestCtx) {
	if s.organism.IsReady() {
		ctx.SuccessString("text/plain; charset=utf-8", "OK")
		return
	}
	ctx.Error("Service Unavailable", fasthttp.StatusServiceUnavailable)

	s.loggerFactory.Create(ctx).
		Info("these limbs are not ready yet", zap.Strings("limbs", limbNames(s.organism.NotReadyLimbs())))
}

func (s *Server) HandleMetrics(ctx *fasthttp.RequestCtx) {
	prometheusHandler := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	prometheusHandler(ctx)
}

func (s *Server) HandlePing(ctx *fasthttp.RequestCtx) {
	ctx.SuccessString("text/plain; charset=utf-8", "PONG")
}

func (s *Server) ReloadConfigs(ctx *fasthttp.RequestCtx) {
	logger := s.loggerFactory.Create(ctx)

	err := s.configsRenewer(ctx)
	if err != nil {
		logger.Error("can't init renew configs", zap.Error(err))
		ctx.Error("can't init renew configs", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SuccessString("text/plain; charset=utf-8", "OK")
}

func limbNames(limbs []*organism.Limb) []string {
	names := make([]string, 0, len(limbs))
	for _, limb := range limbs {
		names = append(names, limb.Name())
	}
	return names
}
