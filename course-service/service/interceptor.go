package service

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		return handler(ctx, req)
	}
}

func Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		return handler(srv, stream)
	}
}
