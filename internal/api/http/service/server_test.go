package service

import (
	"context"
	"errors"
	"fmt"
	organism "go.citilink.cloud/organism/v2"
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"go.citilink.cloud/citizap/factory"
	citizap_factory_ctx "go.citilink.cloud/citizap/factory/ctx"
	http2 "go.citilink.cloud/grpc-skeleton/internal/api/http"
	"go.uber.org/goleak"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type ServerHandlersSuite struct {
	suite.Suite
	ctx           *fasthttp.RequestCtx
	organism      *organism.Organism
	logs          *observer.ObservedLogs
	loggerFactory factory.Factory
	configRenewer func(ctx context.Context) error
	server        *Server
}

func (s *ServerHandlersSuite) SetupTest() {
	s.ctx = &fasthttp.RequestCtx{}
	s.organism = organism.NewOrganism()
	var core zapcore.Core
	core, s.logs = observer.New(zapcore.InfoLevel)
	s.loggerFactory = citizap_factory_ctx.New(zap.New(core))
	s.configRenewer = func(ctx context.Context) error { return nil }
	s.server = NewServer(s.organism, s.loggerFactory, s.configRenewer)
}

func (s *ServerHandlersSuite) Test_LivenessProbe_Ok() {
	s.server.HandleLivenessProbe(s.ctx)

	s.Equal(fasthttp.StatusOK, s.ctx.Response.StatusCode())
	s.Equal([]byte("text/plain; charset=utf-8"), s.ctx.Response.Header.ContentType())
	s.Equal([]byte("OK"), s.ctx.Response.Body())
	s.Zero(s.logs.Len())
}

func (s *ServerHandlersSuite) Test_LivenessProbe_NotOk() {
	s.organism.GrowLimb("limb1").Die()
	s.organism.Die()

	s.server.HandleLivenessProbe(s.ctx)

	s.Equal(fasthttp.StatusServiceUnavailable, s.ctx.Response.StatusCode())
	s.Equal([]byte("text/plain; charset=utf-8"), s.ctx.Response.Header.ContentType())
	s.Equal([]byte("Service Unavailable"), s.ctx.Response.Body())
	s.Equal(1,
		s.logs.
			FilterLevelExact(zapcore.ErrorLevel).
			FilterMessage("some limbs of the organism died").
			FilterField(zap.Strings("limbs", []string{"core", "limb1"})).
			Len(),
	)
	s.Equal(1, s.logs.Len())
}

func (s *ServerHandlersSuite) Test_ReadinessProbe_Ok() {
	s.organism.Ready()
	s.server.HandleReadinessProbe(s.ctx)

	s.Equal(fasthttp.StatusOK, s.ctx.Response.StatusCode())
	s.Equal([]byte("text/plain; charset=utf-8"), s.ctx.Response.Header.ContentType())
	s.Equal([]byte("OK"), s.ctx.Response.Body())
	s.Zero(s.logs.Len())
}

func (s *ServerHandlersSuite) Test_ReadinessProbe_NotOk() {
	s.organism.GrowLimb("limb1")

	s.server.HandleReadinessProbe(s.ctx)

	s.Equal(fasthttp.StatusServiceUnavailable, s.ctx.Response.StatusCode())
	s.Equal([]byte("text/plain; charset=utf-8"), s.ctx.Response.Header.ContentType())
	s.Equal([]byte("Service Unavailable"), s.ctx.Response.Body())
	s.Equal(1,
		s.logs.
			FilterLevelExact(zapcore.InfoLevel).
			FilterMessage("these limbs are not ready yet").
			FilterField(zap.Strings("limbs", []string{"core", "limb1"})).
			Len(),
	)
	s.Equal(1, s.logs.Len())
}

func (s *ServerHandlersSuite) Test_Ping_Ok() {
	s.organism.Ready()
	s.server.HandlePing(s.ctx)

	s.Equal(fasthttp.StatusOK, s.ctx.Response.StatusCode())
	s.Equal([]byte("text/plain; charset=utf-8"), s.ctx.Response.Header.ContentType())
	s.Equal([]byte("PONG"), s.ctx.Response.Body())
	s.Zero(s.logs.Len())
}

func TestServer_Handle(t *testing.T) {
	suite.Run(t, &ServerHandlersSuite{})
}

var (
	logger        = zap.NewNop()
	loggerFactory = citizap_factory_ctx.New(logger)
)

func TestServer_ListenAndServe_pprof(t *testing.T) {
	tests := []struct {
		name         string
		pprofEnabled bool
		status       int
	}{
		{
			name:         "pprof enabled",
			pprofEnabled: true,
			status:       http.StatusOK,
		},
		{
			name:         "pprof disabled",
			pprofEnabled: false,
			status:       http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())

			ln := fasthttputil.NewInmemoryListener()

			org := organism.NewOrganism()
			org.Ready()
			s := &Server{
				loggerFactory: loggerFactory,
				organism:      org,
				httpServer: http2.NewFastHttpServer(http2.WithServe(func(server *fasthttp.Server, _ string) error {
					return server.Serve(ln)
				})),
			}

			const listen = ":64102"

			hostname, err := os.Hostname()
			assert.NoError(t, err)

			host := fmt.Sprintf("%s%s", hostname, listen)
			serverFinishErrCh := s.ListenAndServe(ctx, listen, tt.pprofEnabled)

			client := http.Client{
				Transport: &http.Transport{
					DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
						return ln.Dial()
					},
				},
			}

			resp, err := client.Get(fmt.Sprintf("http://%s%s", host, pprofUrlPrefix))

			assert.NoError(t, err)

			assert.Equal(t, tt.status, resp.StatusCode)
			assert.NoError(t, resp.Body.Close())
			cancel() // завершаем работу сервера
			err = <-serverFinishErrCh
			assert.NoError(t, err)
			// в workerPool.Start у fastHttp остается подвисшая горутина, игнорируем ее.
			goleak.VerifyNone(t, goleak.IgnoreTopFunction("time.Sleep"))
		})
	}
}

