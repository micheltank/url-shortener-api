// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: url.proto

package pb

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

// UrlHandlerClient is the client API for UrlHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlHandlerClient interface {
	Create(ctx context.Context, in *LongUrlParam, opts ...grpc.CallOption) (*ShortUrl, error)
	Get(ctx context.Context, in *ShortUrlParam, opts ...grpc.CallOption) (*LongUrl, error)
	Delete(ctx context.Context, in *ShortUrlParam, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type urlHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlHandlerClient(cc grpc.ClientConnInterface) UrlHandlerClient {
	return &urlHandlerClient{cc}
}

func (c *urlHandlerClient) Create(ctx context.Context, in *LongUrlParam, opts ...grpc.CallOption) (*ShortUrl, error) {
	out := new(ShortUrl)
	err := c.cc.Invoke(ctx, "/pb.UrlHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlHandlerClient) Get(ctx context.Context, in *ShortUrlParam, opts ...grpc.CallOption) (*LongUrl, error) {
	out := new(LongUrl)
	err := c.cc.Invoke(ctx, "/pb.UrlHandler/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlHandlerClient) Delete(ctx context.Context, in *ShortUrlParam, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.UrlHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlHandlerServer is the server API for UrlHandler service.
// All implementations must embed UnimplementedUrlHandlerServer
// for forward compatibility
type UrlHandlerServer interface {
	Create(context.Context, *LongUrlParam) (*ShortUrl, error)
	Get(context.Context, *ShortUrlParam) (*LongUrl, error)
	Delete(context.Context, *ShortUrlParam) (*emptypb.Empty, error)
	mustEmbedUnimplementedUrlHandlerServer()
}

// UnimplementedUrlHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedUrlHandlerServer struct {
}

func (UnimplementedUrlHandlerServer) Create(context.Context, *LongUrlParam) (*ShortUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUrlHandlerServer) Get(context.Context, *ShortUrlParam) (*LongUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUrlHandlerServer) Delete(context.Context, *ShortUrlParam) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUrlHandlerServer) mustEmbedUnimplementedUrlHandlerServer() {}

// UnsafeUrlHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlHandlerServer will
// result in compilation errors.
type UnsafeUrlHandlerServer interface {
	mustEmbedUnimplementedUrlHandlerServer()
}

func RegisterUrlHandlerServer(s grpc.ServiceRegistrar, srv UrlHandlerServer) {
	s.RegisterService(&UrlHandler_ServiceDesc, srv)
}

func _UrlHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LongUrlParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UrlHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlHandlerServer).Create(ctx, req.(*LongUrlParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlHandler_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortUrlParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlHandlerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UrlHandler/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlHandlerServer).Get(ctx, req.(*ShortUrlParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortUrlParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UrlHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlHandlerServer).Delete(ctx, req.(*ShortUrlParam))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlHandler_ServiceDesc is the grpc.ServiceDesc for UrlHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UrlHandler",
	HandlerType: (*UrlHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UrlHandler_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UrlHandler_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UrlHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "url.proto",
}
