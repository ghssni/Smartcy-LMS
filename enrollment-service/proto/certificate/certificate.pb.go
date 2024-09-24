// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: certificate/certificate.proto

package certificate

import (
	meta "github.com/ghssni/Smartcy-LMS/Enrollment-Service/proto/meta"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EnrollmentId   uint32                 `protobuf:"varint,2,opt,name=enrollmentId,proto3" json:"enrollmentId,omitempty"`
	CertificateUrl string                 `protobuf:"bytes,3,opt,name=certificateUrl,proto3" json:"certificateUrl,omitempty"`
	IssuedAt       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=issuedAt,proto3" json:"issuedAt,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_certificate_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *Certificate) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Certificate) GetEnrollmentId() uint32 {
	if x != nil {
		return x.EnrollmentId
	}
	return 0
}

func (x *Certificate) GetCertificateUrl() string {
	if x != nil {
		return x.CertificateUrl
	}
	return ""
}

func (x *Certificate) GetIssuedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IssuedAt
	}
	return nil
}

func (x *Certificate) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Certificate) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnrollmentId   uint32                 `protobuf:"varint,1,opt,name=enrollmentId,proto3" json:"enrollmentId,omitempty"`
	CertificateUrl string                 `protobuf:"bytes,2,opt,name=certificateUrl,proto3" json:"certificateUrl,omitempty"`
	IssuedAt       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=issuedAt,proto3" json:"issuedAt,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *CreateCertificateRequest) Reset() {
	*x = CreateCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_certificate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCertificateRequest) ProtoMessage() {}

func (x *CreateCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_certificate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCertificateRequest.ProtoReflect.Descriptor instead.
func (*CreateCertificateRequest) Descriptor() ([]byte, []int) {
	return file_certificate_certificate_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCertificateRequest) GetEnrollmentId() uint32 {
	if x != nil {
		return x.EnrollmentId
	}
	return 0
}

func (x *CreateCertificateRequest) GetCertificateUrl() string {
	if x != nil {
		return x.CertificateUrl
	}
	return ""
}

func (x *CreateCertificateRequest) GetIssuedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IssuedAt
	}
	return nil
}

func (x *CreateCertificateRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CreateCertificateRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateCertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta        *meta.Meta   `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Certificate *Certificate `protobuf:"bytes,2,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *CreateCertificateResponse) Reset() {
	*x = CreateCertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_certificate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCertificateResponse) ProtoMessage() {}

func (x *CreateCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_certificate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCertificateResponse.ProtoReflect.Descriptor instead.
func (*CreateCertificateResponse) Descriptor() ([]byte, []int) {
	return file_certificate_certificate_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCertificateResponse) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *CreateCertificateResponse) GetCertificate() *Certificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

type GetCertificateByEnrollmentIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnrollmentId uint32 `protobuf:"varint,1,opt,name=enrollmentId,proto3" json:"enrollmentId,omitempty"`
}

func (x *GetCertificateByEnrollmentIdRequest) Reset() {
	*x = GetCertificateByEnrollmentIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_certificate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertificateByEnrollmentIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertificateByEnrollmentIdRequest) ProtoMessage() {}

func (x *GetCertificateByEnrollmentIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_certificate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertificateByEnrollmentIdRequest.ProtoReflect.Descriptor instead.
func (*GetCertificateByEnrollmentIdRequest) Descriptor() ([]byte, []int) {
	return file_certificate_certificate_proto_rawDescGZIP(), []int{3}
}

func (x *GetCertificateByEnrollmentIdRequest) GetEnrollmentId() uint32 {
	if x != nil {
		return x.EnrollmentId
	}
	return 0
}

type GetCertificateByEnrollmentIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta        *meta.Meta   `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Certificate *Certificate `protobuf:"bytes,2,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *GetCertificateByEnrollmentIdResponse) Reset() {
	*x = GetCertificateByEnrollmentIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_certificate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertificateByEnrollmentIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertificateByEnrollmentIdResponse) ProtoMessage() {}

func (x *GetCertificateByEnrollmentIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_certificate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertificateByEnrollmentIdResponse.ProtoReflect.Descriptor instead.
func (*GetCertificateByEnrollmentIdResponse) Descriptor() ([]byte, []int) {
	return file_certificate_certificate_proto_rawDescGZIP(), []int{4}
}

func (x *GetCertificateByEnrollmentIdResponse) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GetCertificateByEnrollmentIdResponse) GetCertificate() *Certificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

var File_certificate_certificate_proto protoreflect.FileDescriptor

var file_certificate_certificate_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95,
	0x02, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x36, 0x0a, 0x08, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x92, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x36, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x77, 0x0a, 0x19, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x22, 0x49, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x79, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x82, 0x01, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x32, 0xfe, 0x01, 0x0a, 0x12, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x25, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x83, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x30, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x79, 0x45,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x3b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_certificate_certificate_proto_rawDescOnce sync.Once
	file_certificate_certificate_proto_rawDescData = file_certificate_certificate_proto_rawDesc
)

func file_certificate_certificate_proto_rawDescGZIP() []byte {
	file_certificate_certificate_proto_rawDescOnce.Do(func() {
		file_certificate_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_certificate_certificate_proto_rawDescData)
	})
	return file_certificate_certificate_proto_rawDescData
}

var file_certificate_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_certificate_certificate_proto_goTypes = []any{
	(*Certificate)(nil),                          // 0: certificate.Certificate
	(*CreateCertificateRequest)(nil),             // 1: certificate.CreateCertificateRequest
	(*CreateCertificateResponse)(nil),            // 2: certificate.CreateCertificateResponse
	(*GetCertificateByEnrollmentIdRequest)(nil),  // 3: certificate.GetCertificateByEnrollmentIdRequest
	(*GetCertificateByEnrollmentIdResponse)(nil), // 4: certificate.GetCertificateByEnrollmentIdResponse
	(*timestamppb.Timestamp)(nil),                // 5: google.protobuf.Timestamp
	(*meta.Meta)(nil),                            // 6: meta.Meta
}
var file_certificate_certificate_proto_depIdxs = []int32{
	5,  // 0: certificate.Certificate.issuedAt:type_name -> google.protobuf.Timestamp
	5,  // 1: certificate.Certificate.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 2: certificate.Certificate.updatedAt:type_name -> google.protobuf.Timestamp
	5,  // 3: certificate.CreateCertificateRequest.issuedAt:type_name -> google.protobuf.Timestamp
	5,  // 4: certificate.CreateCertificateRequest.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 5: certificate.CreateCertificateRequest.updatedAt:type_name -> google.protobuf.Timestamp
	6,  // 6: certificate.CreateCertificateResponse.meta:type_name -> meta.Meta
	0,  // 7: certificate.CreateCertificateResponse.certificate:type_name -> certificate.Certificate
	6,  // 8: certificate.GetCertificateByEnrollmentIdResponse.meta:type_name -> meta.Meta
	0,  // 9: certificate.GetCertificateByEnrollmentIdResponse.certificate:type_name -> certificate.Certificate
	1,  // 10: certificate.CertificateService.CreateCertificate:input_type -> certificate.CreateCertificateRequest
	3,  // 11: certificate.CertificateService.GetCertificateByEnrollmentId:input_type -> certificate.GetCertificateByEnrollmentIdRequest
	2,  // 12: certificate.CertificateService.CreateCertificate:output_type -> certificate.CreateCertificateResponse
	4,  // 13: certificate.CertificateService.GetCertificateByEnrollmentId:output_type -> certificate.GetCertificateByEnrollmentIdResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_certificate_certificate_proto_init() }
func file_certificate_certificate_proto_init() {
	if File_certificate_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_certificate_certificate_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Certificate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_certificate_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCertificateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_certificate_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCertificateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_certificate_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetCertificateByEnrollmentIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_certificate_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetCertificateByEnrollmentIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_certificate_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_certificate_certificate_proto_goTypes,
		DependencyIndexes: file_certificate_certificate_proto_depIdxs,
		MessageInfos:      file_certificate_certificate_proto_msgTypes,
	}.Build()
	File_certificate_certificate_proto = out.File
	file_certificate_certificate_proto_rawDesc = nil
	file_certificate_certificate_proto_goTypes = nil
	file_certificate_certificate_proto_depIdxs = nil
}
