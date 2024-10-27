// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: booking.proto

package booking_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ApartmentId int64                  `protobuf:"varint,2,opt,name=apartment_id,json=apartmentId,proto3" json:"apartment_id,omitempty"`
	DateStart   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateEnd     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
	Price       int64                  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	CustomerId  int64                  `protobuf:"varint,6,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Status      string                 `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Comment     string                 `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	DateCreated *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	mi := &file_booking_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

func (x *Booking) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Booking) GetApartmentId() int64 {
	if x != nil {
		return x.ApartmentId
	}
	return 0
}

func (x *Booking) GetDateStart() *timestamppb.Timestamp {
	if x != nil {
		return x.DateStart
	}
	return nil
}

func (x *Booking) GetDateEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.DateEnd
	}
	return nil
}

func (x *Booking) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Booking) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Booking) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Booking) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Booking) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

type BeginBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BeginBookingRequest) Reset() {
	*x = BeginBookingRequest{}
	mi := &file_booking_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BeginBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginBookingRequest) ProtoMessage() {}

func (x *BeginBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginBookingRequest.ProtoReflect.Descriptor instead.
func (*BeginBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1}
}

func (x *BeginBookingRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BeginBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateStart *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	Id        int64                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BeginBookingResponse) Reset() {
	*x = BeginBookingResponse{}
	mi := &file_booking_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BeginBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginBookingResponse) ProtoMessage() {}

func (x *BeginBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginBookingResponse.ProtoReflect.Descriptor instead.
func (*BeginBookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{2}
}

func (x *BeginBookingResponse) GetDateStart() *timestamppb.Timestamp {
	if x != nil {
		return x.DateStart
	}
	return nil
}

func (x *BeginBookingResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FinishBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FinishBookingRequest) Reset() {
	*x = FinishBookingRequest{}
	mi := &file_booking_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FinishBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishBookingRequest) ProtoMessage() {}

func (x *FinishBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishBookingRequest.ProtoReflect.Descriptor instead.
func (*FinishBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{3}
}

func (x *FinishBookingRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FinishBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateEnd *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
	Id      int64                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FinishBookingResponse) Reset() {
	*x = FinishBookingResponse{}
	mi := &file_booking_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FinishBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishBookingResponse) ProtoMessage() {}

func (x *FinishBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishBookingResponse.ProtoReflect.Descriptor instead.
func (*FinishBookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{4}
}

func (x *FinishBookingResponse) GetDateEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.DateEnd
	}
	return nil
}

func (x *FinishBookingResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NewBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApartmentId int64  `protobuf:"varint,1,opt,name=apartment_id,json=apartmentId,proto3" json:"apartment_id,omitempty"`
	Price       int64  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	CustomerId  int64  `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Comment     string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *NewBookingRequest) Reset() {
	*x = NewBookingRequest{}
	mi := &file_booking_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBookingRequest) ProtoMessage() {}

