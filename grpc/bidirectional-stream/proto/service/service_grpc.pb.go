// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: service.proto

package service

import (
	message "bidirectional-stream/proto/message"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BidirectionalStreamClient is the client API for BidirectionalStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BidirectionalStreamClient interface {
	StartChat(ctx context.Context, opts ...grpc.CallOption) (BidirectionalStream_StartChatClient, error)
}

type bidirectionalStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewBidirectionalStreamClient(cc grpc.ClientConnInterface) BidirectionalStreamClient {
	return &bidirectionalStreamClient{cc}
}

func (c *bidirectionalStreamClient) StartChat(ctx context.Context, opts ...grpc.CallOption) (BidirectionalStream_StartChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &BidirectionalStream_ServiceDesc.Streams[0], "/BidirectionalStream/StartChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &bidirectionalStreamStartChatClient{stream}
	return x, nil
}

type BidirectionalStream_StartChatClient interface {
	Send(*message.Request) error
	Recv() (*message.Response, error)
	grpc.ClientStream
}

type bidirectionalStreamStartChatClient struct {
	grpc.ClientStream
}

func (x *bidirectionalStreamStartChatClient) Send(m *message.Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bidirectionalStreamStartChatClient) Recv() (*message.Response, error) {
	m := new(message.Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BidirectionalStreamServer is the server API for BidirectionalStream service.
// All implementations must embed UnimplementedBidirectionalStreamServer
// for forward compatibility
type BidirectionalStreamServer interface {
	StartChat(BidirectionalStream_StartChatServer) error
	mustEmbedUnimplementedBidirectionalStreamServer()
}

// UnimplementedBidirectionalStreamServer must be embedded to have forward compatible implementations.
type UnimplementedBidirectionalStreamServer struct {
}

func (UnimplementedBidirectionalStreamServer) StartChat(BidirectionalStream_StartChatServer) error {
	return status.Errorf(codes.Unimplemented, "method StartChat not implemented")
}
func (UnimplementedBidirectionalStreamServer) mustEmbedUnimplementedBidirectionalStreamServer() {}

// UnsafeBidirectionalStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BidirectionalStreamServer will
// result in compilation errors.
type UnsafeBidirectionalStreamServer interface {
	mustEmbedUnimplementedBidirectionalStreamServer()
}

func RegisterBidirectionalStreamServer(s grpc.ServiceRegistrar, srv BidirectionalStreamServer) {
	s.RegisterService(&BidirectionalStream_ServiceDesc, srv)
}

func _BidirectionalStream_StartChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BidirectionalStreamServer).StartChat(&bidirectionalStreamStartChatServer{stream})
}

type BidirectionalStream_StartChatServer interface {
	Send(*message.Response) error
	Recv() (*message.Request, error)
	grpc.ServerStream
}

type bidirectionalStreamStartChatServer struct {
	grpc.ServerStream
}

func (x *bidirectionalStreamStartChatServer) Send(m *message.Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bidirectionalStreamStartChatServer) Recv() (*message.Request, error) {
	m := new(message.Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BidirectionalStream_ServiceDesc is the grpc.ServiceDesc for BidirectionalStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BidirectionalStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BidirectionalStream",
	HandlerType: (*BidirectionalStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartChat",
			Handler:       _BidirectionalStream_StartChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}