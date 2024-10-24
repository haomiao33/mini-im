// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: proto/MsgService.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MsgService_SendMessage_FullMethodName = "/myservice.MsgService/SendMessage"
	MsgService_SyncMessage_FullMethodName = "/myservice.MsgService/SyncMessage"
)

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgServiceClient interface {
	// 发送单聊消息
	SendMessage(ctx context.Context, in *ImMsgRequest, opts ...grpc.CallOption) (*ImMsgResponse, error)
	// 同步消息
	SyncMessage(ctx context.Context, in *ImMsgSyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
}

type msgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgServiceClient(cc grpc.ClientConnInterface) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) SendMessage(ctx context.Context, in *ImMsgRequest, opts ...grpc.CallOption) (*ImMsgResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImMsgResponse)
	err := c.cc.Invoke(ctx, MsgService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SyncMessage(ctx context.Context, in *ImMsgSyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, MsgService_SyncMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
// All implementations must embed UnimplementedMsgServiceServer
// for forward compatibility.
type MsgServiceServer interface {
	// 发送单聊消息
	SendMessage(context.Context, *ImMsgRequest) (*ImMsgResponse, error)
	// 同步消息
	SyncMessage(context.Context, *ImMsgSyncRequest) (*SyncResponse, error)
	mustEmbedUnimplementedMsgServiceServer()
}

// UnimplementedMsgServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMsgServiceServer struct{}

func (UnimplementedMsgServiceServer) SendMessage(context.Context, *ImMsgRequest) (*ImMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMsgServiceServer) SyncMessage(context.Context, *ImMsgSyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncMessage not implemented")
}
func (UnimplementedMsgServiceServer) mustEmbedUnimplementedMsgServiceServer() {}
func (UnimplementedMsgServiceServer) testEmbeddedByValue()                    {}

// UnsafeMsgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServiceServer will
// result in compilation errors.
type UnsafeMsgServiceServer interface {
	mustEmbedUnimplementedMsgServiceServer()
}

func RegisterMsgServiceServer(s grpc.ServiceRegistrar, srv MsgServiceServer) {
	// If the following call pancis, it indicates UnimplementedMsgServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MsgService_ServiceDesc, srv)
}

func _MsgService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SendMessage(ctx, req.(*ImMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SyncMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImMsgSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SyncMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgService_SyncMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SyncMessage(ctx, req.(*ImMsgSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MsgService_ServiceDesc is the grpc.ServiceDesc for MsgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MsgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myservice.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _MsgService_SendMessage_Handler,
		},
		{
			MethodName: "SyncMessage",
			Handler:    _MsgService_SyncMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/MsgService.proto",
}
