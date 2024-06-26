// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.0--rc1
// source: app/model/rpc/movie/movie.proto

package movie

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

type AbstractReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AbstractReq) Reset() {
	*x = AbstractReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_model_rpc_movie_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbstractReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbstractReq) ProtoMessage() {}

func (x *AbstractReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_model_rpc_movie_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbstractReq.ProtoReflect.Descriptor instead.
func (*AbstractReq) Descriptor() ([]byte, []int) {
	return file_app_model_rpc_movie_movie_proto_rawDescGZIP(), []int{0}
}

type AddMovieReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Rating      float64 `protobuf:"fixed64,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Image       string  `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Id          int64   `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddMovieReq) Reset() {
	*x = AddMovieReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_model_rpc_movie_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMovieReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMovieReq) ProtoMessage() {}

func (x *AddMovieReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_model_rpc_movie_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMovieReq.ProtoReflect.Descriptor instead.
func (*AddMovieReq) Descriptor() ([]byte, []int) {
	return file_app_model_rpc_movie_movie_proto_rawDescGZIP(), []int{1}
}

func (x *AddMovieReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddMovieReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddMovieReq) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *AddMovieReq) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *AddMovieReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddMovieRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int64      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	RefId     string     `protobuf:"bytes,2,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	Message   string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data      *MovieData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ErrorData []string   `protobuf:"bytes,5,rep,name=ErrorData,proto3" json:"ErrorData,omitempty"`
}

func (x *AddMovieRes) Reset() {
	*x = AddMovieRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_model_rpc_movie_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMovieRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMovieRes) ProtoMessage() {}

func (x *AddMovieRes) ProtoReflect() protoreflect.Message {
	mi := &file_app_model_rpc_movie_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMovieRes.ProtoReflect.Descriptor instead.
func (*AddMovieRes) Descriptor() ([]byte, []int) {
	return file_app_model_rpc_movie_movie_proto_rawDescGZIP(), []int{2}
}

func (x *AddMovieRes) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AddMovieRes) GetRefId() string {
	if x != nil {
		return x.RefId
	}
	return ""
}

func (x *AddMovieRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AddMovieRes) GetData() *MovieData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AddMovieRes) GetErrorData() []string {
	if x != nil {
		return x.ErrorData
	}
	return nil
}

type ListMovieRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	RefId   string       `protobuf:"bytes,2,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	Message string       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*MovieData `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListMovieRes) Reset() {
	*x = ListMovieRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_model_rpc_movie_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMovieRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMovieRes) ProtoMessage() {}

func (x *ListMovieRes) ProtoReflect() protoreflect.Message {
	mi := &file_app_model_rpc_movie_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMovieRes.ProtoReflect.Descriptor instead.
func (*ListMovieRes) Descriptor() ([]byte, []int) {
	return file_app_model_rpc_movie_movie_proto_rawDescGZIP(), []int{3}
}

func (x *ListMovieRes) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListMovieRes) GetRefId() string {
	if x != nil {
		return x.RefId
	}
	return ""
}

func (x *ListMovieRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListMovieRes) GetData() []*MovieData {
	if x != nil {
		return x.Data
	}
	return nil
}

type RemoveMovieRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoveMovieRes) Reset() {
	*x = RemoveMovieRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_model_rpc_movie_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveMovieRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMovieRes) ProtoMessage() {}

func (x *RemoveMovieRes) ProtoReflect() protoreflect.Message {
	mi := &file_app_model_rpc_movie_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMovieRes.ProtoReflect.Descriptor instead.
func (*RemoveMovieRes) Descriptor() ([]byte, []int) {
	return file_app_model_rpc_movie_movie_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveMovieRes) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RemoveMovieRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MovieData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Rating      float64 `protobuf:"fixed64,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Image       string  `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	CreatedAt   string  `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   string  `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *MovieData) Reset() {
	*x = MovieData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_model_rpc_movie_movie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieData) ProtoMessage() {}

func (x *MovieData) ProtoReflect() protoreflect.Message {
	mi := &file_app_model_rpc_movie_movie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieData.ProtoReflect.Descriptor instead.
func (*MovieData) Descriptor() ([]byte, []int) {
	return file_app_model_rpc_movie_movie_proto_rawDescGZIP(), []int{5}
}

func (x *MovieData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MovieData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MovieData) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *MovieData) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *MovieData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MovieData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type DetailMovieReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DetailMovieReq) Reset() {
	*x = DetailMovieReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_model_rpc_movie_movie_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailMovieReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailMovieReq) ProtoMessage() {}

