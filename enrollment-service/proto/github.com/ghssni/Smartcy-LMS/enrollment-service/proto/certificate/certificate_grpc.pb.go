// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: certificate/certificate.proto

package certificate

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
	CertificateService_CreateCertificate_FullMethodName            = "/certificate.CertificateService/CreateCertificate"
	CertificateService_GetCertificateByEnrollmentId_FullMethodName = "/certificate.CertificateService/GetCertificateByEnrollmentId"
)

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// endpoints
type CertificateServiceClient interface {
	CreateCertificate(ctx context.Context, in *CreateCertificateRequest, opts ...grpc.CallOption) (*CreateCertificateResponse, error)
	GetCertificateByEnrollmentId(ctx context.Context, in *GetCertificateByEnrollmentIdRequest, opts ...grpc.CallOption) (*GetCertificateByEnrollmentIdResponse, error)
}

type certificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateServiceClient(cc grpc.ClientConnInterface) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) CreateCertificate(ctx context.Context, in *CreateCertificateRequest, opts ...grpc.CallOption) (*CreateCertificateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCertificateResponse)
	err := c.cc.Invoke(ctx, CertificateService_CreateCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) GetCertificateByEnrollmentId(ctx context.Context, in *GetCertificateByEnrollmentIdRequest, opts ...grpc.CallOption) (*GetCertificateByEnrollmentIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCertificateByEnrollmentIdResponse)
	err := c.cc.Invoke(ctx, CertificateService_GetCertificateByEnrollmentId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
// All implementations must embed UnimplementedCertificateServiceServer
// for forward compatibility.
//
// endpoints
type CertificateServiceServer interface {
	CreateCertificate(context.Context, *CreateCertificateRequest) (*CreateCertificateResponse, error)
	GetCertificateByEnrollmentId(context.Context, *GetCertificateByEnrollmentIdRequest) (*GetCertificateByEnrollmentIdResponse, error)
	mustEmbedUnimplementedCertificateServiceServer()
}

// UnimplementedCertificateServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCertificateServiceServer struct{}

func (UnimplementedCertificateServiceServer) CreateCertificate(context.Context, *CreateCertificateRequest) (*CreateCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) GetCertificateByEnrollmentId(context.Context, *GetCertificateByEnrollmentIdRequest) (*GetCertificateByEnrollmentIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificateByEnrollmentId not implemented")
}
func (UnimplementedCertificateServiceServer) mustEmbedUnimplementedCertificateServiceServer() {}
func (UnimplementedCertificateServiceServer) testEmbeddedByValue()                            {}

// UnsafeCertificateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateServiceServer will
// result in compilation errors.
type UnsafeCertificateServiceServer interface {
	mustEmbedUnimplementedCertificateServiceServer()
}

func RegisterCertificateServiceServer(s grpc.ServiceRegistrar, srv CertificateServiceServer) {
	// If the following call pancis, it indicates UnimplementedCertificateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CertificateService_ServiceDesc, srv)
}

func _CertificateService_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_CreateCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).CreateCertificate(ctx, req.(*CreateCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_GetCertificateByEnrollmentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateByEnrollmentIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).GetCertificateByEnrollmentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_GetCertificateByEnrollmentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).GetCertificateByEnrollmentId(ctx, req.(*GetCertificateByEnrollmentIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateService_ServiceDesc is the grpc.ServiceDesc for CertificateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "certificate.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCertificate",
			Handler:    _CertificateService_CreateCertificate_Handler,
		},
		{
			MethodName: "GetCertificateByEnrollmentId",
			Handler:    _CertificateService_GetCertificateByEnrollmentId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "certificate/certificate.proto",
}
