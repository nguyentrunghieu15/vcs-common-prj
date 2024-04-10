// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: apu/user/user.proto

package user

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	server "github.com/nguyentrunghieu15/vcs-common-prj/apu/server"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UserRole int32

const (
	UserRole_RoleAdmin UserRole = 0
	UserRole_RoleUser  UserRole = 1
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "RoleAdmin",
		1: "RoleUser",
	}
	UserRole_value = map[string]int32{
		"RoleAdmin": 0,
		"RoleUser":  1,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_apu_user_user_proto_enumTypes[0].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_apu_user_user_proto_enumTypes[0]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_apu_user_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     string   `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy     int64    `protobuf:"varint,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt     string   `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy     int64    `protobuf:"varint,5,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Email         string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	FullName      string   `protobuf:"bytes,7,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone         string   `protobuf:"bytes,8,opt,name=phone,proto3" json:"phone,omitempty"`
	Avatar        string   `protobuf:"bytes,9,opt,name=avatar,proto3" json:"avatar,omitempty"`
	IsSupperAdmin bool     `protobuf:"varint,10,opt,name=is_supper_admin,json=isSupperAdmin,proto3" json:"is_supper_admin,omitempty"`
	Roles         UserRole `protobuf:"varint,11,opt,name=roles,proto3,enum=UserRole" json:"roles,omitempty"`
	Password      string   `protobuf:"bytes,12,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apu_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_apu_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_apu_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *User) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *User) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetIsSupperAdmin() bool {
	if x != nil {
		return x.IsSupperAdmin
	}
	return false
}

func (x *User) GetRoles() UserRole {
	if x != nil {
		return x.Roles
	}
	return UserRole_RoleAdmin
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ListUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *server.Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apu_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apu_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_apu_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *ListUsersRequest) GetFilter() *server.Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apu_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apu_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_apu_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *ListUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserByIdRequest) Reset() {
	*x = GetUserByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apu_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdRequest) ProtoMessage() {}

func (x *GetUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apu_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_apu_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserByIdRequest) Reset() {
	*x = DeleteUserByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apu_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserByIdRequest) ProtoMessage() {}

func (x *DeleteUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apu_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_apu_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUserByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email         string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	FullName      string   `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone         string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Avatar        string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	IsSupperAdmin bool     `protobuf:"varint,5,opt,name=is_supper_admin,json=isSupperAdmin,proto3" json:"is_supper_admin,omitempty"`
	Roles         UserRole `protobuf:"varint,6,opt,name=roles,proto3,enum=UserRole" json:"roles,omitempty"`
	Password      string   `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apu_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apu_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_apu_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateUserRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *CreateUserRequest) GetIsSupperAdmin() bool {
	if x != nil {
		return x.IsSupperAdmin
	}
	return false
}

