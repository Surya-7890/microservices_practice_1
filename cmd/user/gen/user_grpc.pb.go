// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: user.proto

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
	UserAuth_UserLogin_FullMethodName  = "/book_store.user_service.UserAuth/UserLogin"
	UserAuth_UserSignup_FullMethodName = "/book_store.user_service.UserAuth/UserSignup"
)

// UserAuthClient is the client API for UserAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAuthClient interface {
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	UserSignup(ctx context.Context, in *UserSignupRequest, opts ...grpc.CallOption) (*UserSignupResponse, error)
}

type userAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAuthClient(cc grpc.ClientConnInterface) UserAuthClient {
	return &userAuthClient{cc}
}

func (c *userAuthClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, UserAuth_UserLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthClient) UserSignup(ctx context.Context, in *UserSignupRequest, opts ...grpc.CallOption) (*UserSignupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserSignupResponse)
	err := c.cc.Invoke(ctx, UserAuth_UserSignup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAuthServer is the server API for UserAuth service.
// All implementations should embed UnimplementedUserAuthServer
// for forward compatibility.
type UserAuthServer interface {
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	UserSignup(context.Context, *UserSignupRequest) (*UserSignupResponse, error)
}

// UnimplementedUserAuthServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserAuthServer struct{}

func (UnimplementedUserAuthServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserAuthServer) UserSignup(context.Context, *UserSignupRequest) (*UserSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignup not implemented")
}
func (UnimplementedUserAuthServer) testEmbeddedByValue() {}

// UnsafeUserAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAuthServer will
// result in compilation errors.
type UnsafeUserAuthServer interface {
	mustEmbedUnimplementedUserAuthServer()
}

func RegisterUserAuthServer(s grpc.ServiceRegistrar, srv UserAuthServer) {
	// If the following call pancis, it indicates UnimplementedUserAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserAuth_ServiceDesc, srv)
}

func _UserAuth_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAuth_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuth_UserSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).UserSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAuth_UserSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).UserSignup(ctx, req.(*UserSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAuth_ServiceDesc is the grpc.ServiceDesc for UserAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_store.user_service.UserAuth",
	HandlerType: (*UserAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserAuth_UserLogin_Handler,
		},
		{
			MethodName: "UserSignup",
			Handler:    _UserAuth_UserSignup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

const (
	UserProfile_GetUser_FullMethodName    = "/book_store.user_service.UserProfile/GetUser"
	UserProfile_UpdateUser_FullMethodName = "/book_store.user_service.UserProfile/UpdateUser"
	UserProfile_DeleteUser_FullMethodName = "/book_store.user_service.UserProfile/DeleteUser"
)

// UserProfileClient is the client API for UserProfile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProfileClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userProfileClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProfileClient(cc grpc.ClientConnInterface) UserProfileClient {
	return &userProfileClient{cc}
}

func (c *userProfileClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, UserProfile_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, UserProfile_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserProfile_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProfileServer is the server API for UserProfile service.
// All implementations should embed UnimplementedUserProfileServer
// for forward compatibility.
type UserProfileServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
}

// UnimplementedUserProfileServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserProfileServer struct{}

func (UnimplementedUserProfileServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserProfileServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserProfileServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserProfileServer) testEmbeddedByValue() {}

// UnsafeUserProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProfileServer will
// result in compilation errors.
type UnsafeUserProfileServer interface {
	mustEmbedUnimplementedUserProfileServer()
}

func RegisterUserProfileServer(s grpc.ServiceRegistrar, srv UserProfileServer) {
	// If the following call pancis, it indicates UnimplementedUserProfileServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserProfile_ServiceDesc, srv)
}

func _UserProfile_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserProfile_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfile_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserProfile_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfile_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserProfile_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserProfile_ServiceDesc is the grpc.ServiceDesc for UserProfile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserProfile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_store.user_service.UserProfile",
	HandlerType: (*UserProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserProfile_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserProfile_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserProfile_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
