// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: assessments/assessments.proto

package assessments

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

type Assessments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EnrollmentId   uint32                 `protobuf:"varint,2,opt,name=enrollmentId,proto3" json:"enrollmentId,omitempty"`
	Score          int32                  `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	AssessmentType string                 `protobuf:"bytes,4,opt,name=assessmentType,proto3" json:"assessmentType,omitempty"`
	TakenAt        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=takenAt,proto3" json:"takenAt,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Assessments) Reset() {
	*x = Assessments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assessments_assessments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Assessments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assessments) ProtoMessage() {}

func (x *Assessments) ProtoReflect() protoreflect.Message {
	mi := &file_assessments_assessments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assessments.ProtoReflect.Descriptor instead.
func (*Assessments) Descriptor() ([]byte, []int) {
	return file_assessments_assessments_proto_rawDescGZIP(), []int{0}
}

func (x *Assessments) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Assessments) GetEnrollmentId() uint32 {
	if x != nil {
		return x.EnrollmentId
	}
	return 0
}

func (x *Assessments) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Assessments) GetAssessmentType() string {
	if x != nil {
		return x.AssessmentType
	}
	return ""
}

func (x *Assessments) GetTakenAt() *timestamppb.Timestamp {
	if x != nil {
		return x.TakenAt
	}
	return nil
}

func (x *Assessments) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Assessments) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateAssessmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnrollmentId   uint32                 `protobuf:"varint,1,opt,name=enrollmentId,proto3" json:"enrollmentId,omitempty"`
	Score          int32                  `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	AssessmentType string                 `protobuf:"bytes,3,opt,name=assessmentType,proto3" json:"assessmentType,omitempty"`
	TakenAt        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=takenAt,proto3" json:"takenAt,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *CreateAssessmentRequest) Reset() {
	*x = CreateAssessmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assessments_assessments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAssessmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssessmentRequest) ProtoMessage() {}

func (x *CreateAssessmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_assessments_assessments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssessmentRequest.ProtoReflect.Descriptor instead.
func (*CreateAssessmentRequest) Descriptor() ([]byte, []int) {
	return file_assessments_assessments_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAssessmentRequest) GetEnrollmentId() uint32 {
	if x != nil {
		return x.EnrollmentId
	}
	return 0
}

func (x *CreateAssessmentRequest) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *CreateAssessmentRequest) GetAssessmentType() string {
	if x != nil {
		return x.AssessmentType
	}
	return ""
}

func (x *CreateAssessmentRequest) GetTakenAt() *timestamppb.Timestamp {
	if x != nil {
		return x.TakenAt
	}
	return nil
}

func (x *CreateAssessmentRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CreateAssessmentRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateAssessmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta        *meta.Meta   `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Assessments *Assessments `protobuf:"bytes,2,opt,name=assessments,proto3" json:"assessments,omitempty"`
}

func (x *CreateAssessmentResponse) Reset() {
	*x = CreateAssessmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assessments_assessments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAssessmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssessmentResponse) ProtoMessage() {}

func (x *CreateAssessmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assessments_assessments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssessmentResponse.ProtoReflect.Descriptor instead.
func (*CreateAssessmentResponse) Descriptor() ([]byte, []int) {
	return file_assessments_assessments_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAssessmentResponse) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *CreateAssessmentResponse) GetAssessments() *Assessments {
	if x != nil {
		return x.Assessments
	}
	return nil
}

type GetAssessmentByStudentIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EnrollmentId uint32 `protobuf:"varint,2,opt,name=enrollmentId,proto3" json:"enrollmentId,omitempty"`
}

func (x *GetAssessmentByStudentIdRequest) Reset() {
	*x = GetAssessmentByStudentIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assessments_assessments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssessmentByStudentIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssessmentByStudentIdRequest) ProtoMessage() {}

func (x *GetAssessmentByStudentIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_assessments_assessments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssessmentByStudentIdRequest.ProtoReflect.Descriptor instead.
func (*GetAssessmentByStudentIdRequest) Descriptor() ([]byte, []int) {
	return file_assessments_assessments_proto_rawDescGZIP(), []int{3}
}

func (x *GetAssessmentByStudentIdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetAssessmentByStudentIdRequest) GetEnrollmentId() uint32 {
	if x != nil {
		return x.EnrollmentId
	}
	return 0
}

type GetAssessmentByStudentIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta        *meta.Meta   `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Assessments *Assessments `protobuf:"bytes,2,opt,name=assessments,proto3" json:"assessments,omitempty"`
}

func (x *GetAssessmentByStudentIdResponse) Reset() {
	*x = GetAssessmentByStudentIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assessments_assessments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssessmentByStudentIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssessmentByStudentIdResponse) ProtoMessage() {}

