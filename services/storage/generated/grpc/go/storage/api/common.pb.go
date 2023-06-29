// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: storage/api/message.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LastModified *timestamp.Timestamp `protobuf:"bytes,2,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	Size         int64                `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	StorageClass string               `protobuf:"bytes,4,opt,name=storageClass,proto3" json:"storageClass,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetLastModified() *timestamp.Timestamp {
	if x != nil {
		return x.LastModified
	}
	return nil
}

func (x *Item) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Item) GetStorageClass() string {
	if x != nil {
		return x.StorageClass
	}
	return ""
}

type GetFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileId string `protobuf:"bytes,1,opt,name=profileId,proto3" json:"profileId,omitempty"`
	AppId     string `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
}

func (x *GetFilesRequest) Reset() {
	*x = GetFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilesRequest) ProtoMessage() {}

func (x *GetFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilesRequest.ProtoReflect.Descriptor instead.
func (*GetFilesRequest) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{1}
}

func (x *GetFilesRequest) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *GetFilesRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content   []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	FileName  string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	ProfileId string `protobuf:"bytes,3,opt,name=profileId,proto3" json:"profileId,omitempty"`
	AppId     string `protobuf:"bytes,4,opt,name=appId,proto3" json:"appId,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{2}
}

func (x *UploadRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *UploadRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadRequest) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *UploadRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type GetFileContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetFileContentResponse) Reset() {
	*x = GetFileContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileContentResponse) ProtoMessage() {}

func (x *GetFileContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileContentResponse.ProtoReflect.Descriptor instead.
func (*GetFileContentResponse) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{3}
}

func (x *GetFileContentResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName  string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	ProfileId string `protobuf:"bytes,2,opt,name=profileId,proto3" json:"profileId,omitempty"`
	AppId     string `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DeleteRequest) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *DeleteRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName  string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	ProfileId string `protobuf:"bytes,2,opt,name=profileId,proto3" json:"profileId,omitempty"`
	AppId     string `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DownloadRequest) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *DownloadRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type GetACLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName  string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	ProfileId string `protobuf:"bytes,2,opt,name=profileId,proto3" json:"profileId,omitempty"`
	AppId     string `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
}

func (x *GetACLRequest) Reset() {
	*x = GetACLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetACLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetACLRequest) ProtoMessage() {}

func (x *GetACLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetACLRequest.ProtoReflect.Descriptor instead.
func (*GetACLRequest) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{6}
}

func (x *GetACLRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *GetACLRequest) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *GetACLRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type SetACLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName  string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	ProfileId string `protobuf:"bytes,2,opt,name=profileId,proto3" json:"profileId,omitempty"`
	AppId     string `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
	AclType   int32  `protobuf:"varint,4,opt,name=aclType,proto3" json:"aclType,omitempty"`
}

func (x *SetACLRequest) Reset() {
	*x = SetACLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetACLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetACLRequest) ProtoMessage() {}

func (x *SetACLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetACLRequest.ProtoReflect.Descriptor instead.
func (*SetACLRequest) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{7}
}

func (x *SetACLRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *SetACLRequest) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *SetACLRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *SetACLRequest) GetAclType() int32 {
	if x != nil {
		return x.AclType
	}
	return 0
}

type GetFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Item `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetFilesResponse) Reset() {
	*x = GetFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilesResponse) ProtoMessage() {}

func (x *GetFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilesResponse.ProtoReflect.Descriptor instead.
func (*GetFilesResponse) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{8}
}

func (x *GetFilesResponse) GetData() []*Item {
	if x != nil {
		return x.Data
	}
	return nil
}

type DownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResponse.ProtoReflect.Descriptor instead.
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{9}
}

func (x *DownloadResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ACLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grantee    string `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Permission string `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *ACLResponse) Reset() {
	*x = ACLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACLResponse) ProtoMessage() {}

func (x *ACLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACLResponse.ProtoReflect.Descriptor instead.
func (*ACLResponse) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{10}
}

func (x *ACLResponse) GetGrantee() string {
	if x != nil {
		return x.Grantee
	}
	return ""
}

func (x *ACLResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ACLResponse) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

type NothingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NothingResponse) Reset() {
	*x = NothingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_api_common_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NothingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NothingResponse) ProtoMessage() {}

func (x *NothingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_common_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NothingResponse.ProtoReflect.Descriptor instead.
func (*NothingResponse) Descriptor() ([]byte, []int) {
	return file_storage_api_common_proto_rawDescGZIP(), []int{11}
}

var File_storage_api_common_proto protoreflect.FileDescriptor

var file_storage_api_common_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x45, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x32,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x5f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x41, 0x43,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x63, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x10,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x22, 0x5b, 0x0a, 0x0b, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x11, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_api_common_proto_rawDescOnce sync.Once
	file_storage_api_common_proto_rawDescData = file_storage_api_common_proto_rawDesc
)

func file_storage_api_common_proto_rawDescGZIP() []byte {
	file_storage_api_common_proto_rawDescOnce.Do(func() {
		file_storage_api_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_api_common_proto_rawDescData)
	})
	return file_storage_api_common_proto_rawDescData
}

var file_storage_api_common_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_storage_api_common_proto_goTypes = []interface{}{
	(*Item)(nil),                   // 0: storage.pb.Item
	(*GetFilesRequest)(nil),        // 1: storage.pb.GetFilesRequest
	(*UploadRequest)(nil),          // 2: storage.pb.UploadRequest
	(*GetFileContentResponse)(nil), // 3: storage.pb.GetFileContentResponse
	(*DeleteRequest)(nil),          // 4: storage.pb.DeleteRequest
	(*DownloadRequest)(nil),        // 5: storage.pb.DownloadRequest
	(*GetACLRequest)(nil),          // 6: storage.pb.GetACLRequest
	(*SetACLRequest)(nil),          // 7: storage.pb.SetACLRequest
	(*GetFilesResponse)(nil),       // 8: storage.pb.GetFilesResponse
	(*DownloadResponse)(nil),       // 9: storage.pb.DownloadResponse
	(*ACLResponse)(nil),            // 10: storage.pb.ACLResponse
	(*NothingResponse)(nil),        // 11: storage.pb.NothingResponse
	(*timestamp.Timestamp)(nil),    // 12: google.protobuf.Timestamp
}
var file_storage_api_common_proto_depIdxs = []int32{
	12, // 0: storage.pb.Item.lastModified:type_name -> google.protobuf.Timestamp
	0,  // 1: storage.pb.GetFilesResponse.data:type_name -> storage.pb.Item
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_storage_api_common_proto_init() }
func file_storage_api_common_proto_init() {
	if File_storage_api_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_api_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_storage_api_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilesRequest); i {
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
		file_storage_api_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_storage_api_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileContentResponse); i {
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
		file_storage_api_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_storage_api_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_storage_api_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetACLRequest); i {
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
		file_storage_api_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetACLRequest); i {
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
		file_storage_api_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilesResponse); i {
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
		file_storage_api_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadResponse); i {
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
		file_storage_api_common_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACLResponse); i {
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
		file_storage_api_common_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NothingResponse); i {
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
			RawDescriptor: file_storage_api_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_api_common_proto_goTypes,
		DependencyIndexes: file_storage_api_common_proto_depIdxs,
		MessageInfos:      file_storage_api_common_proto_msgTypes,
	}.Build()
	File_storage_api_common_proto = out.File
	file_storage_api_common_proto_rawDesc = nil
	file_storage_api_common_proto_goTypes = nil
	file_storage_api_common_proto_depIdxs = nil
}
