// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0
// source: relation.proto

package relation

import (
	common "github.com/czhi-bin/mini-tiktok-backend/biz/model/common"

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

type RelationActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                              // User authentication token
	ToUserId   int64  `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`     // Target user ID
	ActionType int32  `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1 - Follow, 2 - Unfollow
}

func (x *RelationActionRequest) Reset() {
	*x = RelationActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionRequest) ProtoMessage() {}

func (x *RelationActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionRequest.ProtoReflect.Descriptor instead.
func (*RelationActionRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{0}
}

func (x *RelationActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RelationActionRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *RelationActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type RelationActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"` // 0 - success, others - failure
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // Status description
}

func (x *RelationActionResponse) Reset() {
	*x = RelationActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionResponse) ProtoMessage() {}

func (x *RelationActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionResponse.ProtoReflect.Descriptor instead.
func (*RelationActionResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{1}
}

func (x *RelationActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type RelationFollowListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // User ID
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                  // User authentication token
}

func (x *RelationFollowListRequest) Reset() {
	*x = RelationFollowListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowListRequest) ProtoMessage() {}

func (x *RelationFollowListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowListRequest.ProtoReflect.Descriptor instead.
func (*RelationFollowListRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{2}
}

func (x *RelationFollowListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelationFollowListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RelationFollowListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32          `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"` // 0 - success, others - failure
	StatusMsg  string         `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // Status description
	UserList   []*common.User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`        // List of user information
}

func (x *RelationFollowListResponse) Reset() {
	*x = RelationFollowListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowListResponse) ProtoMessage() {}

func (x *RelationFollowListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowListResponse.ProtoReflect.Descriptor instead.
func (*RelationFollowListResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{3}
}

func (x *RelationFollowListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationFollowListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *RelationFollowListResponse) GetUserList() []*common.User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type RelationFollowerListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // User ID
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                  // User authentication token
}

func (x *RelationFollowerListRequest) Reset() {
	*x = RelationFollowerListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowerListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowerListRequest) ProtoMessage() {}

func (x *RelationFollowerListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowerListRequest.ProtoReflect.Descriptor instead.
func (*RelationFollowerListRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{4}
}

func (x *RelationFollowerListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelationFollowerListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RelationFollowerListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32          `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"` // 0 - success, others - failure
	StatusMsg  string         `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // Status description
	UserList   []*common.User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`        // List of user information
}

func (x *RelationFollowerListResponse) Reset() {
	*x = RelationFollowerListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowerListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowerListResponse) ProtoMessage() {}

func (x *RelationFollowerListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowerListResponse.ProtoReflect.Descriptor instead.
func (*RelationFollowerListResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{5}
}

func (x *RelationFollowerListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationFollowerListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *RelationFollowerListResponse) GetUserList() []*common.User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type RelationFriendListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // User ID
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                  // User authentication token
}

func (x *RelationFriendListRequest) Reset() {
	*x = RelationFriendListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFriendListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFriendListRequest) ProtoMessage() {}

func (x *RelationFriendListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFriendListRequest.ProtoReflect.Descriptor instead.
func (*RelationFriendListRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{6}
}

func (x *RelationFriendListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelationFriendListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RelationFriendListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"` // 0 - success, others - failure
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // Status description
	UserList   []*FriendUser `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`        // List of user information
}

func (x *RelationFriendListResponse) Reset() {
	*x = RelationFriendListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFriendListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFriendListResponse) ProtoMessage() {}

func (x *RelationFriendListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFriendListResponse.ProtoReflect.Descriptor instead.
func (*RelationFriendListResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{7}
}

func (x *RelationFriendListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationFriendListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *RelationFriendListResponse) GetUserList() []*FriendUser {
	if x != nil {
		return x.UserList
	}
	return nil
}

type FriendUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                 // latest chat messages with this friend
	MsgType int64  `protobuf:"varint,3,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"` // 0 - Messages current requested by the user, 1 - Messages sent by the current requesting user
}

func (x *FriendUser) Reset() {
	*x = FriendUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendUser) ProtoMessage() {}

func (x *FriendUser) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendUser.ProtoReflect.Descriptor instead.
func (*FriendUser) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{8}
}

func (x *FriendUser) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FriendUser) GetMsgType() int64 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

type FriendUser_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *FriendUser_User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FriendUser_User) Reset() {
	*x = FriendUser_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendUser_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendUser_User) ProtoMessage() {}

func (x *FriendUser_User) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendUser_User.ProtoReflect.Descriptor instead.
func (*FriendUser_User) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{8, 0}
}

