// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: message.proto

package pb

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

type PlatformType_Type int32

const (
	PlatformType_Unknown  PlatformType_Type = 0
	PlatformType_Facebook PlatformType_Type = 1
	PlatformType_Google   PlatformType_Type = 2
	PlatformType_NineYou  PlatformType_Type = 3 //久游
)

// Enum value maps for PlatformType_Type.
var (
	PlatformType_Type_name = map[int32]string{
		0: "Unknown",
		1: "Facebook",
		2: "Google",
		3: "NineYou",
	}
	PlatformType_Type_value = map[string]int32{
		"Unknown":  0,
		"Facebook": 1,
		"Google":   2,
		"NineYou":  3,
	}
)

func (x PlatformType_Type) Enum() *PlatformType_Type {
	p := new(PlatformType_Type)
	*p = x
	return p
}

func (x PlatformType_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlatformType_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (PlatformType_Type) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x PlatformType_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlatformType_Type.Descriptor instead.
func (PlatformType_Type) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0, 0}
}

type PlatformType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlatformType) Reset() {
	*x = PlatformType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformType) ProtoMessage() {}

func (x *PlatformType) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformType.ProtoReflect.Descriptor instead.
func (*PlatformType) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	AppId  string `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId,omitempty"`
	UserId string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Session) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Session) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *Session) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type BindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	// Types that are assignable to Token:
	//
	//	*BindRequest_FacebookToken
	//	*BindRequest_GoogleToken
	Token isBindRequest_Token `protobuf_oneof:"Token"`
}

func (x *BindRequest) Reset() {
	*x = BindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindRequest) ProtoMessage() {}

func (x *BindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindRequest.ProtoReflect.Descriptor instead.
func (*BindRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *BindRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (m *BindRequest) GetToken() isBindRequest_Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (x *BindRequest) GetFacebookToken() string {
	if x, ok := x.GetToken().(*BindRequest_FacebookToken); ok {
		return x.FacebookToken
	}
	return ""
}

func (x *BindRequest) GetGoogleToken() string {
	if x, ok := x.GetToken().(*BindRequest_GoogleToken); ok {
		return x.GoogleToken
	}
	return ""
}

type isBindRequest_Token interface {
	isBindRequest_Token()
}

type BindRequest_FacebookToken struct {
	FacebookToken string `protobuf:"bytes,2,opt,name=facebookToken,proto3,oneof"`
}

type BindRequest_GoogleToken struct {
	GoogleToken string `protobuf:"bytes,3,opt,name=googleToken,proto3,oneof"`
}

func (*BindRequest_FacebookToken) isBindRequest_Token() {}

func (*BindRequest_GoogleToken) isBindRequest_Token() {}

type BindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	IsReLogin bool   `protobuf:"varint,2,opt,name=isReLogin,proto3" json:"isReLogin,omitempty"`
}

func (x *BindResponse) Reset() {
	*x = BindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindResponse) ProtoMessage() {}

func (x *BindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindResponse.ProtoReflect.Descriptor instead.
func (*BindResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *BindResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *BindResponse) GetIsReLogin() bool {
	if x != nil {
		return x.IsReLogin
	}
	return false
}

type NothingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NothingResponse) Reset() {
	*x = NothingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NothingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NothingResponse) ProtoMessage() {}

func (x *NothingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
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
	return file_message_proto_rawDescGZIP(), []int{4}
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x62, 0x22, 0x4a, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x46, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x69, 0x6e, 0x65, 0x59,
	0x6f, 0x75, 0x10, 0x03, 0x22, 0x4d, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0d, 0x66, 0x61, 0x63, 0x65,
	0x62, 0x6f, 0x6f, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0d, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x22, 0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3e, 0x0a,
	0x0c, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x52, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x11, 0x0a,
	0x0f, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0d, 0x5a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_proto_goTypes = []interface{}{
	(PlatformType_Type)(0),  // 0: auth.pb.PlatformType.Type
	(*PlatformType)(nil),    // 1: auth.pb.PlatformType
	(*Session)(nil),         // 2: auth.pb.Session
	(*BindRequest)(nil),     // 3: auth.pb.BindRequest
	(*BindResponse)(nil),    // 4: auth.pb.BindResponse
	(*NothingResponse)(nil), // 5: auth.pb.NothingResponse
}
var file_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformType); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindRequest); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindResponse); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	file_message_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BindRequest_FacebookToken)(nil),
		(*BindRequest_GoogleToken)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
