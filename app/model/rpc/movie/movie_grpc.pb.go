// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.0--rc1
// source: app/model/rpc/movie/movie.proto

package movie

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	AddMovie(ctx context.Context, in *AddMovieReq, opts ...grpc.CallOption) (*AddMovieRes, error)
	DetailMovie(ctx context.Context, in *DetailMovieReq, opts ...grpc.CallOption) (*AddMovieRes, error)
	ListMovie(ctx context.Context, in *AbstractReq, opts ...grpc.CallOption) (*ListMovieRes, error)
	RemoveMovie(ctx context.Context, in *AbstractReq, opts ...grpc.CallOption) (*RemoveMovieRes, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) AddMovie(ctx context.Context, in *AddMovieReq, opts ...grpc.CallOption) (*AddMovieRes, error) {
	out := new(AddMovieRes)
	err := c.cc.Invoke(ctx, "/model.AuthService/AddMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DetailMovie(ctx context.Context, in *DetailMovieReq, opts ...grpc.CallOption) (*AddMovieRes, error) {
	out := new(AddMovieRes)
	err := c.cc.Invoke(ctx, "/model.AuthService/DetailMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListMovie(ctx context.Context, in *AbstractReq, opts ...grpc.CallOption) (*ListMovieRes, error) {
	out := new(ListMovieRes)
	err := c.cc.Invoke(ctx, "/model.AuthService/ListMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RemoveMovie(ctx context.Context, in *AbstractReq, opts ...grpc.CallOption) (*RemoveMovieRes, error) {
	out := new(RemoveMovieRes)
	err := c.cc.Invoke(ctx, "/model.AuthService/RemoveMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	AddMovie(context.Context, *AddMovieReq) (*AddMovieRes, error)
	DetailMovie(context.Context, *DetailMovieReq) (*AddMovieRes, error)
	ListMovie(context.Context, *AbstractReq) (*ListMovieRes, error)
	RemoveMovie(context.Context, *AbstractReq) (*RemoveMovieRes, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) AddMovie(context.Context, *AddMovieReq) (*AddMovieRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMovie not implemented")
}
func (UnimplementedAuthServiceServer) DetailMovie(context.Context, *DetailMovieReq) (*AddMovieRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailMovie not implemented")
}
func (UnimplementedAuthServiceServer) ListMovie(context.Context, *AbstractReq) (*ListMovieRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMovie not implemented")
}
func (UnimplementedAuthServiceServer) RemoveMovie(context.Context, *AbstractReq) (*RemoveMovieRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMovie not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_AddMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMovieReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.AuthService/AddMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddMovie(ctx, req.(*AddMovieReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DetailMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailMovieReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DetailMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.AuthService/DetailMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DetailMovie(ctx, req.(*DetailMovieReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbstractReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.AuthService/ListMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListMovie(ctx, req.(*AbstractReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RemoveMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbstractReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RemoveMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.AuthService/RemoveMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RemoveMovie(ctx, req.(*AbstractReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMovie",
			Handler:    _AuthService_AddMovie_Handler,
		},
		{
			MethodName: "DetailMovie",
			Handler:    _AuthService_DetailMovie_Handler,
		},
		{
			MethodName: "ListMovie",
			Handler:    _AuthService_ListMovie_Handler,
		},
		{
			MethodName: "RemoveMovie",
			Handler:    _AuthService_RemoveMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/model/rpc/movie/movie.proto",
}