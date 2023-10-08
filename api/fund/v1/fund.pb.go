// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: api/fund/v1/fund.proto

package v1

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

type CreateFundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateFundRequest) Reset() {
	*x = CreateFundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fund_v1_fund_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFundRequest) ProtoMessage() {}

func (x *CreateFundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_fund_v1_fund_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFundRequest.ProtoReflect.Descriptor instead.
func (*CreateFundRequest) Descriptor() ([]byte, []int) {
	return file_api_fund_v1_fund_proto_rawDescGZIP(), []int{0}
}

type CreateFundReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateFundReply) Reset() {
	*x = CreateFundReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fund_v1_fund_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFundReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFundReply) ProtoMessage() {}

func (x *CreateFundReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_fund_v1_fund_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFundReply.ProtoReflect.Descriptor instead.
func (*CreateFundReply) Descriptor() ([]byte, []int) {
	return file_api_fund_v1_fund_proto_rawDescGZIP(), []int{1}
}

type UpdateFundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateFundRequest) Reset() {
	*x = UpdateFundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fund_v1_fund_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFundRequest) ProtoMessage() {}

func (x *UpdateFundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_fund_v1_fund_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFundRequest.ProtoReflect.Descriptor instead.
func (*UpdateFundRequest) Descriptor() ([]byte, []int) {
	return file_api_fund_v1_fund_proto_rawDescGZIP(), []int{2}
}

type UpdateFundReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateFundReply) Reset() {
	*x = UpdateFundReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fund_v1_fund_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFundReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFundReply) ProtoMessage() {}

func (x *UpdateFundReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_fund_v1_fund_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFundReply.ProtoReflect.Descriptor instead.
func (*UpdateFundReply) Descriptor() ([]byte, []int) {
	return file_api_fund_v1_fund_proto_rawDescGZIP(), []int{3}
}

type DeleteFundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteFundRequest) Reset() {
	*x = DeleteFundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fund_v1_fund_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFundRequest) ProtoMessage() {}

func (x *DeleteFundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_fund_v1_fund_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFundRequest.ProtoReflect.Descriptor instead.
func (*DeleteFundRequest) Descriptor() ([]byte, []int) {
	return file_api_fund_v1_fund_proto_rawDescGZIP(), []int{4}
}

type DeleteFundReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteFundReply) Reset() {
	*x = DeleteFundReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fund_v1_fund_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFundReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFundReply) ProtoMessage() {}

func (x *DeleteFundReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_fund_v1_fund_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFundReply.ProtoReflect.Descriptor instead.
func (*DeleteFundReply) Descriptor() ([]byte, []int) {
	return file_api_fund_v1_fund_proto_rawDescGZIP(), []int{5}
}

type GetFundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFundRequest) Reset() {
	*x = GetFundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fund_v1_fund_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFundRequest) ProtoMessage() {}

func (x *GetFundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_fund_v1_fund_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFundRequest.ProtoReflect.Descriptor instead.
func (*GetFundRequest) Descriptor() ([]byte, []int) {
	return file_api_fund_v1_fund_proto_rawDescGZIP(), []int{6}
}

type GetFundReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFundReply) Reset() {
	*x = GetFundReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fund_v1_fund_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFundReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFundReply) ProtoMessage() {}

func (x *GetFundReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_fund_v1_fund_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFundReply.ProtoReflect.Descriptor instead.
func (*GetFundReply) Descriptor() ([]byte, []int) {
	return file_api_fund_v1_fund_proto_rawDescGZIP(), []int{7}
}

type ListFundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFundRequest) Reset() {
	*x = ListFundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fund_v1_fund_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFundRequest) ProtoMessage() {}

func (x *ListFundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_fund_v1_fund_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFundRequest.ProtoReflect.Descriptor instead.
func (*ListFundRequest) Descriptor() ([]byte, []int) {
	return file_api_fund_v1_fund_proto_rawDescGZIP(), []int{8}
}

type ListFundReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFundReply) Reset() {
	*x = ListFundReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_fund_v1_fund_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFundReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFundReply) ProtoMessage() {}

