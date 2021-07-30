// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: varieties.proto

package plantmanager

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetVarietyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	JWT string `protobuf:"bytes,2,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *GetVarietyRequest) Reset() {
	*x = GetVarietyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_varieties_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVarietyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVarietyRequest) ProtoMessage() {}

func (x *GetVarietyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_varieties_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVarietyRequest.ProtoReflect.Descriptor instead.
func (*GetVarietyRequest) Descriptor() ([]byte, []int) {
	return file_varieties_proto_rawDescGZIP(), []int{0}
}

func (x *GetVarietyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetVarietyRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

type GetVarietyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variety *Variety `protobuf:"bytes,1,opt,name=variety,proto3" json:"variety,omitempty"`
}

func (x *GetVarietyReply) Reset() {
	*x = GetVarietyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_varieties_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVarietyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVarietyReply) ProtoMessage() {}

func (x *GetVarietyReply) ProtoReflect() protoreflect.Message {
	mi := &file_varieties_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVarietyReply.ProtoReflect.Descriptor instead.
func (*GetVarietyReply) Descriptor() ([]byte, []int) {
	return file_varieties_proto_rawDescGZIP(), []int{1}
}

func (x *GetVarietyReply) GetVariety() *Variety {
	if x != nil {
		return x.Variety
	}
	return nil
}

type FindVarietiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	JWT    string `protobuf:"bytes,3,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *FindVarietiesRequest) Reset() {
	*x = FindVarietiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_varieties_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindVarietiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindVarietiesRequest) ProtoMessage() {}

