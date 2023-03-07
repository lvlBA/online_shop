// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: resource.proto

package passport

import (
	paspport "github.com/lvlBA/online_shop/paspport"
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

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Urn string `protobuf:"bytes,2,opt,name=urn,proto3" json:"urn,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resource) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type CreateResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urn string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
}

func (x *CreateResourceRequest) Reset() {
	*x = CreateResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceRequest) ProtoMessage() {}

func (x *CreateResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceRequest.ProtoReflect.Descriptor instead.
func (*CreateResourceRequest) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResourceRequest) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type CreateResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *CreateResourceResponse) Reset() {
	*x = CreateResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceResponse) ProtoMessage() {}

func (x *CreateResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceResponse.ProtoReflect.Descriptor instead.
func (*CreateResourceResponse) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResourceResponse) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type GetResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetResourceRequest) Reset() {
	*x = GetResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceRequest) ProtoMessage() {}

func (x *GetResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceRequest.ProtoReflect.Descriptor instead.
func (*GetResourceRequest) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{3}
}

func (x *GetResourceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *GetResourceResponse) Reset() {
	*x = GetResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceResponse) ProtoMessage() {}

func (x *GetResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceResponse.ProtoReflect.Descriptor instead.
func (*GetResourceResponse) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{4}
}

func (x *GetResourceResponse) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type DeleteResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteResourceRequest) Reset() {
	*x = DeleteResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceRequest) ProtoMessage() {}

func (x *DeleteResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceRequest.ProtoReflect.Descriptor instead.
func (*DeleteResourceRequest) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteResourceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResourceResponse) Reset() {
	*x = DeleteResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceResponse) ProtoMessage() {}

func (x *DeleteResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceResponse.ProtoReflect.Descriptor instead.
func (*DeleteResourceResponse) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{6}
}

type ListResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *paspport.Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListResourceRequest) Reset() {
	*x = ListResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourceRequest) ProtoMessage() {}

func (x *ListResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourceRequest.ProtoReflect.Descriptor instead.
func (*ListResourceRequest) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{7}
}

func (x *ListResourceRequest) GetPagination() *paspport.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListResourceResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource []*Resource `protobuf:"bytes,1,rep,name=resource,proto3" json:"resource,omitempty"`
}

func (x *ListResourceResponce) Reset() {
	*x = ListResourceResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourceResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourceResponce) ProtoMessage() {}

func (x *ListResourceResponce) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourceResponce.ProtoReflect.Descriptor instead.
func (*ListResourceResponce) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{8}
}

func (x *ListResourceResponce) GetResource() []*Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type SetUserAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *SetUserAccessRequest) Reset() {
	*x = SetUserAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserAccessRequest) ProtoMessage() {}

func (x *SetUserAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserAccessRequest.ProtoReflect.Descriptor instead.
func (*SetUserAccessRequest) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{9}
}

func (x *SetUserAccessRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SetUserAccessRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type SetUserAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetUserAccessResponse) Reset() {
	*x = SetUserAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserAccessResponse) ProtoMessage() {}

func (x *SetUserAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserAccessResponse.ProtoReflect.Descriptor instead.
func (*SetUserAccessResponse) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{10}
}

type DeleteUserAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeleteUserAccessRequest) Reset() {
	*x = DeleteUserAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserAccessRequest) ProtoMessage() {}

func (x *DeleteUserAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserAccessRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserAccessRequest) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteUserAccessRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteUserAccessRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type DeleteUserAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserAccessResponse) Reset() {
	*x = DeleteUserAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserAccessResponse) ProtoMessage() {}

func (x *DeleteUserAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resource_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserAccessResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserAccessResponse) Descriptor() ([]byte, []int) {
	return file_resource_proto_rawDescGZIP(), []int{12}
}

var File_resource_proto protoreflect.FileDescriptor

var file_resource_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61,
	0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x22, 0x29, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6e, 0x22, 0x57, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70,
	0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x24, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x3d, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0x50, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a, 0x17, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22,
	0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc3, 0x05, 0x0a, 0x0f,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x73, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x2e, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2f, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70,
	0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x73, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x2e, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2c, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x30, 0x2e, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x76, 0x6c, 0x42, 0x41, 0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f,
	0x70, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x3b, 0x70, 0x61, 0x73, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_proto_rawDescOnce sync.Once
	file_resource_proto_rawDescData = file_resource_proto_rawDesc
)

func file_resource_proto_rawDescGZIP() []byte {
	file_resource_proto_rawDescOnce.Do(func() {
		file_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_proto_rawDescData)
	})
	return file_resource_proto_rawDescData
}

var file_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_resource_proto_goTypes = []interface{}{
	(*Resource)(nil),                 // 0: online_shop.passport.v1.Resource
	(*CreateResourceRequest)(nil),    // 1: online_shop.passport.v1.CreateResourceRequest
	(*CreateResourceResponse)(nil),   // 2: online_shop.passport.v1.CreateResourceResponse
	(*GetResourceRequest)(nil),       // 3: online_shop.passport.v1.GetResourceRequest
	(*GetResourceResponse)(nil),      // 4: online_shop.passport.v1.GetResourceResponse
	(*DeleteResourceRequest)(nil),    // 5: online_shop.passport.v1.DeleteResourceRequest
	(*DeleteResourceResponse)(nil),   // 6: online_shop.passport.v1.DeleteResourceResponse
	(*ListResourceRequest)(nil),      // 7: online_shop.passport.v1.ListResourceRequest
	(*ListResourceResponce)(nil),     // 8: online_shop.passport.v1.ListResourceResponce
	(*SetUserAccessRequest)(nil),     // 9: online_shop.passport.v1.SetUserAccessRequest
	(*SetUserAccessResponse)(nil),    // 10: online_shop.passport.v1.SetUserAccessResponse
	(*DeleteUserAccessRequest)(nil),  // 11: online_shop.passport.v1.DeleteUserAccessRequest
	(*DeleteUserAccessResponse)(nil), // 12: online_shop.passport.v1.DeleteUserAccessResponse
	(*paspport.Pagination)(nil),      // 13: online_shop.passport.v1.Pagination
}
var file_resource_proto_depIdxs = []int32{
	0,  // 0: online_shop.passport.v1.CreateResourceResponse.resource:type_name -> online_shop.passport.v1.Resource
	0,  // 1: online_shop.passport.v1.GetResourceResponse.resource:type_name -> online_shop.passport.v1.Resource
	13, // 2: online_shop.passport.v1.ListResourceRequest.pagination:type_name -> online_shop.passport.v1.Pagination
	0,  // 3: online_shop.passport.v1.ListResourceResponce.resource:type_name -> online_shop.passport.v1.Resource
	1,  // 4: online_shop.passport.v1.ResourceService.CreateResource:input_type -> online_shop.passport.v1.CreateResourceRequest
	3,  // 5: online_shop.passport.v1.ResourceService.GetResource:input_type -> online_shop.passport.v1.GetResourceRequest
	5,  // 6: online_shop.passport.v1.ResourceService.DeleteResource:input_type -> online_shop.passport.v1.DeleteResourceRequest
	7,  // 7: online_shop.passport.v1.ResourceService.ListResource:input_type -> online_shop.passport.v1.ListResourceRequest
	9,  // 8: online_shop.passport.v1.ResourceService.SetUserAccess:input_type -> online_shop.passport.v1.SetUserAccessRequest
	11, // 9: online_shop.passport.v1.ResourceService.DeleteUserAccess:input_type -> online_shop.passport.v1.DeleteUserAccessRequest
	2,  // 10: online_shop.passport.v1.ResourceService.CreateResource:output_type -> online_shop.passport.v1.CreateResourceResponse
	4,  // 11: online_shop.passport.v1.ResourceService.GetResource:output_type -> online_shop.passport.v1.GetResourceResponse
	6,  // 12: online_shop.passport.v1.ResourceService.DeleteResource:output_type -> online_shop.passport.v1.DeleteResourceResponse
	8,  // 13: online_shop.passport.v1.ResourceService.ListResource:output_type -> online_shop.passport.v1.ListResourceResponce
	10, // 14: online_shop.passport.v1.ResourceService.SetUserAccess:output_type -> online_shop.passport.v1.SetUserAccessResponse
	12, // 15: online_shop.passport.v1.ResourceService.DeleteUserAccess:output_type -> online_shop.passport.v1.DeleteUserAccessResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_resource_proto_init() }
func file_resource_proto_init() {
	if File_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceRequest); i {
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
		file_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceResponse); i {
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
		file_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceRequest); i {
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
		file_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceResponse); i {
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
		file_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResourceRequest); i {
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
		file_resource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResourceResponse); i {
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
		file_resource_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourceRequest); i {
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
		file_resource_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourceResponce); i {
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
		file_resource_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserAccessRequest); i {
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
		file_resource_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserAccessResponse); i {
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
		file_resource_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserAccessRequest); i {
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
		file_resource_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserAccessResponse); i {
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
			RawDescriptor: file_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resource_proto_goTypes,
		DependencyIndexes: file_resource_proto_depIdxs,
		MessageInfos:      file_resource_proto_msgTypes,
	}.Build()
	File_resource_proto = out.File
	file_resource_proto_rawDesc = nil
	file_resource_proto_goTypes = nil
	file_resource_proto_depIdxs = nil
}
