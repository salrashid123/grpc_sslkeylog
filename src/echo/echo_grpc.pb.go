// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package echo

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

// EchoServerClient is the client API for EchoServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoServerClient interface {
	SayHelloUnary(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
	SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (EchoServer_SayHelloClientStreamClient, error)
	SayHelloServerStream(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (EchoServer_SayHelloServerStreamClient, error)
	SayHelloBiDiStream(ctx context.Context, opts ...grpc.CallOption) (EchoServer_SayHelloBiDiStreamClient, error)
}

type echoServerClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServerClient(cc grpc.ClientConnInterface) EchoServerClient {
	return &echoServerClient{cc}
}

func (c *echoServerClient) SayHelloUnary(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := c.cc.Invoke(ctx, "/echo.EchoServer/SayHelloUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServerClient) SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (EchoServer_SayHelloClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EchoServer_ServiceDesc.Streams[0], "/echo.EchoServer/SayHelloClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServerSayHelloClientStreamClient{stream}
	return x, nil
}

type EchoServer_SayHelloClientStreamClient interface {
	Send(*EchoRequest) error
	CloseAndRecv() (*EchoReply, error)
	grpc.ClientStream
}

type echoServerSayHelloClientStreamClient struct {
	grpc.ClientStream
}

func (x *echoServerSayHelloClientStreamClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServerSayHelloClientStreamClient) CloseAndRecv() (*EchoReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EchoReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServerClient) SayHelloServerStream(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (EchoServer_SayHelloServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EchoServer_ServiceDesc.Streams[1], "/echo.EchoServer/SayHelloServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServerSayHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EchoServer_SayHelloServerStreamClient interface {
	Recv() (*EchoReply, error)
	grpc.ClientStream
}

type echoServerSayHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *echoServerSayHelloServerStreamClient) Recv() (*EchoReply, error) {
	m := new(EchoReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServerClient) SayHelloBiDiStream(ctx context.Context, opts ...grpc.CallOption) (EchoServer_SayHelloBiDiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EchoServer_ServiceDesc.Streams[2], "/echo.EchoServer/SayHelloBiDiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServerSayHelloBiDiStreamClient{stream}
	return x, nil
}

type EchoServer_SayHelloBiDiStreamClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoReply, error)
	grpc.ClientStream
}

type echoServerSayHelloBiDiStreamClient struct {
	grpc.ClientStream
}

func (x *echoServerSayHelloBiDiStreamClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServerSayHelloBiDiStreamClient) Recv() (*EchoReply, error) {
	m := new(EchoReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServerServer is the server API for EchoServer service.
// All implementations should embed UnimplementedEchoServerServer
// for forward compatibility
type EchoServerServer interface {
	SayHelloUnary(context.Context, *EchoRequest) (*EchoReply, error)
	SayHelloClientStream(EchoServer_SayHelloClientStreamServer) error
	SayHelloServerStream(*EchoRequest, EchoServer_SayHelloServerStreamServer) error
	SayHelloBiDiStream(EchoServer_SayHelloBiDiStreamServer) error
}

// UnimplementedEchoServerServer should be embedded to have forward compatible implementations.
type UnimplementedEchoServerServer struct {
}

func (UnimplementedEchoServerServer) SayHelloUnary(context.Context, *EchoRequest) (*EchoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloUnary not implemented")
}
func (UnimplementedEchoServerServer) SayHelloClientStream(EchoServer_SayHelloClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStream not implemented")
}
func (UnimplementedEchoServerServer) SayHelloServerStream(*EchoRequest, EchoServer_SayHelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStream not implemented")
}
func (UnimplementedEchoServerServer) SayHelloBiDiStream(EchoServer_SayHelloBiDiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBiDiStream not implemented")
}

// UnsafeEchoServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServerServer will
// result in compilation errors.
type UnsafeEchoServerServer interface {
	mustEmbedUnimplementedEchoServerServer()
}

func RegisterEchoServerServer(s grpc.ServiceRegistrar, srv EchoServerServer) {
	s.RegisterService(&EchoServer_ServiceDesc, srv)
}

func _EchoServer_SayHelloUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServerServer).SayHelloUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.EchoServer/SayHelloUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServerServer).SayHelloUnary(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoServer_SayHelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServerServer).SayHelloClientStream(&echoServerSayHelloClientStreamServer{stream})
}

type EchoServer_SayHelloClientStreamServer interface {
	SendAndClose(*EchoReply) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoServerSayHelloClientStreamServer struct {
	grpc.ServerStream
}

func (x *echoServerSayHelloClientStreamServer) SendAndClose(m *EchoReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServerSayHelloClientStreamServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EchoServer_SayHelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EchoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServerServer).SayHelloServerStream(m, &echoServerSayHelloServerStreamServer{stream})
}

type EchoServer_SayHelloServerStreamServer interface {
	Send(*EchoReply) error
	grpc.ServerStream
}

type echoServerSayHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *echoServerSayHelloServerStreamServer) Send(m *EchoReply) error {
	return x.ServerStream.SendMsg(m)
}

func _EchoServer_SayHelloBiDiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServerServer).SayHelloBiDiStream(&echoServerSayHelloBiDiStreamServer{stream})
}

type EchoServer_SayHelloBiDiStreamServer interface {
	Send(*EchoReply) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoServerSayHelloBiDiStreamServer struct {
	grpc.ServerStream
}

func (x *echoServerSayHelloBiDiStreamServer) Send(m *EchoReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServerSayHelloBiDiStreamServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServer_ServiceDesc is the grpc.ServiceDesc for EchoServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EchoServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "echo.EchoServer",
	HandlerType: (*EchoServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloUnary",
			Handler:    _EchoServer_SayHelloUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloClientStream",
			Handler:       _EchoServer_SayHelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloServerStream",
			Handler:       _EchoServer_SayHelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloBiDiStream",
			Handler:       _EchoServer_SayHelloBiDiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "src/echo/echo.proto",
}
