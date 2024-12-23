// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin.proto

package gen

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
	AdminAuth_AdminLogin_FullMethodName             = "/book_store.admin_service.AdminAuth/AdminLogin"
	AdminAuth_AdminCreate_FullMethodName            = "/book_store.admin_service.AdminAuth/AdminCreate"
	AdminAuth_VerifyAdminCredentials_FullMethodName = "/book_store.admin_service.AdminAuth/VerifyAdminCredentials"
)

// AdminAuthClient is the client API for AdminAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminAuthClient interface {
	AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
	AdminCreate(ctx context.Context, in *AdminCreateRequest, opts ...grpc.CallOption) (*AdminCreateResponse, error)
	VerifyAdminCredentials(ctx context.Context, in *VerifyAdminRequest, opts ...grpc.CallOption) (*VerifyAdminResponse, error)
}

type adminAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminAuthClient(cc grpc.ClientConnInterface) AdminAuthClient {
	return &adminAuthClient{cc}
}

func (c *adminAuthClient) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminLoginResponse)
	err := c.cc.Invoke(ctx, AdminAuth_AdminLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAuthClient) AdminCreate(ctx context.Context, in *AdminCreateRequest, opts ...grpc.CallOption) (*AdminCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminCreateResponse)
	err := c.cc.Invoke(ctx, AdminAuth_AdminCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAuthClient) VerifyAdminCredentials(ctx context.Context, in *VerifyAdminRequest, opts ...grpc.CallOption) (*VerifyAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyAdminResponse)
	err := c.cc.Invoke(ctx, AdminAuth_VerifyAdminCredentials_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminAuthServer is the server API for AdminAuth service.
// All implementations should embed UnimplementedAdminAuthServer
// for forward compatibility.
type AdminAuthServer interface {
	AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error)
	AdminCreate(context.Context, *AdminCreateRequest) (*AdminCreateResponse, error)
	VerifyAdminCredentials(context.Context, *VerifyAdminRequest) (*VerifyAdminResponse, error)
}

// UnimplementedAdminAuthServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdminAuthServer struct{}

func (UnimplementedAdminAuthServer) AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAdminAuthServer) AdminCreate(context.Context, *AdminCreateRequest) (*AdminCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreate not implemented")
}
func (UnimplementedAdminAuthServer) VerifyAdminCredentials(context.Context, *VerifyAdminRequest) (*VerifyAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAdminCredentials not implemented")
}
func (UnimplementedAdminAuthServer) testEmbeddedByValue() {}

// UnsafeAdminAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminAuthServer will
// result in compilation errors.
type UnsafeAdminAuthServer interface {
	mustEmbedUnimplementedAdminAuthServer()
}

func RegisterAdminAuthServer(s grpc.ServiceRegistrar, srv AdminAuthServer) {
	// If the following call pancis, it indicates UnimplementedAdminAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdminAuth_ServiceDesc, srv)
}

func _AdminAuth_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAuthServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAuth_AdminLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAuthServer).AdminLogin(ctx, req.(*AdminLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAuth_AdminCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAuthServer).AdminCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAuth_AdminCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAuthServer).AdminCreate(ctx, req.(*AdminCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAuth_VerifyAdminCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAuthServer).VerifyAdminCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminAuth_VerifyAdminCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAuthServer).VerifyAdminCredentials(ctx, req.(*VerifyAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminAuth_ServiceDesc is the grpc.ServiceDesc for AdminAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_store.admin_service.AdminAuth",
	HandlerType: (*AdminAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminLogin",
			Handler:    _AdminAuth_AdminLogin_Handler,
		},
		{
			MethodName: "AdminCreate",
			Handler:    _AdminAuth_AdminCreate_Handler,
		},
		{
			MethodName: "VerifyAdminCredentials",
			Handler:    _AdminAuth_VerifyAdminCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
