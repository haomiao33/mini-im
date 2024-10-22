// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.3
// source: proto/OnlineService.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateOnlineUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	// 这里是login服务器订阅的topic
	ServerId string `protobuf:"bytes,2,opt,name=serverId,proto3" json:"serverId,omitempty"`
	// ts
	LastUpdateTs int64 `protobuf:"varint,3,opt,name=lastUpdateTs,proto3" json:"lastUpdateTs,omitempty"`
}

func (x *UpdateOnlineUserRequest) Reset() {
	*x = UpdateOnlineUserRequest{}
	mi := &file_proto_OnlineService_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOnlineUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOnlineUserRequest) ProtoMessage() {}

func (x *UpdateOnlineUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OnlineService_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOnlineUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateOnlineUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_OnlineService_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateOnlineUserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateOnlineUserRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *UpdateOnlineUserRequest) GetLastUpdateTs() int64 {
	if x != nil {
		return x.LastUpdateTs
	}
	return 0
}

type GetOnlineUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetOnlineUserRequest) Reset() {
	*x = GetOnlineUserRequest{}
	mi := &file_proto_OnlineService_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOnlineUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOnlineUserRequest) ProtoMessage() {}

func (x *GetOnlineUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OnlineService_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOnlineUserRequest.ProtoReflect.Descriptor instead.
func (*GetOnlineUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_OnlineService_proto_rawDescGZIP(), []int{1}
}

func (x *GetOnlineUserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetOnlineUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	// 这里是login服务器订阅的topic
	ServerId string `protobuf:"bytes,2,opt,name=serverId,proto3" json:"serverId,omitempty"`
	// 最后更新时间戳
	LastUpdateTs int64 `protobuf:"varint,3,opt,name=lastUpdateTs,proto3" json:"lastUpdateTs,omitempty"`
}

func (x *GetOnlineUserResp) Reset() {
	*x = GetOnlineUserResp{}
	mi := &file_proto_OnlineService_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOnlineUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOnlineUserResp) ProtoMessage() {}

func (x *GetOnlineUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OnlineService_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOnlineUserResp.ProtoReflect.Descriptor instead.
func (*GetOnlineUserResp) Descriptor() ([]byte, []int) {
	return file_proto_OnlineService_proto_rawDescGZIP(), []int{2}
}

func (x *GetOnlineUserResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetOnlineUserResp) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *GetOnlineUserResp) GetLastUpdateTs() int64 {
	if x != nil {
		return x.LastUpdateTs
	}
	return 0
}

var File_proto_OnlineService_proto protoreflect.FileDescriptor

var file_proto_OnlineService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x79, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x73, 0x22, 0x2e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x73, 0x32, 0xf7, 0x01, 0x0a, 0x0d, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x6d, 0x79, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x6e, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_OnlineService_proto_rawDescOnce sync.Once
	file_proto_OnlineService_proto_rawDescData = file_proto_OnlineService_proto_rawDesc
)

func file_proto_OnlineService_proto_rawDescGZIP() []byte {
	file_proto_OnlineService_proto_rawDescOnce.Do(func() {
		file_proto_OnlineService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_OnlineService_proto_rawDescData)
	})
	return file_proto_OnlineService_proto_rawDescData
}

var file_proto_OnlineService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_OnlineService_proto_goTypes = []any{
	(*UpdateOnlineUserRequest)(nil), // 0: myservice.UpdateOnlineUserRequest
	(*GetOnlineUserRequest)(nil),    // 1: myservice.GetOnlineUserRequest
	(*GetOnlineUserResp)(nil),       // 2: myservice.GetOnlineUserResp
	(*emptypb.Empty)(nil),           // 3: google.protobuf.Empty
}
var file_proto_OnlineService_proto_depIdxs = []int32{
	0, // 0: myservice.OnlineService.UpdateOnlineUser:input_type -> myservice.UpdateOnlineUserRequest
	1, // 1: myservice.OnlineService.GetOnlineUser:input_type -> myservice.GetOnlineUserRequest
	1, // 2: myservice.OnlineService.OutlineUser:input_type -> myservice.GetOnlineUserRequest
	3, // 3: myservice.OnlineService.UpdateOnlineUser:output_type -> google.protobuf.Empty
	2, // 4: myservice.OnlineService.GetOnlineUser:output_type -> myservice.GetOnlineUserResp
	3, // 5: myservice.OnlineService.OutlineUser:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_OnlineService_proto_init() }
func file_proto_OnlineService_proto_init() {
	if File_proto_OnlineService_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_OnlineService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_OnlineService_proto_goTypes,
		DependencyIndexes: file_proto_OnlineService_proto_depIdxs,
		MessageInfos:      file_proto_OnlineService_proto_msgTypes,
	}.Build()
	File_proto_OnlineService_proto = out.File
	file_proto_OnlineService_proto_rawDesc = nil
	file_proto_OnlineService_proto_goTypes = nil
	file_proto_OnlineService_proto_depIdxs = nil
}