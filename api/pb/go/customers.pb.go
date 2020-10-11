// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: customers.proto

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

type GetCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	JWT string `protobuf:"bytes,2,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *GetCustomerRequest) Reset() {
	*x = GetCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRequest) ProtoMessage() {}

func (x *GetCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customers_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCustomerRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

type GetCustomerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *GetCustomerReply) Reset() {
	*x = GetCustomerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerReply) ProtoMessage() {}

func (x *GetCustomerReply) ProtoReflect() protoreflect.Message {
	mi := &file_customers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerReply.ProtoReflect.Descriptor instead.
func (*GetCustomerReply) Descriptor() ([]byte, []int) {
	return file_customers_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomerReply) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type FindCustomersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	JWT    string `protobuf:"bytes,3,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *FindCustomersRequest) Reset() {
	*x = FindCustomersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCustomersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCustomersRequest) ProtoMessage() {}

func (x *FindCustomersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCustomersRequest.ProtoReflect.Descriptor instead.
func (*FindCustomersRequest) Descriptor() ([]byte, []int) {
	return file_customers_proto_rawDescGZIP(), []int{2}
}

func (x *FindCustomersRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindCustomersRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FindCustomersRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

type FindCustomersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customers []*Customer `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
}

func (x *FindCustomersReply) Reset() {
	*x = FindCustomersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCustomersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCustomersReply) ProtoMessage() {}

func (x *FindCustomersReply) ProtoReflect() protoreflect.Message {
	mi := &file_customers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCustomersReply.ProtoReflect.Descriptor instead.
func (*FindCustomersReply) Descriptor() ([]byte, []int) {
	return file_customers_proto_rawDescGZIP(), []int{3}
}

func (x *FindCustomersReply) GetCustomers() []*Customer {
	if x != nil {
		return x.Customers
	}
	return nil
}

type CreateCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	JWT  string `protobuf:"bytes,2,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *CreateCustomerRequest) Reset() {
	*x = CreateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerRequest) ProtoMessage() {}

func (x *CreateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customers_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCustomerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCustomerRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

type CreateCustomerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *CreateCustomerReply) Reset() {
	*x = CreateCustomerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerReply) ProtoMessage() {}

func (x *CreateCustomerReply) ProtoReflect() protoreflect.Message {
	mi := &file_customers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerReply.ProtoReflect.Descriptor instead.
func (*CreateCustomerReply) Descriptor() ([]byte, []int) {
	return file_customers_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCustomerReply) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type DestroyCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	JWT string `protobuf:"bytes,2,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *DestroyCustomerRequest) Reset() {
	*x = DestroyCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyCustomerRequest) ProtoMessage() {}

func (x *DestroyCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyCustomerRequest.ProtoReflect.Descriptor instead.
func (*DestroyCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customers_proto_rawDescGZIP(), []int{6}
}

func (x *DestroyCustomerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DestroyCustomerRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

var File_customers_proto protoreflect.FileDescriptor

var file_customers_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x1a, 0x0b, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57,
	0x54, 0x22, 0x3f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x22, 0x56, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x22, 0x43, 0x0a, 0x12, 0x46, 0x69,
	0x6e, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2d, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x22,
	0x3d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x4a, 0x57, 0x54, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x22, 0x42,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x22, 0x3a, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x4a, 0x57, 0x54, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x32, 0xce,
	0x02, 0x0a, 0x0b, 0x56, 0x31, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x4b,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x46,
	0x69, 0x6e, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x54,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x12, 0x20, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x10, 0x5a, 0x0e, 0x2e, 0x3b, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_proto_rawDescOnce sync.Once
	file_customers_proto_rawDescData = file_customers_proto_rawDesc
)

func file_customers_proto_rawDescGZIP() []byte {
	file_customers_proto_rawDescOnce.Do(func() {
		file_customers_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_proto_rawDescData)
	})
	return file_customers_proto_rawDescData
}

var file_customers_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_customers_proto_goTypes = []interface{}{
	(*GetCustomerRequest)(nil),     // 0: customers.GetCustomerRequest
	(*GetCustomerReply)(nil),       // 1: customers.GetCustomerReply
	(*FindCustomersRequest)(nil),   // 2: customers.FindCustomersRequest
	(*FindCustomersReply)(nil),     // 3: customers.FindCustomersReply
	(*CreateCustomerRequest)(nil),  // 4: customers.CreateCustomerRequest
	(*CreateCustomerReply)(nil),    // 5: customers.CreateCustomerReply
	(*DestroyCustomerRequest)(nil), // 6: customers.DestroyCustomerRequest
	(*Customer)(nil),               // 7: types.Customer
	(*EmptyReply)(nil),             // 8: types.EmptyReply
}
var file_customers_proto_depIdxs = []int32{
	7, // 0: customers.GetCustomerReply.customer:type_name -> types.Customer
	7, // 1: customers.FindCustomersReply.customers:type_name -> types.Customer
	7, // 2: customers.CreateCustomerReply.customer:type_name -> types.Customer
	0, // 3: customers.V1Customers.GetCustomer:input_type -> customers.GetCustomerRequest
	2, // 4: customers.V1Customers.FindCustomers:input_type -> customers.FindCustomersRequest
	4, // 5: customers.V1Customers.CreateCustomer:input_type -> customers.CreateCustomerRequest
	6, // 6: customers.V1Customers.DestroyCustomer:input_type -> customers.DestroyCustomerRequest
	1, // 7: customers.V1Customers.GetCustomer:output_type -> customers.GetCustomerReply
	3, // 8: customers.V1Customers.FindCustomers:output_type -> customers.FindCustomersReply
	5, // 9: customers.V1Customers.CreateCustomer:output_type -> customers.CreateCustomerReply
	8, // 10: customers.V1Customers.DestroyCustomer:output_type -> types.EmptyReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_customers_proto_init() }
func file_customers_proto_init() {
	if File_customers_proto != nil {
		return
	}
	file_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_customers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerRequest); i {
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
		file_customers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerReply); i {
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
		file_customers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCustomersRequest); i {
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
		file_customers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCustomersReply); i {
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
		file_customers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerRequest); i {
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
		file_customers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerReply); i {
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
		file_customers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyCustomerRequest); i {
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
			RawDescriptor: file_customers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customers_proto_goTypes,
		DependencyIndexes: file_customers_proto_depIdxs,
		MessageInfos:      file_customers_proto_msgTypes,
	}.Build()
	File_customers_proto = out.File
	file_customers_proto_rawDesc = nil
	file_customers_proto_goTypes = nil
	file_customers_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// V1CustomersClient is the client API for V1Customers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type V1CustomersClient interface {
	GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*GetCustomerReply, error)
	FindCustomers(ctx context.Context, in *FindCustomersRequest, opts ...grpc.CallOption) (*FindCustomersReply, error)
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerReply, error)
	DestroyCustomer(ctx context.Context, in *DestroyCustomerRequest, opts ...grpc.CallOption) (*EmptyReply, error)
}

type v1CustomersClient struct {
	cc grpc.ClientConnInterface
}

func NewV1CustomersClient(cc grpc.ClientConnInterface) V1CustomersClient {
	return &v1CustomersClient{cc}
}

func (c *v1CustomersClient) GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*GetCustomerReply, error) {
	out := new(GetCustomerReply)
	err := c.cc.Invoke(ctx, "/customers.V1Customers/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1CustomersClient) FindCustomers(ctx context.Context, in *FindCustomersRequest, opts ...grpc.CallOption) (*FindCustomersReply, error) {
	out := new(FindCustomersReply)
	err := c.cc.Invoke(ctx, "/customers.V1Customers/FindCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1CustomersClient) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerReply, error) {
	out := new(CreateCustomerReply)
	err := c.cc.Invoke(ctx, "/customers.V1Customers/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1CustomersClient) DestroyCustomer(ctx context.Context, in *DestroyCustomerRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/customers.V1Customers/DestroyCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// V1CustomersServer is the server API for V1Customers service.
type V1CustomersServer interface {
	GetCustomer(context.Context, *GetCustomerRequest) (*GetCustomerReply, error)
	FindCustomers(context.Context, *FindCustomersRequest) (*FindCustomersReply, error)
	CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerReply, error)
	DestroyCustomer(context.Context, *DestroyCustomerRequest) (*EmptyReply, error)
}

// UnimplementedV1CustomersServer can be embedded to have forward compatible implementations.
type UnimplementedV1CustomersServer struct {
}

func (*UnimplementedV1CustomersServer) GetCustomer(context.Context, *GetCustomerRequest) (*GetCustomerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (*UnimplementedV1CustomersServer) FindCustomers(context.Context, *FindCustomersRequest) (*FindCustomersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCustomers not implemented")
}
func (*UnimplementedV1CustomersServer) CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (*UnimplementedV1CustomersServer) DestroyCustomer(context.Context, *DestroyCustomerRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyCustomer not implemented")
}

func RegisterV1CustomersServer(s *grpc.Server, srv V1CustomersServer) {
	s.RegisterService(&_V1Customers_serviceDesc, srv)
}

func _V1Customers_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1CustomersServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customers.V1Customers/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1CustomersServer).GetCustomer(ctx, req.(*GetCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Customers_FindCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCustomersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1CustomersServer).FindCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customers.V1Customers/FindCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1CustomersServer).FindCustomers(ctx, req.(*FindCustomersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Customers_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1CustomersServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customers.V1Customers/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1CustomersServer).CreateCustomer(ctx, req.(*CreateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Customers_DestroyCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1CustomersServer).DestroyCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customers.V1Customers/DestroyCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1CustomersServer).DestroyCustomer(ctx, req.(*DestroyCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _V1Customers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customers.V1Customers",
	HandlerType: (*V1CustomersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomer",
			Handler:    _V1Customers_GetCustomer_Handler,
		},
		{
			MethodName: "FindCustomers",
			Handler:    _V1Customers_FindCustomers_Handler,
		},
		{
			MethodName: "CreateCustomer",
			Handler:    _V1Customers_CreateCustomer_Handler,
		},
		{
			MethodName: "DestroyCustomer",
			Handler:    _V1Customers_DestroyCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customers.proto",
}
