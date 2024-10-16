// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: notification/notification.proto

package notification

import (
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

// Email
type EmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Otp         string `protobuf:"bytes,2,opt,name=otp,proto3" json:"otp,omitempty"`
	BodyMessage string `protobuf:"bytes,3,opt,name=body_message,json=bodyMessage,proto3" json:"body_message,omitempty"`
}

func (x *EmailRequest) Reset() {
	*x = EmailRequest{}
	mi := &file_notification_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequest) ProtoMessage() {}

func (x *EmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequest.ProtoReflect.Descriptor instead.
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{0}
}

func (x *EmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *EmailRequest) GetBodyMessage() string {
	if x != nil {
		return x.BodyMessage
	}
	return ""
}

type EmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *EmailResponse) Reset() {
	*x = EmailResponse{}
	mi := &file_notification_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailResponse) ProtoMessage() {}

func (x *EmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailResponse.ProtoReflect.Descriptor instead.
func (*EmailResponse) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{1}
}

func (x *EmailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EmailResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Chat
type CreateChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AdminId      int32 `protobuf:"varint,2,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	SuperAdminId int32 `protobuf:"varint,3,opt,name=super_admin_id,json=superAdminId,proto3" json:"super_admin_id,omitempty"`
}

func (x *CreateChatRequest) Reset() {
	*x = CreateChatRequest{}
	mi := &file_notification_notification_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRequest) ProtoMessage() {}

func (x *CreateChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatRequest.ProtoReflect.Descriptor instead.
func (*CreateChatRequest) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{2}
}

func (x *CreateChatRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateChatRequest) GetAdminId() int32 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *CreateChatRequest) GetSuperAdminId() int32 {
	if x != nil {
		return x.SuperAdminId
	}
	return 0
}

type CreateChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId    string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	StartedAt string `protobuf:"bytes,2,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
}

func (x *CreateChatResponse) Reset() {
	*x = CreateChatResponse{}
	mi := &file_notification_notification_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatResponse) ProtoMessage() {}

func (x *CreateChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatResponse.ProtoReflect.Descriptor instead.
func (*CreateChatResponse) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{3}
}

func (x *CreateChatResponse) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *CreateChatResponse) GetStartedAt() string {
	if x != nil {
		return x.StartedAt
	}
	return ""
}

type AddMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId   string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	SenderId int32  `protobuf:"varint,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddMessageRequest) Reset() {
	*x = AddMessageRequest{}
	mi := &file_notification_notification_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessageRequest) ProtoMessage() {}

func (x *AddMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessageRequest.ProtoReflect.Descriptor instead.
func (*AddMessageRequest) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{4}
}

func (x *AddMessageRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *AddMessageRequest) GetSenderId() int32 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *AddMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AddMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int32  `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	SentAt    string `protobuf:"bytes,2,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *AddMessageResponse) Reset() {
	*x = AddMessageResponse{}
	mi := &file_notification_notification_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessageResponse) ProtoMessage() {}

func (x *AddMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessageResponse.ProtoReflect.Descriptor instead.
func (*AddMessageResponse) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{5}
}

func (x *AddMessageResponse) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *AddMessageResponse) GetSentAt() string {
	if x != nil {
		return x.SentAt
	}
	return ""
}

type EndChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *EndChatRequest) Reset() {
	*x = EndChatRequest{}
	mi := &file_notification_notification_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndChatRequest) ProtoMessage() {}

func (x *EndChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndChatRequest.ProtoReflect.Descriptor instead.
func (*EndChatRequest) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{6}
}

func (x *EndChatRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type EndChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId  int32  `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	EndedAt string `protobuf:"bytes,2,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
}

func (x *EndChatResponse) Reset() {
	*x = EndChatResponse{}
	mi := &file_notification_notification_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndChatResponse) ProtoMessage() {}

func (x *EndChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndChatResponse.ProtoReflect.Descriptor instead.
func (*EndChatResponse) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{7}
}

func (x *EndChatResponse) GetChatId() int32 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *EndChatResponse) GetEndedAt() string {
	if x != nil {
		return x.EndedAt
	}
	return ""
}

type StreamMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int32 `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *StreamMessagesRequest) Reset() {
	*x = StreamMessagesRequest{}
	mi := &file_notification_notification_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessagesRequest) ProtoMessage() {}

func (x *StreamMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessagesRequest.ProtoReflect.Descriptor instead.
func (*StreamMessagesRequest) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{8}
}