func (x *NewBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBookingRequest.ProtoReflect.Descriptor instead.
func (*NewBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{5}
}

func (x *NewBookingRequest) GetApartmentId() int64 {
	if x != nil {
		return x.ApartmentId
	}
	return 0
}

func (x *NewBookingRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *NewBookingRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *NewBookingRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type UpdateBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ApartmentId *wrapperspb.Int64Value  `protobuf:"bytes,2,opt,name=apartment_id,json=apartmentId,proto3" json:"apartment_id,omitempty"`
	Price       *wrapperspb.Int64Value  `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	CustomerId  *wrapperspb.Int64Value  `protobuf:"bytes,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	DateStart   *timestamppb.Timestamp  `protobuf:"bytes,5,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateEnd     *timestamppb.Timestamp  `protobuf:"bytes,6,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
	Comment     *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *UpdateBookingRequest) Reset() {
	*x = UpdateBookingRequest{}
	mi := &file_booking_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookingRequest) ProtoMessage() {}

func (x *UpdateBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookingRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBookingRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateBookingRequest) GetApartmentId() *wrapperspb.Int64Value {
	if x != nil {
		return x.ApartmentId
	}
	return nil
}

func (x *UpdateBookingRequest) GetPrice() *wrapperspb.Int64Value {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *UpdateBookingRequest) GetCustomerId() *wrapperspb.Int64Value {
	if x != nil {
		return x.CustomerId
	}
	return nil
}

func (x *UpdateBookingRequest) GetDateStart() *timestamppb.Timestamp {
	if x != nil {
		return x.DateStart
	}
	return nil
}

func (x *UpdateBookingRequest) GetDateEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.DateEnd
	}
	return nil
}

func (x *UpdateBookingRequest) GetComment() *wrapperspb.StringValue {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CancelBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelBookingRequest) Reset() {
	*x = CancelBookingRequest{}
	mi := &file_booking_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBookingRequest) ProtoMessage() {}

func (x *CancelBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBookingRequest.ProtoReflect.Descriptor instead.
func (*CancelBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{7}
}

func (x *CancelBookingRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CancelBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CancelBookingResponse) Reset() {
	*x = CancelBookingResponse{}
	mi := &file_booking_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBookingResponse) ProtoMessage() {}

func (x *CancelBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBookingResponse.ProtoReflect.Descriptor instead.
func (*CancelBookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{8}
}

func (x *CancelBookingResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_booking_proto protoreflect.FileDescriptor

var file_booking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x02, 0x0a,
	0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65,
	0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x14,
	0x42, 0x65, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x26, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x81, 0x03, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x0c, 0x61, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x61,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65,
	0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x36, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a,
	0x15, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0x95, 0x03, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x4e,
	0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22,
	0x00, 0x12, 0x53, 0x0a, 0x0c, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x1f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x42,
	0x65, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x2e,
	0x42, 0x65, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x20, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booking_proto_rawDescOnce sync.Once
	file_booking_proto_rawDescData = file_booking_proto_rawDesc
)

func file_booking_proto_rawDescGZIP() []byte {
	file_booking_proto_rawDescOnce.Do(func() {
		file_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_proto_rawDescData)
	})
	return file_booking_proto_rawDescData
}

var file_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_booking_proto_goTypes = []any{
	(*Booking)(nil),                // 0: booking_v1.Booking
	(*BeginBookingRequest)(nil),    // 1: booking_v1.BeginBookingRequest
	(*BeginBookingResponse)(nil),   // 2: booking_v1.BeginBookingResponse
	(*FinishBookingRequest)(nil),   // 3: booking_v1.FinishBookingRequest
	(*FinishBookingResponse)(nil),  // 4: booking_v1.FinishBookingResponse
	(*NewBookingRequest)(nil),      // 5: booking_v1.NewBookingRequest
	(*UpdateBookingRequest)(nil),   // 6: booking_v1.UpdateBookingRequest
	(*CancelBookingRequest)(nil),   // 7: booking_v1.CancelBookingRequest
	(*CancelBookingResponse)(nil),  // 8: booking_v1.CancelBookingResponse
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
	(*wrapperspb.Int64Value)(nil),  // 10: google.protobuf.Int64Value
	(*wrapperspb.StringValue)(nil), // 11: google.protobuf.StringValue
}
var file_booking_proto_depIdxs = []int32{
	9,  // 0: booking_v1.Booking.date_start:type_name -> google.protobuf.Timestamp
	9,  // 1: booking_v1.Booking.date_end:type_name -> google.protobuf.Timestamp
	9,  // 2: booking_v1.Booking.date_created:type_name -> google.protobuf.Timestamp
	9,  // 3: booking_v1.BeginBookingResponse.date_start:type_name -> google.protobuf.Timestamp
	9,  // 4: booking_v1.FinishBookingResponse.date_end:type_name -> google.protobuf.Timestamp
	10, // 5: booking_v1.UpdateBookingRequest.apartment_id:type_name -> google.protobuf.Int64Value
	10, // 6: booking_v1.UpdateBookingRequest.price:type_name -> google.protobuf.Int64Value
	10, // 7: booking_v1.UpdateBookingRequest.customer_id:type_name -> google.protobuf.Int64Value
	9,  // 8: booking_v1.UpdateBookingRequest.date_start:type_name -> google.protobuf.Timestamp
	9,  // 9: booking_v1.UpdateBookingRequest.date_end:type_name -> google.protobuf.Timestamp
	11, // 10: booking_v1.UpdateBookingRequest.comment:type_name -> google.protobuf.StringValue
	5,  // 11: booking_v1.BookingService.NewBooking:input_type -> booking_v1.NewBookingRequest
	6,  // 12: booking_v1.BookingService.UpdateBooking:input_type -> booking_v1.UpdateBookingRequest
	1,  // 13: booking_v1.BookingService.BeginBooking:input_type -> booking_v1.BeginBookingRequest
	3,  // 14: booking_v1.BookingService.FinishBooking:input_type -> booking_v1.FinishBookingRequest
	7,  // 15: booking_v1.BookingService.CancelBooking:input_type -> booking_v1.CancelBookingRequest
	0,  // 16: booking_v1.BookingService.NewBooking:output_type -> booking_v1.Booking
	0,  // 17: booking_v1.BookingService.UpdateBooking:output_type -> booking_v1.Booking
	2,  // 18: booking_v1.BookingService.BeginBooking:output_type -> booking_v1.BeginBookingResponse
	0,  // 19: booking_v1.BookingService.FinishBooking:output_type -> booking_v1.Booking
	8,  // 20: booking_v1.BookingService.CancelBooking:output_type -> booking_v1.CancelBookingResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_booking_proto_init() }
func file_booking_proto_init() {
	if File_booking_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_proto_goTypes,
		DependencyIndexes: file_booking_proto_depIdxs,
		MessageInfos:      file_booking_proto_msgTypes,
	}.Build()
	File_booking_proto = out.File
	file_booking_proto_rawDesc = nil
	file_booking_proto_goTypes = nil
	file_booking_proto_depIdxs = nil
}
