// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: assessments/assessments.proto

package assessments

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
	AssessmentsService_CreateAssessment_FullMethodName         = "/assessments.AssessmentsService/CreateAssessment"
	AssessmentsService_GetAssessmentByStudentId_FullMethodName = "/assessments.AssessmentsService/GetAssessmentByStudentId"
)

// AssessmentsServiceClient is the client API for AssessmentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// endpoints
type AssessmentsServiceClient interface {
	CreateAssessment(ctx context.Context, in *CreateAssessmentRequest, opts ...grpc.CallOption) (*CreateAssessmentResponse, error)
	GetAssessmentByStudentId(ctx context.Context, in *GetAssessmentByStudentIdRequest, opts ...grpc.CallOption) (*GetAssessmentByStudentIdResponse, error)
}

type assessmentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssessmentsServiceClient(cc grpc.ClientConnInterface) AssessmentsServiceClient {
	return &assessmentsServiceClient{cc}
}

func (c *assessmentsServiceClient) CreateAssessment(ctx context.Context, in *CreateAssessmentRequest, opts ...grpc.CallOption) (*CreateAssessmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAssessmentResponse)
	err := c.cc.Invoke(ctx, AssessmentsService_CreateAssessment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assessmentsServiceClient) GetAssessmentByStudentId(ctx context.Context, in *GetAssessmentByStudentIdRequest, opts ...grpc.CallOption) (*GetAssessmentByStudentIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAssessmentByStudentIdResponse)
	err := c.cc.Invoke(ctx, AssessmentsService_GetAssessmentByStudentId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssessmentsServiceServer is the server API for AssessmentsService service.
// All implementations must embed UnimplementedAssessmentsServiceServer
// for forward compatibility.
//
// endpoints
type AssessmentsServiceServer interface {
	CreateAssessment(context.Context, *CreateAssessmentRequest) (*CreateAssessmentResponse, error)
	GetAssessmentByStudentId(context.Context, *GetAssessmentByStudentIdRequest) (*GetAssessmentByStudentIdResponse, error)
	mustEmbedUnimplementedAssessmentsServiceServer()
}

// UnimplementedAssessmentsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAssessmentsServiceServer struct{}

func (UnimplementedAssessmentsServiceServer) CreateAssessment(context.Context, *CreateAssessmentRequest) (*CreateAssessmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAssessment not implemented")
}
func (UnimplementedAssessmentsServiceServer) GetAssessmentByStudentId(context.Context, *GetAssessmentByStudentIdRequest) (*GetAssessmentByStudentIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssessmentByStudentId not implemented")
}
func (UnimplementedAssessmentsServiceServer) mustEmbedUnimplementedAssessmentsServiceServer() {}
func (UnimplementedAssessmentsServiceServer) testEmbeddedByValue()                            {}

// UnsafeAssessmentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssessmentsServiceServer will
// result in compilation errors.
type UnsafeAssessmentsServiceServer interface {
	mustEmbedUnimplementedAssessmentsServiceServer()
}

func RegisterAssessmentsServiceServer(s grpc.ServiceRegistrar, srv AssessmentsServiceServer) {
	// If the following call pancis, it indicates UnimplementedAssessmentsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AssessmentsService_ServiceDesc, srv)
}

func _AssessmentsService_CreateAssessment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssessmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssessmentsServiceServer).CreateAssessment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssessmentsService_CreateAssessment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssessmentsServiceServer).CreateAssessment(ctx, req.(*CreateAssessmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssessmentsService_GetAssessmentByStudentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssessmentByStudentIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssessmentsServiceServer).GetAssessmentByStudentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssessmentsService_GetAssessmentByStudentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssessmentsServiceServer).GetAssessmentByStudentId(ctx, req.(*GetAssessmentByStudentIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssessmentsService_ServiceDesc is the grpc.ServiceDesc for AssessmentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssessmentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "assessments.AssessmentsService",
	HandlerType: (*AssessmentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAssessment",
			Handler:    _AssessmentsService_CreateAssessment_Handler,
		},
		{
			MethodName: "GetAssessmentByStudentId",
			Handler:    _AssessmentsService_GetAssessmentByStudentId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assessments/assessments.proto",
}