func TestServer_ListenAndServe_ServiceRoutes(t *testing.T) {
	ln := fasthttputil.NewInmemoryListener()

	org := organism.NewOrganism()
	org.Ready()
	s := &Server{
		loggerFactory: loggerFactory,
		organism:      org,
		httpServer: http2.NewFastHttpServer(http2.WithServe(func(server *fasthttp.Server, _ string) error {
			return server.Serve(ln)
		})),
		configsRenewer: func(ctx context.Context) error {
			return nil
		},
	}

	const listen = ":64102"

	hostname, err := os.Hostname()
	assert.NoError(t, err)

	host := fmt.Sprintf("%s%s", hostname, listen)

	ctx, cancel := context.WithCancel(context.Background())
	serverFinishErrCh := s.ListenAndServe(ctx, listen, false)

	routes := []string{
		"/healthz",
		"/healthy",
		"/metrics",
		"/ping",
		"/reload-configs",
	}

	for _, r := range routes {
		r := r
		t.Run(fmt.Sprintf("route %s", r), func(t *testing.T) {
			client := http.Client{
				Transport: &http.Transport{
					DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
						return ln.Dial()
					},
				},
			}
			resp, err := client.Get(fmt.Sprintf("http://%s%s", host, r))

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.NoError(t, resp.Body.Close())
		})
	}
	cancel() // завершаем работу сервера
	err = <-serverFinishErrCh
	assert.NoError(t, err)
	// в workerPool.Start у fastHttp остается подвисшая горутина, игнорируем ее.
	goleak.VerifyNone(t, goleak.IgnoreTopFunction("time.Sleep"))
}

func TestServer_ReloadConfigs(t *testing.T) {
	org := organism.NewOrganism()
	ctx := &fasthttp.RequestCtx{}

	type args struct {
		callback func(ctx context.Context) error
	}

	tests := []struct {
		name string
		args *args
		want func() []byte
	}{
		{
			name: "success",
			args: &args{
				callback: func(ctx context.Context) error {
					return nil
				},
			},
			want: func() []byte {
				return []byte("OK")
			},
		},
		{
			name: "error",
			args: &args{
				callback: func(ctx context.Context) error {
					return errors.New("mock")
				},
			},
			want: func() []byte {
				return []byte("can't init renew configs")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			wantBody := tt.want()
			s := NewServer(org, loggerFactory, tt.args.callback)
			s.ReloadConfigs(ctx)

			assert.Equal(t, []byte("text/plain; charset=utf-8"), ctx.Response.Header.ContentType())
			assert.Equal(t, wantBody, ctx.Response.Body())
		})
	}
}
