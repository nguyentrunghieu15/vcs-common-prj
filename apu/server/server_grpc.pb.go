// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: apu/server/server.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServerServiceClient is the client API for ServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerServiceClient interface {
	ListServers(ctx context.Context, in *ListServersRequest, opts ...grpc.CallOption) (*ListServersResponse, error)
	GetServerByName(ctx context.Context, in *GetServerByNameRequest, opts ...grpc.CallOption) (*Server, error)
	GetServerById(ctx context.Context, in *GetServerByIdRequest, opts ...grpc.CallOption) (*Server, error)
	CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*Server, error)
	UpdateServer(ctx context.Context, in *UpdateServerRequest, opts ...grpc.CallOption) (*Server, error)
	DeleteServerByName(ctx context.Context, in *DeleteServerByNameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteServerById(ctx context.Context, in *DeleteServerByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ExportServer(ctx context.Context, in *ExportServerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ImportServer(ctx context.Context, opts ...grpc.CallOption) (ServerService_ImportServerClient, error)
}

type serverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerServiceClient(cc grpc.ClientConnInterface) ServerServiceClient {
	return &serverServiceClient{cc}
}

func (c *serverServiceClient) ListServers(ctx context.Context, in *ListServersRequest, opts ...grpc.CallOption) (*ListServersResponse, error) {
	out := new(ListServersResponse)
	err := c.cc.Invoke(ctx, "/ServerService/ListServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) GetServerByName(ctx context.Context, in *GetServerByNameRequest, opts ...grpc.CallOption) (*Server, error) {
	out := new(Server)
	err := c.cc.Invoke(ctx, "/ServerService/GetServerByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) GetServerById(ctx context.Context, in *GetServerByIdRequest, opts ...grpc.CallOption) (*Server, error) {
	out := new(Server)
	err := c.cc.Invoke(ctx, "/ServerService/GetServerById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*Server, error) {
	out := new(Server)
	err := c.cc.Invoke(ctx, "/ServerService/CreateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) UpdateServer(ctx context.Context, in *UpdateServerRequest, opts ...grpc.CallOption) (*Server, error) {
	out := new(Server)
	err := c.cc.Invoke(ctx, "/ServerService/UpdateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) DeleteServerByName(ctx context.Context, in *DeleteServerByNameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ServerService/DeleteServerByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) DeleteServerById(ctx context.Context, in *DeleteServerByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ServerService/DeleteServerById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) ExportServer(ctx context.Context, in *ExportServerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ServerService/ExportServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) ImportServer(ctx context.Context, opts ...grpc.CallOption) (ServerService_ImportServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServerService_ServiceDesc.Streams[0], "/ServerService/ImportServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverServiceImportServerClient{stream}
	return x, nil
}

type ServerService_ImportServerClient interface {
	Send(*ImportServerRequest) error
	CloseAndRecv() (*ImportServerResponse, error)
	grpc.ClientStream
}

type serverServiceImportServerClient struct {
	grpc.ClientStream
}

func (x *serverServiceImportServerClient) Send(m *ImportServerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serverServiceImportServerClient) CloseAndRecv() (*ImportServerResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ImportServerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerServiceServer is the server API for ServerService service.
// All implementations must embed UnimplementedServerServiceServer
// for forward compatibility
type ServerServiceServer interface {
	ListServers(context.Context, *ListServersRequest) (*ListServersResponse, error)
	GetServerByName(context.Context, *GetServerByNameRequest) (*Server, error)
	GetServerById(context.Context, *GetServerByIdRequest) (*Server, error)
	CreateServer(context.Context, *CreateServerRequest) (*Server, error)
	UpdateServer(context.Context, *UpdateServerRequest) (*Server, error)
	DeleteServerByName(context.Context, *DeleteServerByNameRequest) (*emptypb.Empty, error)
	DeleteServerById(context.Context, *DeleteServerByIdRequest) (*emptypb.Empty, error)
	ExportServer(context.Context, *ExportServerRequest) (*emptypb.Empty, error)
	ImportServer(ServerService_ImportServerServer) error
	mustEmbedUnimplementedServerServiceServer()
}

// UnimplementedServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerServiceServer struct {
}

func (UnimplementedServerServiceServer) ListServers(context.Context, *ListServersRequest) (*ListServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServers not implemented")
}
func (UnimplementedServerServiceServer) GetServerByName(context.Context, *GetServerByNameRequest) (*Server, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerByName not implemented")
}
func (UnimplementedServerServiceServer) GetServerById(context.Context, *GetServerByIdRequest) (*Server, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerById not implemented")
}
func (UnimplementedServerServiceServer) CreateServer(context.Context, *CreateServerRequest) (*Server, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServer not implemented")
}
func (UnimplementedServerServiceServer) UpdateServer(context.Context, *UpdateServerRequest) (*Server, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServer not implemented")
}
func (UnimplementedServerServiceServer) DeleteServerByName(context.Context, *DeleteServerByNameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServerByName not implemented")
}
func (UnimplementedServerServiceServer) DeleteServerById(context.Context, *DeleteServerByIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServerById not implemented")
}
func (UnimplementedServerServiceServer) ExportServer(context.Context, *ExportServerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportServer not implemented")
}
func (UnimplementedServerServiceServer) ImportServer(ServerService_ImportServerServer) error {
	return status.Errorf(codes.Unimplemented, "method ImportServer not implemented")
}
func (UnimplementedServerServiceServer) mustEmbedUnimplementedServerServiceServer() {}

// UnsafeServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServiceServer will
// result in compilation errors.
type UnsafeServerServiceServer interface {
	mustEmbedUnimplementedServerServiceServer()
}

func RegisterServerServiceServer(s grpc.ServiceRegistrar, srv ServerServiceServer) {
	s.RegisterService(&ServerService_ServiceDesc, srv)
}

func _ServerService_ListServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).ListServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerService/ListServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).ListServers(ctx, req.(*ListServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_GetServerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).GetServerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerService/GetServerByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).GetServerByName(ctx, req.(*GetServerByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_GetServerById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).GetServerById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerService/GetServerById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).GetServerById(ctx, req.(*GetServerByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_CreateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).CreateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerService/CreateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).CreateServer(ctx, req.(*CreateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_UpdateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).UpdateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerService/UpdateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).UpdateServer(ctx, req.(*UpdateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_DeleteServerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServerByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).DeleteServerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerService/DeleteServerByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).DeleteServerByName(ctx, req.(*DeleteServerByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_DeleteServerById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServerByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).DeleteServerById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerService/DeleteServerById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).DeleteServerById(ctx, req.(*DeleteServerByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_ExportServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).ExportServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerService/ExportServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).ExportServer(ctx, req.(*ExportServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_ImportServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServerServiceServer).ImportServer(&serverServiceImportServerServer{stream})
}

type ServerService_ImportServerServer interface {
	SendAndClose(*ImportServerResponse) error
	Recv() (*ImportServerRequest, error)
	grpc.ServerStream
}

type serverServiceImportServerServer struct {
	grpc.ServerStream
}

func (x *serverServiceImportServerServer) SendAndClose(m *ImportServerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serverServiceImportServerServer) Recv() (*ImportServerRequest, error) {
	m := new(ImportServerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerService_ServiceDesc is the grpc.ServiceDesc for ServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServerService",
	HandlerType: (*ServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServers",
			Handler:    _ServerService_ListServers_Handler,
		},
		{
			MethodName: "GetServerByName",
			Handler:    _ServerService_GetServerByName_Handler,
		},
		{
			MethodName: "GetServerById",
			Handler:    _ServerService_GetServerById_Handler,
		},
		{
			MethodName: "CreateServer",
			Handler:    _ServerService_CreateServer_Handler,
		},
		{
			MethodName: "UpdateServer",
			Handler:    _ServerService_UpdateServer_Handler,
		},
		{
			MethodName: "DeleteServerByName",
			Handler:    _ServerService_DeleteServerByName_Handler,
		},
		{
			MethodName: "DeleteServerById",
			Handler:    _ServerService_DeleteServerById_Handler,
		},
		{
			MethodName: "ExportServer",
			Handler:    _ServerService_ExportServer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ImportServer",
			Handler:       _ServerService_ImportServer_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "apu/server/server.proto",
}