func (x *DetailMovieReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_model_rpc_movie_movie_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailMovieReq.ProtoReflect.Descriptor instead.
func (*DetailMovieReq) Descriptor() ([]byte, []int) {
	return file_app_model_rpc_movie_movie_proto_rawDescGZIP(), []int{6}
}

func (x *DetailMovieReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_app_model_rpc_movie_movie_proto protoreflect.FileDescriptor

var file_app_model_rpc_movie_movie_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x62, 0x73, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x96, 0x01,
	0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x65, 0x66, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x79, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65,
	0x66, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x66, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x3e, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xbd, 0x01, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xeb, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12,
	0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0b, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x12, 0x34, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x12,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_app_model_rpc_movie_movie_proto_rawDescOnce sync.Once
	file_app_model_rpc_movie_movie_proto_rawDescData = file_app_model_rpc_movie_movie_proto_rawDesc
)

func file_app_model_rpc_movie_movie_proto_rawDescGZIP() []byte {
	file_app_model_rpc_movie_movie_proto_rawDescOnce.Do(func() {
		file_app_model_rpc_movie_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_model_rpc_movie_movie_proto_rawDescData)
	})
	return file_app_model_rpc_movie_movie_proto_rawDescData
}

var file_app_model_rpc_movie_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_app_model_rpc_movie_movie_proto_goTypes = []interface{}{
	(*AbstractReq)(nil),    // 0: model.AbstractReq
	(*AddMovieReq)(nil),    // 1: model.AddMovieReq
	(*AddMovieRes)(nil),    // 2: model.AddMovieRes
	(*ListMovieRes)(nil),   // 3: model.ListMovieRes
	(*RemoveMovieRes)(nil), // 4: model.RemoveMovieRes
	(*MovieData)(nil),      // 5: model.MovieData
	(*DetailMovieReq)(nil), // 6: model.DetailMovieReq
}
var file_app_model_rpc_movie_movie_proto_depIdxs = []int32{
	5, // 0: model.AddMovieRes.data:type_name -> model.MovieData
	5, // 1: model.ListMovieRes.data:type_name -> model.MovieData
	1, // 2: model.AuthService.AddMovie:input_type -> model.AddMovieReq
	6, // 3: model.AuthService.DetailMovie:input_type -> model.DetailMovieReq
	0, // 4: model.AuthService.ListMovie:input_type -> model.AbstractReq
	0, // 5: model.AuthService.RemoveMovie:input_type -> model.AbstractReq
	2, // 6: model.AuthService.AddMovie:output_type -> model.AddMovieRes
	2, // 7: model.AuthService.DetailMovie:output_type -> model.AddMovieRes
	3, // 8: model.AuthService.ListMovie:output_type -> model.ListMovieRes
	4, // 9: model.AuthService.RemoveMovie:output_type -> model.RemoveMovieRes
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_model_rpc_movie_movie_proto_init() }
func file_app_model_rpc_movie_movie_proto_init() {
	if File_app_model_rpc_movie_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_model_rpc_movie_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbstractReq); i {
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
		file_app_model_rpc_movie_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMovieReq); i {
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
		file_app_model_rpc_movie_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMovieRes); i {
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
		file_app_model_rpc_movie_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMovieRes); i {
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
		file_app_model_rpc_movie_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveMovieRes); i {
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
		file_app_model_rpc_movie_movie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieData); i {
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
		file_app_model_rpc_movie_movie_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailMovieReq); i {
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
			RawDescriptor: file_app_model_rpc_movie_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_model_rpc_movie_movie_proto_goTypes,
		DependencyIndexes: file_app_model_rpc_movie_movie_proto_depIdxs,
		MessageInfos:      file_app_model_rpc_movie_movie_proto_msgTypes,
	}.Build()
	File_app_model_rpc_movie_movie_proto = out.File
	file_app_model_rpc_movie_movie_proto_rawDesc = nil
	file_app_model_rpc_movie_movie_proto_goTypes = nil
	file_app_model_rpc_movie_movie_proto_depIdxs = nil
}
