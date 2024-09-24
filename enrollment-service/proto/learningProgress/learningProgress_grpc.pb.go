// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: learningProgress/learningProgress.proto

package learningProgress

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LearningProgressService_CreateLearningProgress_FullMethodName = "/learningProgress.LearningProgressService/CreateLearningProgress"
)

// LearningProgressServiceClient is the client API for LearningProgressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// endpoints
type LearningProgressServiceClient interface {
	CreateLearningProgress(ctx context.Context, in *CreateLearningProgressRequest, opts ...grpc.CallOption) (*CreateLearningProgressResponse, error)
}

type learningProgressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLearningProgressServiceClient(cc grpc.ClientConnInterface) LearningProgressServiceClient {
	return &learningProgressServiceClient{cc}
}

func (c *learningProgressServiceClient) CreateLearningProgress(ctx context.Context, in *CreateLearningProgressRequest, opts ...grpc.CallOption) (*CreateLearningProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLearningProgressResponse)
	err := c.cc.Invoke(ctx, LearningProgressService_CreateLearningProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LearningProgressServiceServer is the server API for LearningProgressService service.
// All implementations must embed UnimplementedLearningProgressServiceServer
// for forward compatibility.
//
// endpoints
type LearningProgressServiceServer interface {
	CreateLearningProgress(context.Context, *CreateLearningProgressRequest) (*CreateLearningProgressResponse, error)
	mustEmbedUnimplementedLearningProgressServiceServer()
}

// UnimplementedLearningProgressServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLearningProgressServiceServer struct{}

func (UnimplementedLearningProgressServiceServer) CreateLearningProgress(context.Context, *CreateLearningProgressRequest) (*CreateLearningProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLearningProgress not implemented")
}
func (UnimplementedLearningProgressServiceServer) mustEmbedUnimplementedLearningProgressServiceServer() {
}
func (UnimplementedLearningProgressServiceServer) testEmbeddedByValue() {}

// UnsafeLearningProgressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LearningProgressServiceServer will
// result in compilation errors.
type UnsafeLearningProgressServiceServer interface {
	mustEmbedUnimplementedLearningProgressServiceServer()
}

func RegisterLearningProgressServiceServer(s grpc.ServiceRegistrar, srv LearningProgressServiceServer) {
	// If the following call pancis, it indicates UnimplementedLearningProgressServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LearningProgressService_ServiceDesc, srv)
}

func _LearningProgressService_CreateLearningProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLearningProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningProgressServiceServer).CreateLearningProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningProgressService_CreateLearningProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningProgressServiceServer).CreateLearningProgress(ctx, req.(*CreateLearningProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LearningProgressService_ServiceDesc is the grpc.ServiceDesc for LearningProgressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LearningProgressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "learningProgress.LearningProgressService",
	HandlerType: (*LearningProgressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLearningProgress",
			Handler:    _LearningProgressService_CreateLearningProgress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "learningProgress/learningProgress.proto",
}
