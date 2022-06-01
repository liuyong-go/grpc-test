// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0--rc2
// source: proto/sumnum.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SumNumClient is the client API for SumNum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SumNumClient interface {
	SumTwoNums(ctx context.Context, in *SumTwoRequest, opts ...grpc.CallOption) (*SumReply, error)
	SumNums(ctx context.Context, opts ...grpc.CallOption) (SumNum_SumNumsClient, error)
}

type sumNumClient struct {
	cc grpc.ClientConnInterface
}

func NewSumNumClient(cc grpc.ClientConnInterface) SumNumClient {
	return &sumNumClient{cc}
}

func (c *sumNumClient) SumTwoNums(ctx context.Context, in *SumTwoRequest, opts ...grpc.CallOption) (*SumReply, error) {
	out := new(SumReply)
	err := c.cc.Invoke(ctx, "/sumnum.SumNum/SumTwoNums", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumNumClient) SumNums(ctx context.Context, opts ...grpc.CallOption) (SumNum_SumNumsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SumNum_ServiceDesc.Streams[0], "/sumnum.SumNum/SumNums", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumNumSumNumsClient{stream}
	return x, nil
}

type SumNum_SumNumsClient interface {
	Send(*NumsRequest) error
	CloseAndRecv() (*SumReply, error)
	grpc.ClientStream
}

type sumNumSumNumsClient struct {
	grpc.ClientStream
}

func (x *sumNumSumNumsClient) Send(m *NumsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumNumSumNumsClient) CloseAndRecv() (*SumReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SumReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumNumServer is the server API for SumNum service.
// All implementations must embed UnimplementedSumNumServer
// for forward compatibility
type SumNumServer interface {
	SumTwoNums(context.Context, *SumTwoRequest) (*SumReply, error)
	SumNums(SumNum_SumNumsServer) error
	mustEmbedUnimplementedSumNumServer()
}

// UnimplementedSumNumServer must be embedded to have forward compatible implementations.
type UnimplementedSumNumServer struct {
}

func (UnimplementedSumNumServer) SumTwoNums(context.Context, *SumTwoRequest) (*SumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumTwoNums not implemented")
}
func (UnimplementedSumNumServer) SumNums(SumNum_SumNumsServer) error {
	return status.Errorf(codes.Unimplemented, "method SumNums not implemented")
}
func (UnimplementedSumNumServer) mustEmbedUnimplementedSumNumServer() {}

// UnsafeSumNumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SumNumServer will
// result in compilation errors.
type UnsafeSumNumServer interface {
	mustEmbedUnimplementedSumNumServer()
}

func RegisterSumNumServer(s grpc.ServiceRegistrar, srv SumNumServer) {
	s.RegisterService(&SumNum_ServiceDesc, srv)
}

func _SumNum_SumTwoNums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumTwoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumNumServer).SumTwoNums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sumnum.SumNum/SumTwoNums",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumNumServer).SumTwoNums(ctx, req.(*SumTwoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumNum_SumNums_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumNumServer).SumNums(&sumNumSumNumsServer{stream})
}

type SumNum_SumNumsServer interface {
	SendAndClose(*SumReply) error
	Recv() (*NumsRequest, error)
	grpc.ServerStream
}

type sumNumSumNumsServer struct {
	grpc.ServerStream
}

func (x *sumNumSumNumsServer) SendAndClose(m *SumReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumNumSumNumsServer) Recv() (*NumsRequest, error) {
	m := new(NumsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumNum_ServiceDesc is the grpc.ServiceDesc for SumNum service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SumNum_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sumnum.SumNum",
	HandlerType: (*SumNumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SumTwoNums",
			Handler:    _SumNum_SumTwoNums_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SumNums",
			Handler:       _SumNum_SumNums_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/sumnum.proto",
}