func (x *CreateUserRequest) GetRoles() UserRole {
	if x != nil {
		return x.Roles
	}
	return UserRole_RoleAdmin
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateUserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email         string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	FullName      string   `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone         string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Avatar        string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	IsSupperAdmin bool     `protobuf:"varint,5,opt,name=is_supper_admin,json=isSupperAdmin,proto3" json:"is_supper_admin,omitempty"`
	Roles         UserRole `protobuf:"varint,6,opt,name=roles,proto3,enum=UserRole" json:"roles,omitempty"`
	Password      string   `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	Id            int64    `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateUserByIdRequest) Reset() {
	*x = UpdateUserByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apu_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserByIdRequest) ProtoMessage() {}

func (x *UpdateUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apu_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserByIdRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_apu_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserByIdRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserByIdRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateUserByIdRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateUserByIdRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UpdateUserByIdRequest) GetIsSupperAdmin() bool {
	if x != nil {
		return x.IsSupperAdmin
	}
	return false
}

func (x *UpdateUserByIdRequest) GetRoles() UserRole {
	if x != nil {
		return x.Roles
	}
	return UserRole_RoleAdmin
}

func (x *UpdateUserByIdRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateUserByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_apu_user_user_proto protoreflect.FileDescriptor

var file_apu_user_user_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x75, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x61, 0x70, 0x75, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe1, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x53, 0x75,
	0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x33, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x24, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd9, 0x01, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73,
	0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xed, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73,
	0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x27, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x10,
	0x01, 0x32, 0xf3, 0x05, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x96, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x11, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x62, 0x92, 0x41, 0x4b, 0x4a, 0x32, 0x0a, 0x03, 0x34,
	0x30, 0x31, 0x12, 0x2b, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x72,
	0x15, 0x0a, 0x13, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x28, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x8e, 0x01, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x22, 0x67, 0x92, 0x41, 0x4b, 0x4a, 0x32, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x2b,
	0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x72, 0x15, 0x0a, 0x13, 0x0a,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x28, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x8e, 0x01, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x65, 0x92, 0x41, 0x4b, 0x4a, 0x32, 0x0a, 0x03, 0x34, 0x30,
	0x31, 0x12, 0x2b, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x72, 0x15,
	0x0a, 0x13, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x28, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x80, 0x01, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x53, 0x92, 0x41, 0x34, 0x4a,
	0x32, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x2b, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x32, 0x11, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0xa5, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x67,
	0x92, 0x41, 0x4b, 0x4a, 0x32, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x2b, 0x0a, 0x0c, 0x55, 0x6e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17,
	0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70,
	0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x72, 0x15, 0x0a, 0x13, 0x0a, 0x0d, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x28, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x67, 0x75, 0x79, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x6e,
	0x67, 0x68, 0x69, 0x65, 0x75, 0x31, 0x35, 0x2f, 0x76, 0x63, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6a, 0x2f, 0x61, 0x70, 0x75, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apu_user_user_proto_rawDescOnce sync.Once
	file_apu_user_user_proto_rawDescData = file_apu_user_user_proto_rawDesc
)

func file_apu_user_user_proto_rawDescGZIP() []byte {
	file_apu_user_user_proto_rawDescOnce.Do(func() {
		file_apu_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_apu_user_user_proto_rawDescData)
	})
	return file_apu_user_user_proto_rawDescData
}

var file_apu_user_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apu_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_apu_user_user_proto_goTypes = []interface{}{
	(UserRole)(0),                 // 0: UserRole
	(*User)(nil),                  // 1: User
	(*ListUsersRequest)(nil),      // 2: ListUsersRequest
	(*ListUsersResponse)(nil),     // 3: ListUsersResponse
	(*GetUserByIdRequest)(nil),    // 4: GetUserByIdRequest
	(*DeleteUserByIdRequest)(nil), // 5: DeleteUserByIdRequest
	(*CreateUserRequest)(nil),     // 6: CreateUserRequest
	(*UpdateUserByIdRequest)(nil), // 7: UpdateUserByIdRequest
	(*server.Filter)(nil),         // 8: Filter
	(*empty.Empty)(nil),           // 9: google.protobuf.Empty
}
var file_apu_user_user_proto_depIdxs = []int32{
	0,  // 0: User.roles:type_name -> UserRole
	8,  // 1: ListUsersRequest.filter:type_name -> Filter
	1,  // 2: ListUsersResponse.users:type_name -> User
	0,  // 3: CreateUserRequest.roles:type_name -> UserRole
	0,  // 4: UpdateUserByIdRequest.roles:type_name -> UserRole
	2,  // 5: UserService.ListUsers:input_type -> ListUsersRequest
	4,  // 6: UserService.GetUser:input_type -> GetUserByIdRequest
	6,  // 7: UserService.CreateUser:input_type -> CreateUserRequest
	7,  // 8: UserService.UpdateUser:input_type -> UpdateUserByIdRequest
	5,  // 9: UserService.DeleteUser:input_type -> DeleteUserByIdRequest
	3,  // 10: UserService.ListUsers:output_type -> ListUsersResponse
	1,  // 11: UserService.GetUser:output_type -> User
	1,  // 12: UserService.CreateUser:output_type -> User
	1,  // 13: UserService.UpdateUser:output_type -> User
	9,  // 14: UserService.DeleteUser:output_type -> google.protobuf.Empty
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_apu_user_user_proto_init() }
func file_apu_user_user_proto_init() {
	if File_apu_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apu_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_apu_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersRequest); i {
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
		file_apu_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersResponse); i {
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
		file_apu_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdRequest); i {
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
		file_apu_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserByIdRequest); i {
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
		file_apu_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_apu_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserByIdRequest); i {
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
			RawDescriptor: file_apu_user_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apu_user_user_proto_goTypes,
		DependencyIndexes: file_apu_user_user_proto_depIdxs,
		EnumInfos:         file_apu_user_user_proto_enumTypes,
		MessageInfos:      file_apu_user_user_proto_msgTypes,
	}.Build()
	File_apu_user_user_proto = out.File
	file_apu_user_user_proto_rawDesc = nil
	file_apu_user_user_proto_goTypes = nil
	file_apu_user_user_proto_depIdxs = nil
}