func (x *ListFundReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_fund_v1_fund_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFundReply.ProtoReflect.Descriptor instead.
func (*ListFundReply) Descriptor() ([]byte, []int) {
	return file_api_fund_v1_fund_proto_rawDescGZIP(), []int{9}
}

var File_api_fund_v1_fund_proto protoreflect.FileDescriptor

var file_api_fund_v1_fund_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75,
	0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x75,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0xf3, 0x02, 0x0a, 0x04, 0x46, 0x75, 0x6e, 0x64, 0x12, 0x4a, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x75, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x75, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x75,
	0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66,
	0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x12,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x27, 0x0a, 0x0b, 0x61, 0x70, 0x69,
	0x2e, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x16, 0x74, 0x61, 0x6f, 0x2d,
	0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_fund_v1_fund_proto_rawDescOnce sync.Once
	file_api_fund_v1_fund_proto_rawDescData = file_api_fund_v1_fund_proto_rawDesc
)

func file_api_fund_v1_fund_proto_rawDescGZIP() []byte {
	file_api_fund_v1_fund_proto_rawDescOnce.Do(func() {
		file_api_fund_v1_fund_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_fund_v1_fund_proto_rawDescData)
	})
	return file_api_fund_v1_fund_proto_rawDescData
}

var file_api_fund_v1_fund_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_fund_v1_fund_proto_goTypes = []interface{}{
	(*CreateFundRequest)(nil), // 0: api.fund.v1.CreateFundRequest
	(*CreateFundReply)(nil),   // 1: api.fund.v1.CreateFundReply
	(*UpdateFundRequest)(nil), // 2: api.fund.v1.UpdateFundRequest
	(*UpdateFundReply)(nil),   // 3: api.fund.v1.UpdateFundReply
	(*DeleteFundRequest)(nil), // 4: api.fund.v1.DeleteFundRequest
	(*DeleteFundReply)(nil),   // 5: api.fund.v1.DeleteFundReply
	(*GetFundRequest)(nil),    // 6: api.fund.v1.GetFundRequest
	(*GetFundReply)(nil),      // 7: api.fund.v1.GetFundReply
	(*ListFundRequest)(nil),   // 8: api.fund.v1.ListFundRequest
	(*ListFundReply)(nil),     // 9: api.fund.v1.ListFundReply
}
var file_api_fund_v1_fund_proto_depIdxs = []int32{
	0, // 0: api.fund.v1.Fund.CreateFund:input_type -> api.fund.v1.CreateFundRequest
	2, // 1: api.fund.v1.Fund.UpdateFund:input_type -> api.fund.v1.UpdateFundRequest
	4, // 2: api.fund.v1.Fund.DeleteFund:input_type -> api.fund.v1.DeleteFundRequest
	6, // 3: api.fund.v1.Fund.GetFund:input_type -> api.fund.v1.GetFundRequest
	8, // 4: api.fund.v1.Fund.ListFund:input_type -> api.fund.v1.ListFundRequest
	1, // 5: api.fund.v1.Fund.CreateFund:output_type -> api.fund.v1.CreateFundReply
	3, // 6: api.fund.v1.Fund.UpdateFund:output_type -> api.fund.v1.UpdateFundReply
	5, // 7: api.fund.v1.Fund.DeleteFund:output_type -> api.fund.v1.DeleteFundReply
	7, // 8: api.fund.v1.Fund.GetFund:output_type -> api.fund.v1.GetFundReply
	9, // 9: api.fund.v1.Fund.ListFund:output_type -> api.fund.v1.ListFundReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_fund_v1_fund_proto_init() }
func file_api_fund_v1_fund_proto_init() {
	if File_api_fund_v1_fund_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_fund_v1_fund_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFundRequest); i {
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
		file_api_fund_v1_fund_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFundReply); i {
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
		file_api_fund_v1_fund_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFundRequest); i {
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
		file_api_fund_v1_fund_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFundReply); i {
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
		file_api_fund_v1_fund_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFundRequest); i {
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
		file_api_fund_v1_fund_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFundReply); i {
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
		file_api_fund_v1_fund_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFundRequest); i {
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
		file_api_fund_v1_fund_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFundReply); i {
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
		file_api_fund_v1_fund_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFundRequest); i {
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
		file_api_fund_v1_fund_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFundReply); i {
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
			RawDescriptor: file_api_fund_v1_fund_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_fund_v1_fund_proto_goTypes,
		DependencyIndexes: file_api_fund_v1_fund_proto_depIdxs,
		MessageInfos:      file_api_fund_v1_fund_proto_msgTypes,
	}.Build()
	File_api_fund_v1_fund_proto = out.File
	file_api_fund_v1_fund_proto_rawDesc = nil
	file_api_fund_v1_fund_proto_goTypes = nil
	file_api_fund_v1_fund_proto_depIdxs = nil
}