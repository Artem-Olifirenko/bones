package api

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	citizap_factory_ctx "go.citilink.cloud/citizap/factory/ctx"
	http2 "go.citilink.cloud/grpc-skeleton/internal/api/http"
	"go.uber.org/goleak"
	"go.uber.org/zap"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
)

var (
	logger        = zap.NewNop()
	loggerFactory = citizap_factory_ctx.New(logger)
)

func TestServer_Hello(t *testing.T) {
	s := NewServer(loggerFactory)

	ctx := &fasthttp.RequestCtx{}
	s.HandleHello(ctx)

	assert.Equal(t, []byte("text/plain; charset=utf-8"), ctx.Response.Header.ContentType())
	assert.Equal(t, []byte("Hello, World!\n"), ctx.Response.Body())
}

func TestServer_HelloName(t *testing.T) {
	s := NewServer(loggerFactory)

	ctx := &fasthttp.RequestCtx{}
	ctx.SetUserValue("name", "test")
	s.HandleHello(ctx)

	assert.Equal(t, []byte("text/plain; charset=utf-8"), ctx.Response.Header.ContentType())
	assert.Equal(t, []byte("Hello, test!\n"), ctx.Response.Body())
}

func TestServer_HelloJson(t *testing.T) {
	s := NewServer(loggerFactory)

	ctx := &fasthttp.RequestCtx{}
	s.HandleJSONHello(ctx)

	assert.Equal(t, []byte("application/json"), ctx.Response.Header.ContentType())
	assert.Equal(t, `{"say":"hello","name":"World"}`, strings.TrimSpace(string(ctx.Response.Body())))
}

func TestServer_HelloJsonName(t *testing.T) {
	s := NewServer(loggerFactory)

	ctx := &fasthttp.RequestCtx{}
	ctx.SetUserValue("name", "test")
	s.HandleJSONHello(ctx)

	assert.Equal(t, []byte("application/json"), ctx.Response.Header.ContentType())
	assert.Equal(t, `{"say":"hello","name":"test"}`, strings.TrimSpace(string(ctx.Response.Body())))
}

func TestServer_ListenAndServe(t *testing.T) {
	ln := fasthttputil.NewInmemoryListener()

	s := &Server{
		loggerFactory: loggerFactory,
		httpServer: http2.NewFastHttpServer(http2.WithServe(func(server *fasthttp.Server, _ string) error {
			return server.Serve(ln)
		})),
	}

	const listen = ":64102"

	hostname, err := os.Hostname()
	assert.NoError(t, err)

	host := fmt.Sprintf("%s%s", hostname, listen)

	ctx, cancel := context.WithCancel(context.Background())
	serverFinishErrCh := s.ListenAndServe(ctx, listen)

	routes := []string{
		"/api/hello",
		"/api/hello/name",
		"/api/hello-json",
		"/api/hello-json/name",
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
