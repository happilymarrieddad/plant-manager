// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: users.proto

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

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	JWT string `protobuf:"bytes,2,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

type GetUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserReply) Reset() {
	*x = GetUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReply) ProtoMessage() {}

func (x *GetUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReply.ProtoReflect.Descriptor instead.
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserReply) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type VerifyJWTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route string `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	JWT   string `protobuf:"bytes,2,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *VerifyJWTRequest) Reset() {
	*x = VerifyJWTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyJWTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyJWTRequest) ProtoMessage() {}

func (x *VerifyJWTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyJWTRequest.ProtoReflect.Descriptor instead.
func (*VerifyJWTRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyJWTRequest) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *VerifyJWTRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

type VerifyJWTReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid         bool          `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	HasPermission bool          `protobuf:"varint,2,opt,name=hasPermission,proto3" json:"hasPermission,omitempty"`
	Permissions   []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *VerifyJWTReply) Reset() {
	*x = VerifyJWTReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyJWTReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyJWTReply) ProtoMessage() {}

func (x *VerifyJWTReply) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyJWTReply.ProtoReflect.Descriptor instead.
func (*VerifyJWTReply) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyJWTReply) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *VerifyJWTReply) GetHasPermission() bool {
	if x != nil {
		return x.HasPermission
	}
	return false
}

func (x *VerifyJWTReply) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Permissions []*Permission `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{5}
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginReply) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type AddUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	JWT       string `protobuf:"bytes,4,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *AddUserRequest) Reset() {
	*x = AddUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRequest) ProtoMessage() {}

func (x *AddUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRequest.ProtoReflect.Descriptor instead.
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{6}
}

func (x *AddUserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *AddUserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *AddUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddUserRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

type AddUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *AddUserReply) Reset() {
	*x = AddUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserReply) ProtoMessage() {}

func (x *AddUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserReply.ProtoReflect.Descriptor instead.
func (*AddUserReply) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{7}
}

func (x *AddUserReply) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type FindUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	JWT    string `protobuf:"bytes,3,opt,name=JWT,proto3" json:"JWT,omitempty"`
}

func (x *FindUsersRequest) Reset() {
	*x = FindUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUsersRequest) ProtoMessage() {}

func (x *FindUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUsersRequest.ProtoReflect.Descriptor instead.
func (*FindUsersRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{8}
}

func (x *FindUsersRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindUsersRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FindUsersRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

type FindUsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *FindUsersReply) Reset() {
	*x = FindUsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUsersReply) ProtoMessage() {}

func (x *FindUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUsersReply.ProtoReflect.Descriptor instead.
func (*FindUsersReply) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{9}
}

func (x *FindUsersReply) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_users_proto protoreflect.FileDescriptor

var file_users_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x32, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4a, 0x57, 0x54, 0x22, 0x2f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x4a, 0x57, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a,
	0x57, 0x54, 0x22, 0x81, 0x01, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4a, 0x57, 0x54,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x68,
	0x61, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x33, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x40, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x57, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x33, 0x0a, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x72, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4a, 0x57, 0x54, 0x22, 0x2f, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x52, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x22, 0x33, 0x0a, 0x0e, 0x46, 0x69,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32,
	0xac, 0x02, 0x0a, 0x07, 0x56, 0x31, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4a, 0x57, 0x54, 0x12, 0x17, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x10,
	0x5a, 0x0e, 0x2e, 0x3b, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_proto_rawDescOnce sync.Once
	file_users_proto_rawDescData = file_users_proto_rawDesc
)

func file_users_proto_rawDescGZIP() []byte {
	file_users_proto_rawDescOnce.Do(func() {
		file_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_proto_rawDescData)
	})
	return file_users_proto_rawDescData
}

