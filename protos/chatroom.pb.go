// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: protos/chatroom.proto

package protos

import (
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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chatroom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chatroom_proto_msgTypes[0]
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
	return file_protos_chatroom_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ChatToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ChatToken) Reset() {
	*x = ChatToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chatroom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatToken) ProtoMessage() {}

func (x *ChatToken) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chatroom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatToken.ProtoReflect.Descriptor instead.
func (*ChatToken) Descriptor() ([]byte, []int) {
	return file_protos_chatroom_proto_rawDescGZIP(), []int{1}
}

func (x *ChatToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Login) Reset() {
	*x = Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chatroom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login) ProtoMessage() {}

func (x *Login) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chatroom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login.ProtoReflect.Descriptor instead.
func (*Login) Descriptor() ([]byte, []int) {
	return file_protos_chatroom_proto_rawDescGZIP(), []int{2}
}

func (x *Login) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Logout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Logout) Reset() {
	*x = Logout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chatroom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logout) ProtoMessage() {}

func (x *Logout) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chatroom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logout.ProtoReflect.Descriptor instead.
func (*Logout) Descriptor() ([]byte, []int) {
	return file_protos_chatroom_proto_rawDescGZIP(), []int{3}
}

func (x *Logout) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content   string                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Username  string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chatroom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chatroom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_protos_chatroom_proto_rawDescGZIP(), []int{4}
}

func (x *ChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ChatMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type StreamMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Message:
	//
	//	*StreamMessage_ClientLogin
	//	*StreamMessage_ClientMessage
	//	*StreamMessage_ClientLogout
	Message isStreamMessage_Message `protobuf_oneof:"message"`
}

func (x *StreamMessage) Reset() {
	*x = StreamMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chatroom_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessage) ProtoMessage() {}

func (x *StreamMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chatroom_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessage.ProtoReflect.Descriptor instead.
func (*StreamMessage) Descriptor() ([]byte, []int) {
	return file_protos_chatroom_proto_rawDescGZIP(), []int{5}
}

func (x *StreamMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (m *StreamMessage) GetMessage() isStreamMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *StreamMessage) GetClientLogin() *Login {
	if x, ok := x.GetMessage().(*StreamMessage_ClientLogin); ok {
		return x.ClientLogin
	}
	return nil
}

func (x *StreamMessage) GetClientMessage() *ChatMessage {
	if x, ok := x.GetMessage().(*StreamMessage_ClientMessage); ok {
		return x.ClientMessage
	}
	return nil
}

func (x *StreamMessage) GetClientLogout() *Logout {
	if x, ok := x.GetMessage().(*StreamMessage_ClientLogout); ok {
		return x.ClientLogout
	}
	return nil
}

type isStreamMessage_Message interface {
	isStreamMessage_Message()
}

type StreamMessage_ClientLogin struct {
	ClientLogin *Login `protobuf:"bytes,2,opt,name=client_login,json=clientLogin,proto3,oneof"`
}

type StreamMessage_ClientMessage struct {
	ClientMessage *ChatMessage `protobuf:"bytes,3,opt,name=client_message,json=clientMessage,proto3,oneof"`
}

type StreamMessage_ClientLogout struct {
	ClientLogout *Logout `protobuf:"bytes,4,opt,name=client_logout,json=clientLogout,proto3,oneof"`
}

func (*StreamMessage_ClientLogin) isStreamMessage_Message() {}

func (*StreamMessage_ClientMessage) isStreamMessage_Message() {}

func (*StreamMessage_ClientLogout) isStreamMessage_Message() {}

var File_protos_chatroom_proto protoreflect.FileDescriptor

var file_protos_chatroom_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f,
	0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21,
	0x0a, 0x09, 0x43, 0x68, 0x61, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x23, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7d, 0x0a, 0x0b,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x83, 0x02, 0x0a, 0x0d,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x34, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x00,
	0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x3e, 0x0a,
	0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a,
	0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2e,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x87, 0x01, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x39,
	0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x64, 0x6b, 0x6e, 0x69,
	0x67, 0x68, 0x74, 0x32, 0x34, 0x2f, 0x74, 0x65, 0x72, 0x6d, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_chatroom_proto_rawDescOnce sync.Once
	file_protos_chatroom_proto_rawDescData = file_protos_chatroom_proto_rawDesc
)

func file_protos_chatroom_proto_rawDescGZIP() []byte {
	file_protos_chatroom_proto_rawDescOnce.Do(func() {
		file_protos_chatroom_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_chatroom_proto_rawDescData)
	})
	return file_protos_chatroom_proto_rawDescData
}

var file_protos_chatroom_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_chatroom_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),          // 0: chatroom.LoginRequest
	(*ChatToken)(nil),             // 1: chatroom.ChatToken
	(*Login)(nil),                 // 2: chatroom.Login
	(*Logout)(nil),                // 3: chatroom.Logout
	(*ChatMessage)(nil),           // 4: chatroom.ChatMessage
	(*StreamMessage)(nil),         // 5: chatroom.StreamMessage
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_protos_chatroom_proto_depIdxs = []int32{
	6, // 0: chatroom.ChatMessage.timestamp:type_name -> google.protobuf.Timestamp
	6, // 1: chatroom.StreamMessage.timestamp:type_name -> google.protobuf.Timestamp
	2, // 2: chatroom.StreamMessage.client_login:type_name -> chatroom.Login
	4, // 3: chatroom.StreamMessage.client_message:type_name -> chatroom.ChatMessage
	3, // 4: chatroom.StreamMessage.client_logout:type_name -> chatroom.Logout
	0, // 5: chatroom.ChatRoom.JoinChat:input_type -> chatroom.LoginRequest
	5, // 6: chatroom.ChatRoom.Stream:input_type -> chatroom.StreamMessage
	1, // 7: chatroom.ChatRoom.JoinChat:output_type -> chatroom.ChatToken
	5, // 8: chatroom.ChatRoom.Stream:output_type -> chatroom.StreamMessage
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_chatroom_proto_init() }
func file_protos_chatroom_proto_init() {
	if File_protos_chatroom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_chatroom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_chatroom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatToken); i {
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
		file_protos_chatroom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Login); i {
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
		file_protos_chatroom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logout); i {
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
		file_protos_chatroom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_protos_chatroom_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMessage); i {
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
	file_protos_chatroom_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*StreamMessage_ClientLogin)(nil),
		(*StreamMessage_ClientMessage)(nil),
		(*StreamMessage_ClientLogout)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_chatroom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_chatroom_proto_goTypes,
		DependencyIndexes: file_protos_chatroom_proto_depIdxs,
		MessageInfos:      file_protos_chatroom_proto_msgTypes,
	}.Build()
	File_protos_chatroom_proto = out.File
	file_protos_chatroom_proto_rawDesc = nil
	file_protos_chatroom_proto_goTypes = nil
	file_protos_chatroom_proto_depIdxs = nil
}
