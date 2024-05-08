package http

import (
	"context"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type Option interface {
	apply(*FastHttpServer)
}

type serveOption func(server *fasthttp.Server, addr string) error

func (o serveOption) apply(server *FastHttpServer) {
	server.serve = o
}

func WithServe(s func(server *fasthttp.Server, addr string) error) Option {
	return serveOption(s)
}

type FastHttpServer struct {
	serve func(server *fasthttp.Server, addr string) error
}

func NewFastHttpServer(opts ...Option) *FastHttpServer {
	server := &FastHttpServer{
		serve: defaultServe,
	}

	for _, o := range opts {
		o.apply(server)
	}

	return server
}

// Run запускает fasthttp сервер, и возвращает канал с ошибкой при завершении работы сервера.
func (s *FastHttpServer) Run(
	ctx context.Context,
	handler fasthttp.RequestHandler,
	listen string,
	logger *zap.Logger,
) <-chan error {
	ctx, cancel := context.WithCancel(ctx)
	server := &fasthttp.Server{
		Handler: handler,
	}

	result := make(chan error)
	go func() {
		defer close(result)
		<-ctx.Done()
		result <- server.Shutdown()
	}()

	go func() {
		err := s.serve(server, listen)
		if err != nil {
			logger.Error("server listenAndServe failed", zap.Error(err))
		}
		cancel()
	}()

	return result
}

func defaultServe(server *fasthttp.Server, addr string) error {
	return server.ListenAndServe(addr)
}
