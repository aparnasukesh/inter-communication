// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: payment/payment.proto

package payment

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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId     int32   `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	BookingId         int32   `protobuf:"varint,2,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	UserId            int32   `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PaymentMethodId   int32   `protobuf:"varint,4,opt,name=payment_method_id,json=paymentMethodId,proto3" json:"payment_method_id,omitempty"`
	TransactionDate   string  `protobuf:"bytes,5,opt,name=transaction_date,json=transactionDate,proto3" json:"transaction_date,omitempty"`
	Amount            float64 `protobuf:"fixed64,6,opt,name=amount,proto3" json:"amount,omitempty"`
	OrderId           string  `protobuf:"bytes,7,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status            string  `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	RazorpayPaymentId string  `protobuf:"bytes,9,opt,name=razorpay_payment_id,json=razorpayPaymentId,proto3" json:"razorpay_payment_id,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_payment_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetTransactionId() int32 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (x *Transaction) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *Transaction) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Transaction) GetPaymentMethodId() int32 {
	if x != nil {
		return x.PaymentMethodId
	}
	return 0
}

func (x *Transaction) GetTransactionDate() string {
	if x != nil {
		return x.TransactionDate
	}
	return ""
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Transaction) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Transaction) GetRazorpayPaymentId() string {
	if x != nil {
		return x.RazorpayPaymentId
	}
	return ""
}

type ProcessPaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId       int32   `protobuf:"varint,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	UserId          int32   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount          float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	PaymentMethodId int32   `protobuf:"varint,4,opt,name=payment_method_id,json=paymentMethodId,proto3" json:"payment_method_id,omitempty"`
}

func (x *ProcessPaymentRequest) Reset() {
	*x = ProcessPaymentRequest{}
	mi := &file_payment_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentRequest) ProtoMessage() {}

func (x *ProcessPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentRequest.ProtoReflect.Descriptor instead.
func (*ProcessPaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessPaymentRequest) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *ProcessPaymentRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ProcessPaymentRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ProcessPaymentRequest) GetPaymentMethodId() int32 {
	if x != nil {
		return x.PaymentMethodId
	}
	return 0
}

type ProcessPaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *ProcessPaymentResponse) Reset() {
	*x = ProcessPaymentResponse{}
	mi := &file_payment_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentResponse) ProtoMessage() {}

func (x *ProcessPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentResponse.ProtoReflect.Descriptor instead.
func (*ProcessPaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessPaymentResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type GetTransactionStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId int32 `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *GetTransactionStatusRequest) Reset() {
	*x = GetTransactionStatusRequest{}
	mi := &file_payment_payment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionStatusRequest) ProtoMessage() {}

func (x *GetTransactionStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionStatusRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionStatusRequest) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{3}
}

func (x *GetTransactionStatusRequest) GetTransactionId() int32 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

type GetTransactionStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId   int32   `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Status          string  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Amount          float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	PaymentMethodId int32   `protobuf:"varint,4,opt,name=payment_method_id,json=paymentMethodId,proto3" json:"payment_method_id,omitempty"`
	TransactionDate string  `protobuf:"bytes,5,opt,name=transaction_date,json=transactionDate,proto3" json:"transaction_date,omitempty"`
}

func (x *GetTransactionStatusResponse) Reset() {
	*x = GetTransactionStatusResponse{}
	mi := &file_payment_payment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionStatusResponse) ProtoMessage() {}

func (x *GetTransactionStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionStatusResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionStatusResponse) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{4}
}

func (x *GetTransactionStatusResponse) GetTransactionId() int32 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (x *GetTransactionStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetTransactionStatusResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *GetTransactionStatusResponse) GetPaymentMethodId() int32 {
	if x != nil {
		return x.PaymentMethodId
	}
	return 0
}

func (x *GetTransactionStatusResponse) GetTransactionDate() string {
	if x != nil {
		return x.TransactionDate
	}
	return ""
}

type PaymentSuccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId           string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	RazorpayPaymentId string `protobuf:"bytes,2,opt,name=razorpay_payment_id,json=razorpayPaymentId,proto3" json:"razorpay_payment_id,omitempty"`
	BookingId         int32  `protobuf:"varint,3,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
}

func (x *PaymentSuccessRequest) Reset() {
	*x = PaymentSuccessRequest{}
	mi := &file_payment_payment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentSuccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentSuccessRequest) ProtoMessage() {}

func (x *PaymentSuccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentSuccessRequest.ProtoReflect.Descriptor instead.
func (*PaymentSuccessRequest) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{5}
}

func (x *PaymentSuccessRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *PaymentSuccessRequest) GetRazorpayPaymentId() string {
	if x != nil {
		return x.RazorpayPaymentId
	}
	return ""
}

func (x *PaymentSuccessRequest) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

type PaymentSuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PaymentSuccessResponse) Reset() {
	*x = PaymentSuccessResponse{}
	mi := &file_payment_payment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentSuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentSuccessResponse) ProtoMessage() {}

func (x *PaymentSuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentSuccessResponse.ProtoReflect.Descriptor instead.
func (*PaymentSuccessResponse) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{6}
}

type PaymentFailureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId           string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	RazorpayPaymentId string `protobuf:"bytes,2,opt,name=razorpay_payment_id,json=razorpayPaymentId,proto3" json:"razorpay_payment_id,omitempty"`
	BookingId         int32  `protobuf:"varint,3,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
}