func (x *FriendUser_User) GetUser() *FriendUser_User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_relation_proto protoreflect.FileDescriptor

var file_relation_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x70, 0x0a, 0x16, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x16, 0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x73, 0x67, 0x22, 0x4a, 0x0a, 0x19, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x98, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x16, 0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x1b, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9a, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x16,
	0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73,
	0x67, 0x12, 0x22, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x19, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0xa7, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x16, 0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x31, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x78, 0x0a, 0x0a, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x35,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0x9b, 0x04, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x71, 0x0a, 0x0e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c,
	0xd2, 0xc1, 0x18, 0x18, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x12, 0x82, 0x01, 0x0a,
	0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21,
	0xca, 0xc1, 0x18, 0x1d, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x2f, 0x12, 0x8a, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0xca, 0xc1, 0x18, 0x1f, 0x2f,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x12, 0x82,
	0x01, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x21, 0xca, 0xc1, 0x18, 0x1d, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x2f, 0x42, 0x11, 0x5a, 0x0f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relation_proto_rawDescOnce sync.Once
	file_relation_proto_rawDescData = file_relation_proto_rawDesc
)

func file_relation_proto_rawDescGZIP() []byte {
	file_relation_proto_rawDescOnce.Do(func() {
		file_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_relation_proto_rawDescData)
	})
	return file_relation_proto_rawDescData
}

var file_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_relation_proto_goTypes = []interface{}{
	(*RelationActionRequest)(nil),        // 0: relation.RelationActionRequest
	(*RelationActionResponse)(nil),       // 1: relation.RelationActionResponse
	(*RelationFollowListRequest)(nil),    // 2: relation.RelationFollowListRequest
	(*RelationFollowListResponse)(nil),   // 3: relation.RelationFollowListResponse
	(*RelationFollowerListRequest)(nil),  // 4: relation.RelationFollowerListRequest
	(*RelationFollowerListResponse)(nil), // 5: relation.RelationFollowerListResponse
	(*RelationFriendListRequest)(nil),    // 6: relation.RelationFriendListRequest
	(*RelationFriendListResponse)(nil),   // 7: relation.RelationFriendListResponse
	(*FriendUser)(nil),                   // 8: relation.FriendUser
	(*FriendUser_User)(nil),              // 9: relation.FriendUser.User
	(*common.User)(nil),                  // 10: User
}
var file_relation_proto_depIdxs = []int32{
	10, // 0: relation.RelationFollowListResponse.user_list:type_name -> User
	10, // 1: relation.RelationFollowerListResponse.user_list:type_name -> User
	8,  // 2: relation.RelationFriendListResponse.user_list:type_name -> relation.FriendUser
	9,  // 3: relation.FriendUser.User.user:type_name -> relation.FriendUser.User
	0,  // 4: relation.RelationHandler.RelationAction:input_type -> relation.RelationActionRequest
	2,  // 5: relation.RelationHandler.RelationFollowList:input_type -> relation.RelationFollowListRequest
	4,  // 6: relation.RelationHandler.RelationFollowerList:input_type -> relation.RelationFollowerListRequest
	6,  // 7: relation.RelationHandler.RelationFriendList:input_type -> relation.RelationFriendListRequest
	1,  // 8: relation.RelationHandler.RelationAction:output_type -> relation.RelationActionResponse
	3,  // 9: relation.RelationHandler.RelationFollowList:output_type -> relation.RelationFollowListResponse
	5,  // 10: relation.RelationHandler.RelationFollowerList:output_type -> relation.RelationFollowerListResponse
	7,  // 11: relation.RelationHandler.RelationFriendList:output_type -> relation.RelationFriendListResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_relation_proto_init() }
func file_relation_proto_init() {
	if File_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionRequest); i {
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
		file_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionResponse); i {
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
		file_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowListRequest); i {
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
		file_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowListResponse); i {
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
		file_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowerListRequest); i {
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
		file_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowerListResponse); i {
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
		file_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFriendListRequest); i {
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
		file_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFriendListResponse); i {
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
		file_relation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendUser); i {
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
		file_relation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendUser_User); i {
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
			RawDescriptor: file_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relation_proto_goTypes,
		DependencyIndexes: file_relation_proto_depIdxs,
		MessageInfos:      file_relation_proto_msgTypes,
	}.Build()
	File_relation_proto = out.File
	file_relation_proto_rawDesc = nil
	file_relation_proto_goTypes = nil
	file_relation_proto_depIdxs = nil
}