var file_users_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_users_proto_goTypes = []interface{}{
	(*GetUserRequest)(nil),   // 0: users.GetUserRequest
	(*GetUserReply)(nil),     // 1: users.GetUserReply
	(*VerifyJWTRequest)(nil), // 2: users.VerifyJWTRequest
	(*VerifyJWTReply)(nil),   // 3: users.VerifyJWTReply
	(*LoginRequest)(nil),     // 4: users.LoginRequest
	(*LoginReply)(nil),       // 5: users.LoginReply
	(*AddUserRequest)(nil),   // 6: users.AddUserRequest
	(*AddUserReply)(nil),     // 7: users.AddUserReply
	(*FindUsersRequest)(nil), // 8: users.FindUsersRequest
	(*FindUsersReply)(nil),   // 9: users.FindUsersReply
	(*User)(nil),             // 10: types.User
	(*Permission)(nil),       // 11: types.Permission
}
var file_users_proto_depIdxs = []int32{
	10, // 0: users.GetUserReply.user:type_name -> types.User
	11, // 1: users.VerifyJWTReply.permissions:type_name -> types.Permission
	11, // 2: users.LoginReply.permissions:type_name -> types.Permission
	10, // 3: users.AddUserReply.user:type_name -> types.User
	10, // 4: users.FindUsersReply.users:type_name -> types.User
	4,  // 5: users.V1Users.Login:input_type -> users.LoginRequest
	2,  // 6: users.V1Users.VerifyJWT:input_type -> users.VerifyJWTRequest
	6,  // 7: users.V1Users.AddUser:input_type -> users.AddUserRequest
	0,  // 8: users.V1Users.GetUser:input_type -> users.GetUserRequest
	8,  // 9: users.V1Users.FindUsers:input_type -> users.FindUsersRequest
	5,  // 10: users.V1Users.Login:output_type -> users.LoginReply
	3,  // 11: users.V1Users.VerifyJWT:output_type -> users.VerifyJWTReply
	7,  // 12: users.V1Users.AddUser:output_type -> users.AddUserReply
	1,  // 13: users.V1Users.GetUser:output_type -> users.GetUserReply
	9,  // 14: users.V1Users.FindUsers:output_type -> users.FindUsersReply
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_users_proto_init() }
func file_users_proto_init() {
	if File_users_proto != nil {
		return
	}
	file_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReply); i {
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
		file_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyJWTRequest); i {
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
		file_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyJWTReply); i {
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
		file_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_users_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserRequest); i {
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
		file_users_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserReply); i {
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
		file_users_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUsersRequest); i {
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
		file_users_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUsersReply); i {
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
			RawDescriptor: file_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_proto_goTypes,
		DependencyIndexes: file_users_proto_depIdxs,
		MessageInfos:      file_users_proto_msgTypes,
	}.Build()
	File_users_proto = out.File
	file_users_proto_rawDesc = nil
	file_users_proto_goTypes = nil
	file_users_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// V1UsersClient is the client API for V1Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type V1UsersClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	VerifyJWT(ctx context.Context, in *VerifyJWTRequest, opts ...grpc.CallOption) (*VerifyJWTReply, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserReply, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	FindUsers(ctx context.Context, in *FindUsersRequest, opts ...grpc.CallOption) (*FindUsersReply, error)
}

type v1UsersClient struct {
	cc grpc.ClientConnInterface
}

func NewV1UsersClient(cc grpc.ClientConnInterface) V1UsersClient {
	return &v1UsersClient{cc}
}

func (c *v1UsersClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/users.V1Users/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1UsersClient) VerifyJWT(ctx context.Context, in *VerifyJWTRequest, opts ...grpc.CallOption) (*VerifyJWTReply, error) {
	out := new(VerifyJWTReply)
	err := c.cc.Invoke(ctx, "/users.V1Users/VerifyJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1UsersClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserReply, error) {
	out := new(AddUserReply)
	err := c.cc.Invoke(ctx, "/users.V1Users/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1UsersClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, "/users.V1Users/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1UsersClient) FindUsers(ctx context.Context, in *FindUsersRequest, opts ...grpc.CallOption) (*FindUsersReply, error) {
	out := new(FindUsersReply)
	err := c.cc.Invoke(ctx, "/users.V1Users/FindUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// V1UsersServer is the server API for V1Users service.
type V1UsersServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	VerifyJWT(context.Context, *VerifyJWTRequest) (*VerifyJWTReply, error)
	AddUser(context.Context, *AddUserRequest) (*AddUserReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	FindUsers(context.Context, *FindUsersRequest) (*FindUsersReply, error)
}

// UnimplementedV1UsersServer can be embedded to have forward compatible implementations.
type UnimplementedV1UsersServer struct {
}

func (*UnimplementedV1UsersServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedV1UsersServer) VerifyJWT(context.Context, *VerifyJWTRequest) (*VerifyJWTReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyJWT not implemented")
}
func (*UnimplementedV1UsersServer) AddUser(context.Context, *AddUserRequest) (*AddUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (*UnimplementedV1UsersServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedV1UsersServer) FindUsers(context.Context, *FindUsersRequest) (*FindUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUsers not implemented")
}

func RegisterV1UsersServer(s *grpc.Server, srv V1UsersServer) {
	s.RegisterService(&_V1Users_serviceDesc, srv)
}

func _V1Users_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1UsersServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.V1Users/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1UsersServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Users_VerifyJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyJWTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1UsersServer).VerifyJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.V1Users/VerifyJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1UsersServer).VerifyJWT(ctx, req.(*VerifyJWTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Users_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1UsersServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.V1Users/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1UsersServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Users_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1UsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.V1Users/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1UsersServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1Users_FindUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1UsersServer).FindUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.V1Users/FindUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1UsersServer).FindUsers(ctx, req.(*FindUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _V1Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.V1Users",
	HandlerType: (*V1UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _V1Users_Login_Handler,
		},
		{
			MethodName: "VerifyJWT",
			Handler:    _V1Users_VerifyJWT_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _V1Users_AddUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _V1Users_GetUser_Handler,
		},
		{
			MethodName: "FindUsers",
			Handler:    _V1Users_FindUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
