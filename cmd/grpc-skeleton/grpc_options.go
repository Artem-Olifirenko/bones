package main

import (
	"context"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"time"
)

var (
	grpcClientDialOpts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{Time: 6 * time.Minute, Timeout: 1 * time.Second}),
		grpc.WithChainUnaryInterceptor(
			grpc_prometheus.UnaryClientInterceptor,
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(opentracing.GlobalTracer()),
			),
			UnaryClientOutContextInterceptor,
		),
		grpc.WithChainStreamInterceptor(
			grpc_prometheus.StreamClientInterceptor,
			grpc_opentracing.StreamClientInterceptor(
				grpc_opentracing.WithTracer(opentracing.GlobalTracer()),
			),
			StreamClientOutContextInterceptor,
		),
	}

	UnaryClientOutContextInterceptor = func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		headersIn, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		return invoker(metadata.NewOutgoingContext(ctx, headersIn), method, req, reply, cc, opts...)
	}

	StreamClientOutContextInterceptor = func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		headersIn, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return streamer(ctx, desc, cc, method, opts...)
		}

		return streamer(metadata.NewOutgoingContext(ctx, headersIn), desc, cc, method, opts...)
	}
)