func (x *StreamMessagesRequest) GetChatId() int32 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type StreamMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int32  `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	SenderId  int32  `protobuf:"varint,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	SentAt    string `protobuf:"bytes,4,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *StreamMessagesResponse) Reset() {
	*x = StreamMessagesResponse{}
	mi := &file_notification_notification_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessagesResponse) ProtoMessage() {}

func (x *StreamMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessagesResponse.ProtoReflect.Descriptor instead.
func (*StreamMessagesResponse) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{9}
}

func (x *StreamMessagesResponse) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *StreamMessagesResponse) GetSenderId() int32 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *StreamMessagesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StreamMessagesResponse) GetSentAt() string {
	if x != nil {
		return x.SentAt
	}
	return ""
}

var File_notification_notification_proto protoreflect.FileDescriptor

var file_notification_notification_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x59, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x64, 0x79, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62,
	0x6f, 0x64, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x0d, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x6d, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x63, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4c, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x22, 0x29, 0x0a, 0x0e, 0x45,
	0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0f, 0x45, 0x6e, 0x64, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x30, 0x0a,
	0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22,
	0x87, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x32, 0xa7, 0x01, 0x0a, 0x0c, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x53, 0x65,
	0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x51, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x57, 0x6f, 0x72, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xd6, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x1c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x45, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e,
	0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a,
	0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x23, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e,
	0x2e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_notification_proto_rawDescOnce sync.Once
	file_notification_notification_proto_rawDescData = file_notification_notification_proto_rawDesc
)

func file_notification_notification_proto_rawDescGZIP() []byte {
	file_notification_notification_proto_rawDescOnce.Do(func() {
		file_notification_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_notification_proto_rawDescData)
	})
	return file_notification_notification_proto_rawDescData
}

var file_notification_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_notification_notification_proto_goTypes = []any{
	(*EmailRequest)(nil),           // 0: notification.EmailRequest
	(*EmailResponse)(nil),          // 1: notification.EmailResponse
	(*CreateChatRequest)(nil),      // 2: notification.CreateChatRequest
	(*CreateChatResponse)(nil),     // 3: notification.CreateChatResponse
	(*AddMessageRequest)(nil),      // 4: notification.AddMessageRequest
	(*AddMessageResponse)(nil),     // 5: notification.AddMessageResponse
	(*EndChatRequest)(nil),         // 6: notification.EndChatRequest
	(*EndChatResponse)(nil),        // 7: notification.EndChatResponse
	(*StreamMessagesRequest)(nil),  // 8: notification.StreamMessagesRequest
	(*StreamMessagesResponse)(nil), // 9: notification.StreamMessagesResponse
}
var file_notification_notification_proto_depIdxs = []int32{
	0, // 0: notification.EmailService.SendEmail:input_type -> notification.EmailRequest
	0, // 1: notification.EmailService.SendResetPassWordEmail:input_type -> notification.EmailRequest
	2, // 2: notification.ChatService.CreateChat:input_type -> notification.CreateChatRequest
	4, // 3: notification.ChatService.AddMessage:input_type -> notification.AddMessageRequest
	6, // 4: notification.ChatService.EndChat:input_type -> notification.EndChatRequest
	8, // 5: notification.ChatService.StreamMessages:input_type -> notification.StreamMessagesRequest
	1, // 6: notification.EmailService.SendEmail:output_type -> notification.EmailResponse
	1, // 7: notification.EmailService.SendResetPassWordEmail:output_type -> notification.EmailResponse
	3, // 8: notification.ChatService.CreateChat:output_type -> notification.CreateChatResponse
	5, // 9: notification.ChatService.AddMessage:output_type -> notification.AddMessageResponse
	7, // 10: notification.ChatService.EndChat:output_type -> notification.EndChatResponse
	9, // 11: notification.ChatService.StreamMessages:output_type -> notification.StreamMessagesResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notification_notification_proto_init() }
func file_notification_notification_proto_init() {
	if File_notification_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notification_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_notification_notification_proto_goTypes,
		DependencyIndexes: file_notification_notification_proto_depIdxs,
		MessageInfos:      file_notification_notification_proto_msgTypes,
	}.Build()
	File_notification_notification_proto = out.File
	file_notification_notification_proto_rawDesc = nil
	file_notification_notification_proto_goTypes = nil
	file_notification_notification_proto_depIdxs = nil
}
