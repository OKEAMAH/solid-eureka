// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package spyv1

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

// SpyRPCServiceClient is the client API for SpyRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpyRPCServiceClient interface {
	// SubscribeSignedVAA returns a stream of signed VAA messages received on the network.
	SubscribeSignedVAA(ctx context.Context, in *SubscribeSignedVAARequest, opts ...grpc.CallOption) (SpyRPCService_SubscribeSignedVAAClient, error)
}

type spyRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpyRPCServiceClient(cc grpc.ClientConnInterface) SpyRPCServiceClient {
	return &spyRPCServiceClient{cc}
}

func (c *spyRPCServiceClient) SubscribeSignedVAA(ctx context.Context, in *SubscribeSignedVAARequest, opts ...grpc.CallOption) (SpyRPCService_SubscribeSignedVAAClient, error) {
	stream, err := c.cc.NewStream(ctx, &SpyRPCService_ServiceDesc.Streams[0], "/spy.v1.SpyRPCService/SubscribeSignedVAA", opts...)
	if err != nil {
		return nil, err
	}
	x := &spyRPCServiceSubscribeSignedVAAClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SpyRPCService_SubscribeSignedVAAClient interface {
	Recv() (*SubscribeSignedVAAResponse, error)
	grpc.ClientStream
}

type spyRPCServiceSubscribeSignedVAAClient struct {
	grpc.ClientStream
}

func (x *spyRPCServiceSubscribeSignedVAAClient) Recv() (*SubscribeSignedVAAResponse, error) {
	m := new(SubscribeSignedVAAResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpyRPCServiceServer is the server API for SpyRPCService service.
// All implementations must embed UnimplementedSpyRPCServiceServer
// for forward compatibility
type SpyRPCServiceServer interface {
	// SubscribeSignedVAA returns a stream of signed VAA messages received on the network.
	SubscribeSignedVAA(*SubscribeSignedVAARequest, SpyRPCService_SubscribeSignedVAAServer) error
	mustEmbedUnimplementedSpyRPCServiceServer()
}

// UnimplementedSpyRPCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpyRPCServiceServer struct {
}

func (UnimplementedSpyRPCServiceServer) SubscribeSignedVAA(*SubscribeSignedVAARequest, SpyRPCService_SubscribeSignedVAAServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeSignedVAA not implemented")
}
func (UnimplementedSpyRPCServiceServer) mustEmbedUnimplementedSpyRPCServiceServer() {}

// UnsafeSpyRPCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpyRPCServiceServer will
// result in compilation errors.
type UnsafeSpyRPCServiceServer interface {
	mustEmbedUnimplementedSpyRPCServiceServer()
}

func RegisterSpyRPCServiceServer(s grpc.ServiceRegistrar, srv SpyRPCServiceServer) {
	s.RegisterService(&SpyRPCService_ServiceDesc, srv)
}

func _SpyRPCService_SubscribeSignedVAA_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeSignedVAARequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpyRPCServiceServer).SubscribeSignedVAA(m, &spyRPCServiceSubscribeSignedVAAServer{stream})
}

type SpyRPCService_SubscribeSignedVAAServer interface {
	Send(*SubscribeSignedVAAResponse) error
	grpc.ServerStream
}

type spyRPCServiceSubscribeSignedVAAServer struct {
	grpc.ServerStream
}

func (x *spyRPCServiceSubscribeSignedVAAServer) Send(m *SubscribeSignedVAAResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SpyRPCService_ServiceDesc is the grpc.ServiceDesc for SpyRPCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpyRPCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spy.v1.SpyRPCService",
	HandlerType: (*SpyRPCServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeSignedVAA",
			Handler:       _SpyRPCService_SubscribeSignedVAA_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spy/v1/spy.proto",
}