func (x *PaymentFailureRequest) Reset() {
	*x = PaymentFailureRequest{}
	mi := &file_payment_payment_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentFailureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentFailureRequest) ProtoMessage() {}

func (x *PaymentFailureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentFailureRequest.ProtoReflect.Descriptor instead.
func (*PaymentFailureRequest) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{7}
}

func (x *PaymentFailureRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *PaymentFailureRequest) GetRazorpayPaymentId() string {
	if x != nil {
		return x.RazorpayPaymentId
	}
	return ""
}

func (x *PaymentFailureRequest) GetBookingId() int32 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

type PaymentFailureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PaymentFailureResponse) Reset() {
	*x = PaymentFailureResponse{}
	mi := &file_payment_payment_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentFailureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentFailureResponse) ProtoMessage() {}

func (x *PaymentFailureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_payment_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentFailureResponse.ProtoReflect.Descriptor instead.
func (*PaymentFailureResponse) Descriptor() ([]byte, []int) {
	return file_payment_payment_proto_rawDescGZIP(), []int{8}
}

var File_payment_payment_proto protoreflect.FileDescriptor

var file_payment_payment_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0xbe, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x61, 0x7a, 0x6f, 0x72, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x72, 0x61, 0x7a, 0x6f, 0x72, 0x70, 0x61, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x93, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x44, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0xcc, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x81,
	0x01, 0x0a, 0x15, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x61, 0x7a, 0x6f, 0x72, 0x70, 0x61, 0x79, 0x5f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x72, 0x61, 0x7a, 0x6f, 0x72, 0x70, 0x61, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x81, 0x01, 0x0a,
	0x15, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x61, 0x7a, 0x6f, 0x72, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x72, 0x61, 0x7a, 0x6f, 0x72, 0x70, 0x61, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x22, 0x18, 0x0a, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xee, 0x02, 0x0a, 0x0e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a,
	0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x63, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_payment_proto_rawDescOnce sync.Once
	file_payment_payment_proto_rawDescData = file_payment_payment_proto_rawDesc
)

func file_payment_payment_proto_rawDescGZIP() []byte {
	file_payment_payment_proto_rawDescOnce.Do(func() {
		file_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_payment_proto_rawDescData)
	})
	return file_payment_payment_proto_rawDescData
}

var file_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_payment_payment_proto_goTypes = []any{
	(*Transaction)(nil),                  // 0: payment.Transaction
	(*ProcessPaymentRequest)(nil),        // 1: payment.ProcessPaymentRequest
	(*ProcessPaymentResponse)(nil),       // 2: payment.ProcessPaymentResponse
	(*GetTransactionStatusRequest)(nil),  // 3: payment.GetTransactionStatusRequest
	(*GetTransactionStatusResponse)(nil), // 4: payment.GetTransactionStatusResponse
	(*PaymentSuccessRequest)(nil),        // 5: payment.PaymentSuccessRequest
	(*PaymentSuccessResponse)(nil),       // 6: payment.PaymentSuccessResponse
	(*PaymentFailureRequest)(nil),        // 7: payment.PaymentFailureRequest
	(*PaymentFailureResponse)(nil),       // 8: payment.PaymentFailureResponse
}
var file_payment_payment_proto_depIdxs = []int32{
	0, // 0: payment.ProcessPaymentResponse.transaction:type_name -> payment.Transaction
	1, // 1: payment.PaymentService.ProcessPayment:input_type -> payment.ProcessPaymentRequest
	3, // 2: payment.PaymentService.GetTransactionStatus:input_type -> payment.GetTransactionStatusRequest
	5, // 3: payment.PaymentService.PaymentSuccess:input_type -> payment.PaymentSuccessRequest
	7, // 4: payment.PaymentService.PaymentFailure:input_type -> payment.PaymentFailureRequest
	2, // 5: payment.PaymentService.ProcessPayment:output_type -> payment.ProcessPaymentResponse
	4, // 6: payment.PaymentService.GetTransactionStatus:output_type -> payment.GetTransactionStatusResponse
	6, // 7: payment.PaymentService.PaymentSuccess:output_type -> payment.PaymentSuccessResponse
	8, // 8: payment.PaymentService.PaymentFailure:output_type -> payment.PaymentFailureResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payment_payment_proto_init() }
func file_payment_payment_proto_init() {
	if File_payment_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payment_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_payment_proto_goTypes,
		DependencyIndexes: file_payment_payment_proto_depIdxs,
		MessageInfos:      file_payment_payment_proto_msgTypes,
	}.Build()
	File_payment_payment_proto = out.File
	file_payment_payment_proto_rawDesc = nil
	file_payment_payment_proto_goTypes = nil
	file_payment_payment_proto_depIdxs = nil
}
