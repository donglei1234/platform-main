// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: notification/api/service.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_notification_api_service_proto protoreflect.FileDescriptor

var file_notification_api_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa4,
	0x04, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x75, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x6e, 0x12, 0x23,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x61, 0x72, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x97, 0x01, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x72, 0x6e, 0x12, 0x21, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x3f, 0x2a, 0x3d, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x72, 0x6e, 0x2f, 0x7b, 0x61, 0x70, 0x70, 0x49, 0x64, 0x7d, 0x2f,
	0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x7d, 0x2f, 0x7b, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x49, 0x64, 0x7d, 0x2f, 0x7b, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x7d,
	0x12, 0x7f, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x81, 0x01, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x4e,
	0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x15, 0x5a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_notification_api_service_proto_goTypes = []interface{}{
	(*RegisterArnRequest)(nil),    // 0: notification.pb.RegisterArnRequest
	(*DeleteArnRequest)(nil),      // 1: notification.pb.DeleteArnRequest
	(*PublishMessageRequest)(nil), // 2: notification.pb.PublishMessageRequest
	(*SubscribeTopicRequest)(nil), // 3: notification.pb.SubscribeTopicRequest
	(*NothingResponse)(nil),       // 4: notification.pb.NothingResponse
}
var file_notification_api_service_proto_depIdxs = []int32{
	0, // 0: notification.pb.Notification.RegisterArn:input_type -> notification.pb.RegisterArnRequest
	1, // 1: notification.pb.Notification.DeleteArn:input_type -> notification.pb.DeleteArnRequest
	2, // 2: notification.pb.Notification.PublishMessage:input_type -> notification.pb.PublishMessageRequest
	3, // 3: notification.pb.Notification.SubscribeTopic:input_type -> notification.pb.SubscribeTopicRequest
	4, // 4: notification.pb.Notification.RegisterArn:output_type -> notification.pb.NothingResponse
	4, // 5: notification.pb.Notification.DeleteArn:output_type -> notification.pb.NothingResponse
	4, // 6: notification.pb.Notification.PublishMessage:output_type -> notification.pb.NothingResponse
	4, // 7: notification.pb.Notification.SubscribeTopic:output_type -> notification.pb.NothingResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notification_api_service_proto_init() }
func file_notification_api_service_proto_init() {
	if File_notification_api_service_proto != nil {
		return
	}
	file_notification_api_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notification_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_api_service_proto_goTypes,
		DependencyIndexes: file_notification_api_service_proto_depIdxs,
	}.Build()
	File_notification_api_service_proto = out.File
	file_notification_api_service_proto_rawDesc = nil
	file_notification_api_service_proto_goTypes = nil
	file_notification_api_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationClient interface {
	RegisterArn(ctx context.Context, in *RegisterArnRequest, opts ...grpc.CallOption) (*NothingResponse, error)
	DeleteArn(ctx context.Context, in *DeleteArnRequest, opts ...grpc.CallOption) (*NothingResponse, error)
	PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...grpc.CallOption) (*NothingResponse, error)
	SubscribeTopic(ctx context.Context, in *SubscribeTopicRequest, opts ...grpc.CallOption) (*NothingResponse, error)
}

type notificationClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationClient(cc grpc.ClientConnInterface) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) RegisterArn(ctx context.Context, in *RegisterArnRequest, opts ...grpc.CallOption) (*NothingResponse, error) {
	out := new(NothingResponse)
	err := c.cc.Invoke(ctx, "/notification.pb.Notification/RegisterArn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) DeleteArn(ctx context.Context, in *DeleteArnRequest, opts ...grpc.CallOption) (*NothingResponse, error) {
	out := new(NothingResponse)
	err := c.cc.Invoke(ctx, "/notification.pb.Notification/DeleteArn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...grpc.CallOption) (*NothingResponse, error) {
	out := new(NothingResponse)
	err := c.cc.Invoke(ctx, "/notification.pb.Notification/PublishMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) SubscribeTopic(ctx context.Context, in *SubscribeTopicRequest, opts ...grpc.CallOption) (*NothingResponse, error) {
	out := new(NothingResponse)
	err := c.cc.Invoke(ctx, "/notification.pb.Notification/SubscribeTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServer is the server API for Notification service.
type NotificationServer interface {
	RegisterArn(context.Context, *RegisterArnRequest) (*NothingResponse, error)
	DeleteArn(context.Context, *DeleteArnRequest) (*NothingResponse, error)
	PublishMessage(context.Context, *PublishMessageRequest) (*NothingResponse, error)
	SubscribeTopic(context.Context, *SubscribeTopicRequest) (*NothingResponse, error)
}

// UnimplementedNotificationServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (*UnimplementedNotificationServer) RegisterArn(context.Context, *RegisterArnRequest) (*NothingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterArn not implemented")
}
func (*UnimplementedNotificationServer) DeleteArn(context.Context, *DeleteArnRequest) (*NothingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArn not implemented")
}
func (*UnimplementedNotificationServer) PublishMessage(context.Context, *PublishMessageRequest) (*NothingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishMessage not implemented")
}
func (*UnimplementedNotificationServer) SubscribeTopic(context.Context, *SubscribeTopicRequest) (*NothingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeTopic not implemented")
}

func RegisterNotificationServer(s *grpc.Server, srv NotificationServer) {
	s.RegisterService(&_Notification_serviceDesc, srv)
}

func _Notification_RegisterArn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterArnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).RegisterArn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.pb.Notification/RegisterArn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).RegisterArn(ctx, req.(*RegisterArnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_DeleteArn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).DeleteArn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.pb.Notification/DeleteArn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).DeleteArn(ctx, req.(*DeleteArnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_PublishMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).PublishMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.pb.Notification/PublishMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).PublishMessage(ctx, req.(*PublishMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_SubscribeTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).SubscribeTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.pb.Notification/SubscribeTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).SubscribeTopic(ctx, req.(*SubscribeTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notification_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notification.pb.Notification",
	HandlerType: (*NotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterArn",
			Handler:    _Notification_RegisterArn_Handler,
		},
		{
			MethodName: "DeleteArn",
			Handler:    _Notification_DeleteArn_Handler,
		},
		{
			MethodName: "PublishMessage",
			Handler:    _Notification_PublishMessage_Handler,
		},
		{
			MethodName: "SubscribeTopic",
			Handler:    _Notification_SubscribeTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification/api/service.proto",
}
