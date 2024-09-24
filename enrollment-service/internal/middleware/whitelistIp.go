package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"net"
)

func IpWhiteListInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Unable to get Client Ip")
	}
	addr := p.Addr.String()
	ip, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to parse client IP: %v", err)
	}

	// only allow access ip localhost:8080
	if ip != "127.0.0.1" && ip != "::1" {
		return nil, status.Errorf(codes.PermissionDenied, "Your IP is not allowed")
	}
	return handler(ctx, req)
}