func (x *GetAssessmentByStudentIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assessments_assessments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssessmentByStudentIdResponse.ProtoReflect.Descriptor instead.
func (*GetAssessmentByStudentIdResponse) Descriptor() ([]byte, []int) {
	return file_assessments_assessments_proto_rawDescGZIP(), []int{4}
}

func (x *GetAssessmentByStudentIdResponse) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GetAssessmentByStudentIdResponse) GetAssessments() *Assessments {
	if x != nil {
		return x.Assessments
	}
	return nil
}

var File_assessments_assessments_proto protoreflect.FileDescriptor

var file_assessments_assessments_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x73,
	0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9,
	0x02, 0x0a, 0x0b, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x34, 0x0a, 0x07, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x74,
	0x61, 0x6b, 0x65, 0x6e, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa5, 0x02, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6e,
	0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x74, 0x61, 0x6b, 0x65,
	0x6e, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x41, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x76, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65,
	0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3a,
	0x0a, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0b, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x55, 0x0a, 0x1f, 0x47, 0x65,
	0x74, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x7e, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x32, 0xee, 0x01, 0x0a, 0x12, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x77, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x42,
	0x79, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3b, 0x61, 0x73,
	0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_assessments_assessments_proto_rawDescOnce sync.Once
	file_assessments_assessments_proto_rawDescData = file_assessments_assessments_proto_rawDesc
)

func file_assessments_assessments_proto_rawDescGZIP() []byte {
	file_assessments_assessments_proto_rawDescOnce.Do(func() {
		file_assessments_assessments_proto_rawDescData = protoimpl.X.CompressGZIP(file_assessments_assessments_proto_rawDescData)
	})
	return file_assessments_assessments_proto_rawDescData
}

var file_assessments_assessments_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_assessments_assessments_proto_goTypes = []any{
	(*Assessments)(nil),                      // 0: assessments.Assessments
	(*CreateAssessmentRequest)(nil),          // 1: assessments.CreateAssessmentRequest
	(*CreateAssessmentResponse)(nil),         // 2: assessments.CreateAssessmentResponse
	(*GetAssessmentByStudentIdRequest)(nil),  // 3: assessments.GetAssessmentByStudentIdRequest
	(*GetAssessmentByStudentIdResponse)(nil), // 4: assessments.GetAssessmentByStudentIdResponse
	(*timestamppb.Timestamp)(nil),            // 5: google.protobuf.Timestamp
	(*meta.Meta)(nil),                        // 6: meta.Meta
}
var file_assessments_assessments_proto_depIdxs = []int32{
	5,  // 0: assessments.Assessments.takenAt:type_name -> google.protobuf.Timestamp
	5,  // 1: assessments.Assessments.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 2: assessments.Assessments.updatedAt:type_name -> google.protobuf.Timestamp
	5,  // 3: assessments.CreateAssessmentRequest.takenAt:type_name -> google.protobuf.Timestamp
	5,  // 4: assessments.CreateAssessmentRequest.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 5: assessments.CreateAssessmentRequest.updatedAt:type_name -> google.protobuf.Timestamp
	6,  // 6: assessments.CreateAssessmentResponse.meta:type_name -> meta.Meta
	0,  // 7: assessments.CreateAssessmentResponse.assessments:type_name -> assessments.Assessments
	6,  // 8: assessments.GetAssessmentByStudentIdResponse.meta:type_name -> meta.Meta
	0,  // 9: assessments.GetAssessmentByStudentIdResponse.assessments:type_name -> assessments.Assessments
	1,  // 10: assessments.AssessmentsService.CreateAssessment:input_type -> assessments.CreateAssessmentRequest
	3,  // 11: assessments.AssessmentsService.GetAssessmentByStudentId:input_type -> assessments.GetAssessmentByStudentIdRequest
	2,  // 12: assessments.AssessmentsService.CreateAssessment:output_type -> assessments.CreateAssessmentResponse
	4,  // 13: assessments.AssessmentsService.GetAssessmentByStudentId:output_type -> assessments.GetAssessmentByStudentIdResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_assessments_assessments_proto_init() }
func file_assessments_assessments_proto_init() {
	if File_assessments_assessments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_assessments_assessments_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Assessments); i {
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
		file_assessments_assessments_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAssessmentRequest); i {
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
		file_assessments_assessments_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAssessmentResponse); i {
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
		file_assessments_assessments_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAssessmentByStudentIdRequest); i {
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
		file_assessments_assessments_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAssessmentByStudentIdResponse); i {
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
			RawDescriptor: file_assessments_assessments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_assessments_assessments_proto_goTypes,
		DependencyIndexes: file_assessments_assessments_proto_depIdxs,
		MessageInfos:      file_assessments_assessments_proto_msgTypes,
	}.Build()
	File_assessments_assessments_proto = out.File
	file_assessments_assessments_proto_rawDesc = nil
	file_assessments_assessments_proto_goTypes = nil
	file_assessments_assessments_proto_depIdxs = nil
}
