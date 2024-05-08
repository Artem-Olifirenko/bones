package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	citizap_factory "go.citilink.cloud/citizap/factory"
	"go.citilink.cloud/grpc-skeleton/internal/api/http"
)

func NewServer(loggerFactory citizap_factory.Factory) *Server {
	return &Server{
		loggerFactory: loggerFactory,
		httpServer:    http.NewFastHttpServer(),
	}
}

type Server struct {
	loggerFactory citizap_factory.Factory
	httpServer    *http.FastHttpServer
}

func (s *Server) ListenAndServe(ctx context.Context, listen string) <-chan error {
	logger := s.loggerFactory.Create(ctx)
	router := fasthttprouter.New()
	// обязательный путь для проверки валидной работы самого HTTP-сервера
	router.GET("/ping", func(ctx *fasthttp.RequestCtx) {
		ctx.SuccessString("text/plain", "PONG")
	})

	router.GET("/api/hello/", s.HandleHello)
	router.GET("/api/hello/:name", s.HandleHello)
	router.GET("/api/hello-json/", s.HandleJSONHello)
	router.GET("/api/hello-json/:name", s.HandleJSONHello)

	return s.httpServer.Run(ctx, router.Handler, listen, logger)
}

// HandleHello пример обработчика, всегда возвращает html с кодом состояния 200
func (s *Server) HandleHello(ctx *fasthttp.RequestCtx) {
	logger := s.loggerFactory.Create(ctx)
	name := ctx.UserValue("name")
	if name == nil {
		name = "World"
	}
	nameStr, _ := name.(string)

	resp := fmt.Sprintf("Hello, %s!\n", nameStr)
	logger.Info("Action handleHello was called")

	ctx.SetBody([]byte(resp))
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// HandleJSONHello пример обработчика HelloWorldController, всегда возращает json с кодом состояния 200
func (s *Server) HandleJSONHello(ctx *fasthttp.RequestCtx) {
	name := ctx.UserValue("name")
	if name == nil {
		name = "World"
	}
	nameStr, _ := name.(string)

	encoder := json.NewEncoder(ctx)
	err := encoder.Encode(&HelloWorldJSONResponse{Say: "hello", Name: nameStr})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusServiceUnavailable)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// HelloWorldJSONResponse пример структуры json ответа
type HelloWorldJSONResponse struct {
	Say  string `json:"say"`
	Name string `json:"name"`
}