func (x *FindVarietiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_varieties_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindVarietiesRequest.ProtoReflect.Descriptor instead.
func (*FindVarietiesRequest) Descriptor() ([]byte, []int) {
	return file_varieties_proto_rawDescGZIP(), []int{2}
}

func (x *FindVarietiesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindVarietiesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FindVarietiesRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

type FindVarietiesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Varieties []*Variety `protobuf:"bytes,1,rep,name=varieties,proto3" json:"varieties,omitempty"`
}

func (x *FindVarietiesReply) Reset() {
	*x = FindVarietiesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_varieties_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindVarietiesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindVarietiesReply) ProtoMessage() {}

func (x *FindVarietiesReply) ProtoReflect() protoreflect.Message {
	mi := &file_varieties_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindVarietiesReply.ProtoReflect.Descriptor instead.
func (*FindVarietiesReply) Descriptor() ([]byte, []int) {
	return file_varieties_proto_rawDescGZIP(), []int{3}
}

func (x *FindVarietiesReply) GetVarieties() []*Variety {
	if x != nil {
		return x.Varieties
	}
	return nil
}

type CreateVarietyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PlantTypeId int64  `protobuf:"varint,2,opt,name=plantTypeId,proto3" json:"plantTypeId,omitempty"`
	JWT         string `protobuf:"bytes,3,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *CreateVarietyRequest) Reset() {
	*x = CreateVarietyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_varieties_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVarietyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVarietyRequest) ProtoMessage() {}

func (x *CreateVarietyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_varieties_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVarietyRequest.ProtoReflect.Descriptor instead.
func (*CreateVarietyRequest) Descriptor() ([]byte, []int) {
	return file_varieties_proto_rawDescGZIP(), []int{4}
}

func (x *CreateVarietyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateVarietyRequest) GetPlantTypeId() int64 {
	if x != nil {
		return x.PlantTypeId
	}
	return 0
}

func (x *CreateVarietyRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

type CreateVarietyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variety *Variety `protobuf:"bytes,1,opt,name=variety,proto3" json:"variety,omitempty"`
}

func (x *CreateVarietyReply) Reset() {
	*x = CreateVarietyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_varieties_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVarietyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVarietyReply) ProtoMessage() {}

func (x *CreateVarietyReply) ProtoReflect() protoreflect.Message {
	mi := &file_varieties_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVarietyReply.ProtoReflect.Descriptor instead.
func (*CreateVarietyReply) Descriptor() ([]byte, []int) {
	return file_varieties_proto_rawDescGZIP(), []int{5}
}

func (x *CreateVarietyReply) GetVariety() *Variety {
	if x != nil {
		return x.Variety
	}
	return nil
}

type DestroyVarietyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	JWT string `protobuf:"bytes,2,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *DestroyVarietyRequest) Reset() {
	*x = DestroyVarietyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_varieties_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyVarietyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyVarietyRequest) ProtoMessage() {}

func (x *DestroyVarietyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_varieties_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyVarietyRequest.ProtoReflect.Descriptor instead.
func (*DestroyVarietyRequest) Descriptor() ([]byte, []int) {
	return file_varieties_proto_rawDescGZIP(), []int{6}
}

func (x *DestroyVarietyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DestroyVarietyRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

type UpdateVarietyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PlantTypeId int64  `protobuf:"varint,3,opt,name=plantTypeId,proto3" json:"plantTypeId,omitempty"`
	JWT         string `protobuf:"bytes,4,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *UpdateVarietyRequest) Reset() {
	*x = UpdateVarietyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_varieties_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVarietyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVarietyRequest) ProtoMessage() {}

func (x *UpdateVarietyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_varieties_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVarietyRequest.ProtoReflect.Descriptor instead.
func (*UpdateVarietyRequest) Descriptor() ([]byte, []int) {
	return file_varieties_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateVarietyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateVarietyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateVarietyRequest) GetPlantTypeId() int64 {
	if x != nil {
		return x.PlantTypeId
	}
	return 0
}

func (x *UpdateVarietyRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

var File_varieties_proto protoreflect.FileDescriptor

var file_varieties_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x0b, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54,
	0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x72,
	0x69, 0x65, 0x74, 0x79, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x22, 0x56, 0x0a,
	0x14, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4a, 0x57, 0x54, 0x22, 0x42, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x72,
	0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x09, 0x76,
	0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x09,
	0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x22, 0x3e, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x28, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79,
	0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x22, 0x39, 0x0a, 0x15, 0x44, 0x65, 0x73,
	0x74, 0x72, 0x6f, 0x79, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4a, 0x57, 0x54, 0x22, 0x6e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x61,
	0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4a, 0x57, 0x54, 0x32, 0x8d, 0x03, 0x0a, 0x0b, 0x56, 0x31, 0x56, 0x61, 0x72, 0x69, 0x65,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x61, 0x72, 0x69, 0x65,
	0x74, 0x79, 0x12, 0x1c, 0x2e, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x51,
	0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x1f, 0x2e, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x65,
	0x74, 0x79, 0x12, 0x1f, 0x2e, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x61,
	0x72, 0x69, 0x65, 0x74, 0x79, 0x12, 0x1f, 0x2e, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e, 0x44,
	0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x12, 0x20, 0x2e,
	0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f,
	0x79, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x3b, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_varieties_proto_rawDescOnce sync.Once
	file_varieties_proto_rawDescData = file_varieties_proto_rawDesc
)

func file_varieties_proto_rawDescGZIP() []byte {
	file_varieties_proto_rawDescOnce.Do(func() {
		file_varieties_proto_rawDescData = protoimpl.X.CompressGZIP(file_varieties_proto_rawDescData)
	})
	return file_varieties_proto_rawDescData
}

var file_varieties_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_varieties_proto_goTypes = []interface{}{
	(*GetVarietyRequest)(nil),     // 0: varieties.GetVarietyRequest
	(*GetVarietyReply)(nil),       // 1: varieties.GetVarietyReply
	(*FindVarietiesRequest)(nil),  // 2: varieties.FindVarietiesRequest
	(*FindVarietiesReply)(nil),    // 3: varieties.FindVarietiesReply
	(*CreateVarietyRequest)(nil),  // 4: varieties.CreateVarietyRequest
	(*CreateVarietyReply)(nil),    // 5: varieties.CreateVarietyReply
	(*DestroyVarietyRequest)(nil), // 6: varieties.DestroyVarietyRequest
	(*UpdateVarietyRequest)(nil),  // 7: varieties.UpdateVarietyRequest
	(*Variety)(nil),               // 8: types.Variety
	(*EmptyReply)(nil),            // 9: types.EmptyReply
}
var file_varieties_proto_depIdxs = []int32{
	8, // 0: varieties.GetVarietyReply.variety:type_name -> types.Variety
	8, // 1: varieties.FindVarietiesReply.varieties:type_name -> types.Variety
	8, // 2: varieties.CreateVarietyReply.variety:type_name -> types.Variety
	0, // 3: varieties.V1Varieties.GetVariety:input_type -> varieties.GetVarietyRequest
	2, // 4: varieties.V1Varieties.FindVarieties:input_type -> varieties.FindVarietiesRequest
	4, // 5: varieties.V1Varieties.CreateVariety:input_type -> varieties.CreateVarietyRequest
	7, // 6: varieties.V1Varieties.UpdateVariety:input_type -> varieties.UpdateVarietyRequest
	6, // 7: varieties.V1Varieties.DestroyVariety:input_type -> varieties.DestroyVarietyRequest
	1, // 8: varieties.V1Varieties.GetVariety:output_type -> varieties.GetVarietyReply
	3, // 9: varieties.V1Varieties.FindVarieties:output_type -> varieties.FindVarietiesReply
	5, // 10: varieties.V1Varieties.CreateVariety:output_type -> varieties.CreateVarietyReply
	9, // 11: varieties.V1Varieties.UpdateVariety:output_type -> types.EmptyReply
	9, // 12: varieties.V1Varieties.DestroyVariety:output_type -> types.EmptyReply
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_varieties_proto_init() }
func file_varieties_proto_init() {
	if File_varieties_proto != nil {
		return
	}
	file_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_varieties_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVarietyRequest); i {
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
		file_varieties_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVarietyReply); i {
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
		file_varieties_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindVarietiesRequest); i {
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
		file_varieties_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindVarietiesReply); i {
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
		file_varieties_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVarietyRequest); i {
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
		file_varieties_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVarietyReply); i {
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
		file_varieties_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyVarietyRequest); i {
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
		file_varieties_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVarietyRequest); i {
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
			RawDescriptor: file_varieties_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_varieties_proto_goTypes,
		DependencyIndexes: file_varieties_proto_depIdxs,
		MessageInfos:      file_varieties_proto_msgTypes,
	}.Build()
	File_varieties_proto = out.File
	file_varieties_proto_rawDesc = nil
	file_varieties_proto_goTypes = nil
	file_varieties_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// V1VarietiesClient is the client API for V1Varieties service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type V1VarietiesClient interface {
	GetVariety(ctx context.Context, in *GetVarietyRequest, opts ...grpc.CallOption) (*GetVarietyReply, error)
	FindVarieties(ctx context.Context, in *FindVarietiesRequest, opts ...grpc.CallOption) (*FindVarietiesReply, error)
	CreateVariety(ctx context.Context, in *CreateVarietyRequest, opts ...grpc.CallOption) (*CreateVarietyReply, error)
	UpdateVariety(ctx context.Context, in *UpdateVarietyRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	DestroyVariety(ctx context.Context, in *DestroyVarietyRequest, opts ...grpc.CallOption) (*EmptyReply, error)
}

type v1VarietiesClient struct {
	cc grpc.ClientConnInterface
}

func NewV1VarietiesClient(cc grpc.ClientConnInterface) V1VarietiesClient {
	return &v1VarietiesClient{cc}
}

func (c *v1VarietiesClient) GetVariety(ctx context.Context, in *GetVarietyRequest, opts ...grpc.CallOption) (*GetVarietyReply, error) {
	out := new(GetVarietyReply)
	err := c.cc.Invoke(ctx, "/varieties.V1Varieties/GetVariety", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1VarietiesClient) FindVarieties(ctx context.Context, in *FindVarietiesRequest, opts ...grpc.CallOption) (*FindVarietiesReply, error) {
	out := new(FindVarietiesReply)
	err := c.cc.Invoke(ctx, "/varieties.V1Varieties/FindVarieties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1VarietiesClient) CreateVariety(ctx context.Context, in *CreateVarietyRequest, opts ...grpc.CallOption) (*CreateVarietyReply, error) {
	out := new(CreateVarietyReply)
	err := c.cc.Invoke(ctx, "/varieties.V1Varieties/CreateVariety", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1VarietiesClient) UpdateVariety(ctx context.Context, in *UpdateVarietyRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/varieties.V1Varieties/UpdateVariety", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1VarietiesClient) DestroyVariety(ctx context.Context, in *DestroyVarietyRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/varieties.V1Varieties/DestroyVariety", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// V1VarietiesServer is the server API for V1Varieties service.
type V1VarietiesServer interface {
	GetVariety(context.Context, *GetVarietyRequest) (*GetVarietyReply, error)
	FindVarieties(context.Context, *FindVarietiesRequest) (*FindVarietiesReply, error)
	CreateVariety(context.Context, *CreateVarietyRequest) (*CreateVarietyReply, error)
	UpdateVariety(context.Context, *UpdateVarietyRequest) (*EmptyReply, error)
	DestroyVariety(context.Context, *DestroyVarietyRequest) (*EmptyReply, error)
}

// UnimplementedV1VarietiesServer can be embedded to have forward compatible implementations.
type UnimplementedV1VarietiesServer struct {
}

func (*UnimplementedV1VarietiesServer) GetVariety(context.Context, *GetVarietyRequest) (*GetVarietyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVariety not implemented")
}
func (*UnimplementedV1VarietiesServer) FindVarieties(context.Context, *FindVarietiesRequest) (*FindVarietiesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindVarieties not implemented")
}
func (*UnimplementedV1VarietiesServer) CreateVariety(context.Context, *CreateVarietyRequest) (*CreateVarietyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVariety not implemented")
}
func (*UnimplementedV1VarietiesServer) UpdateVariety(context.Context, *UpdateVarietyRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVariety not implemented")
}
func (*UnimplementedV1VarietiesServer) DestroyVariety(context.Context, *DestroyVarietyRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyVariety not implemented")
}

func RegisterV1VarietiesServer(s *grpc.Server, srv V1VarietiesServer) {
	s.RegisterService(&_V1Varieties_serviceDesc, srv)
}

func _V1Varieties_GetVariety_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVarietyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1VarietiesServer).GetVariety(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/varieties.V1Varieties/GetVariety",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1VarietiesServer).GetVariety(ctx, req.(*GetVarietyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Varieties_FindVarieties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindVarietiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1VarietiesServer).FindVarieties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/varieties.V1Varieties/FindVarieties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1VarietiesServer).FindVarieties(ctx, req.(*FindVarietiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Varieties_CreateVariety_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVarietyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1VarietiesServer).CreateVariety(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/varieties.V1Varieties/CreateVariety",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1VarietiesServer).CreateVariety(ctx, req.(*CreateVarietyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Varieties_UpdateVariety_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVarietyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1VarietiesServer).UpdateVariety(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/varieties.V1Varieties/UpdateVariety",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1VarietiesServer).UpdateVariety(ctx, req.(*UpdateVarietyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Varieties_DestroyVariety_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyVarietyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1VarietiesServer).DestroyVariety(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/varieties.V1Varieties/DestroyVariety",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1VarietiesServer).DestroyVariety(ctx, req.(*DestroyVarietyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _V1Varieties_serviceDesc = grpc.ServiceDesc{
	ServiceName: "varieties.V1Varieties",
	HandlerType: (*V1VarietiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVariety",
			Handler:    _V1Varieties_GetVariety_Handler,
		},
		{
			MethodName: "FindVarieties",
			Handler:    _V1Varieties_FindVarieties_Handler,
		},
		{
			MethodName: "CreateVariety",
			Handler:    _V1Varieties_CreateVariety_Handler,
		},
		{
			MethodName: "UpdateVariety",
			Handler:    _V1Varieties_UpdateVariety_Handler,
		},
		{
			MethodName: "DestroyVariety",
			Handler:    _V1Varieties_DestroyVariety_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "varieties.proto",
}
