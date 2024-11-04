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
	AdminAuth_AdminLogin_FullMethodName  = "/book_store.admin_service.AdminAuth/AdminLogin"
	AdminAuth_AdminCreate_FullMethodName = "/book_store.admin_service.AdminAuth/AdminCreate"
)

// AdminAuthClient is the client API for AdminAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminAuthClient interface {
	AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
	AdminCreate(ctx context.Context, in *AdminCreateRequest, opts ...grpc.CallOption) (*AdminCreateResponse, error)
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

// AdminAuthServer is the server API for AdminAuth service.
// All implementations should embed UnimplementedAdminAuthServer
// for forward compatibility.
type AdminAuthServer interface {
	AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error)
	AdminCreate(context.Context, *AdminCreateRequest) (*AdminCreateResponse, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}

const (
	AdminBooks_NewBook_FullMethodName     = "/book_store.admin_service.AdminBooks/NewBook"
	AdminBooks_UpdateBooks_FullMethodName = "/book_store.admin_service.AdminBooks/UpdateBooks"
	AdminBooks_DeleteBooks_FullMethodName = "/book_store.admin_service.AdminBooks/DeleteBooks"
)

// AdminBooksClient is the client API for AdminBooks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminBooksClient interface {
	NewBook(ctx context.Context, in *NewBookRequest, opts ...grpc.CallOption) (*NewBookResponse, error)
	UpdateBooks(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error)
	DeleteBooks(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error)
}

type adminBooksClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminBooksClient(cc grpc.ClientConnInterface) AdminBooksClient {
	return &adminBooksClient{cc}
}

func (c *adminBooksClient) NewBook(ctx context.Context, in *NewBookRequest, opts ...grpc.CallOption) (*NewBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NewBookResponse)
	err := c.cc.Invoke(ctx, AdminBooks_NewBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminBooksClient) UpdateBooks(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBookResponse)
	err := c.cc.Invoke(ctx, AdminBooks_UpdateBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminBooksClient) DeleteBooks(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBookResponse)
	err := c.cc.Invoke(ctx, AdminBooks_DeleteBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminBooksServer is the server API for AdminBooks service.
// All implementations should embed UnimplementedAdminBooksServer
// for forward compatibility.
type AdminBooksServer interface {
	NewBook(context.Context, *NewBookRequest) (*NewBookResponse, error)
	UpdateBooks(context.Context, *UpdateBookRequest) (*UpdateBookResponse, error)
	DeleteBooks(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error)
}

// UnimplementedAdminBooksServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdminBooksServer struct{}

func (UnimplementedAdminBooksServer) NewBook(context.Context, *NewBookRequest) (*NewBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBook not implemented")
}
func (UnimplementedAdminBooksServer) UpdateBooks(context.Context, *UpdateBookRequest) (*UpdateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBooks not implemented")
}
func (UnimplementedAdminBooksServer) DeleteBooks(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBooks not implemented")
}
func (UnimplementedAdminBooksServer) testEmbeddedByValue() {}

// UnsafeAdminBooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminBooksServer will
// result in compilation errors.
type UnsafeAdminBooksServer interface {
	mustEmbedUnimplementedAdminBooksServer()
}

func RegisterAdminBooksServer(s grpc.ServiceRegistrar, srv AdminBooksServer) {
	// If the following call pancis, it indicates UnimplementedAdminBooksServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdminBooks_ServiceDesc, srv)
}

func _AdminBooks_NewBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminBooksServer).NewBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminBooks_NewBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminBooksServer).NewBook(ctx, req.(*NewBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminBooks_UpdateBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminBooksServer).UpdateBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminBooks_UpdateBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminBooksServer).UpdateBooks(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminBooks_DeleteBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminBooksServer).DeleteBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminBooks_DeleteBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminBooksServer).DeleteBooks(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminBooks_ServiceDesc is the grpc.ServiceDesc for AdminBooks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminBooks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_store.admin_service.AdminBooks",
	HandlerType: (*AdminBooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewBook",
			Handler:    _AdminBooks_NewBook_Handler,
		},
		{
			MethodName: "UpdateBooks",
			Handler:    _AdminBooks_UpdateBooks_Handler,
		},
		{
			MethodName: "DeleteBooks",
			Handler:    _AdminBooks_DeleteBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}

const (
	BookCategories_CreateCategory_FullMethodName = "/book_store.admin_service.BookCategories/CreateCategory"
	BookCategories_UpdateCategory_FullMethodName = "/book_store.admin_service.BookCategories/UpdateCategory"
	BookCategories_DeleteCategory_FullMethodName = "/book_store.admin_service.BookCategories/DeleteCategory"
)

// BookCategoriesClient is the client API for BookCategories service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookCategoriesClient interface {
	CreateCategory(ctx context.Context, in *NewCategoryRequest, opts ...grpc.CallOption) (*NewCategoryResponse, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error)
}

type bookCategoriesClient struct {
	cc grpc.ClientConnInterface
}

func NewBookCategoriesClient(cc grpc.ClientConnInterface) BookCategoriesClient {
	return &bookCategoriesClient{cc}
}

func (c *bookCategoriesClient) CreateCategory(ctx context.Context, in *NewCategoryRequest, opts ...grpc.CallOption) (*NewCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NewCategoryResponse)
	err := c.cc.Invoke(ctx, BookCategories_CreateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookCategoriesClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCategoryResponse)
	err := c.cc.Invoke(ctx, BookCategories_UpdateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookCategoriesClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCategoryResponse)
	err := c.cc.Invoke(ctx, BookCategories_DeleteCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookCategoriesServer is the server API for BookCategories service.
// All implementations should embed UnimplementedBookCategoriesServer
// for forward compatibility.
type BookCategoriesServer interface {
	CreateCategory(context.Context, *NewCategoryRequest) (*NewCategoryResponse, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error)
}

// UnimplementedBookCategoriesServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookCategoriesServer struct{}

func (UnimplementedBookCategoriesServer) CreateCategory(context.Context, *NewCategoryRequest) (*NewCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedBookCategoriesServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedBookCategoriesServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedBookCategoriesServer) testEmbeddedByValue() {}

// UnsafeBookCategoriesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookCategoriesServer will
// result in compilation errors.
type UnsafeBookCategoriesServer interface {
	mustEmbedUnimplementedBookCategoriesServer()
}

func RegisterBookCategoriesServer(s grpc.ServiceRegistrar, srv BookCategoriesServer) {
	// If the following call pancis, it indicates UnimplementedBookCategoriesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookCategories_ServiceDesc, srv)
}

func _BookCategories_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookCategoriesServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookCategories_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookCategoriesServer).CreateCategory(ctx, req.(*NewCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookCategories_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookCategoriesServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookCategories_UpdateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookCategoriesServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookCategories_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookCategoriesServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookCategories_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookCategoriesServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookCategories_ServiceDesc is the grpc.ServiceDesc for BookCategories service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookCategories_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_store.admin_service.BookCategories",
	HandlerType: (*BookCategoriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCategory",
			Handler:    _BookCategories_CreateCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _BookCategories_UpdateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _BookCategories_DeleteCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}

const (
	SalesReport_GetReport_FullMethodName = "/book_store.admin_service.SalesReport/GetReport"
)

// SalesReportClient is the client API for SalesReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SalesReportClient interface {
	GetReport(ctx context.Context, in *SalesReportRequest, opts ...grpc.CallOption) (*SalesReportResponse, error)
}

type salesReportClient struct {
	cc grpc.ClientConnInterface
}

func NewSalesReportClient(cc grpc.ClientConnInterface) SalesReportClient {
	return &salesReportClient{cc}
}

func (c *salesReportClient) GetReport(ctx context.Context, in *SalesReportRequest, opts ...grpc.CallOption) (*SalesReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SalesReportResponse)
	err := c.cc.Invoke(ctx, SalesReport_GetReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SalesReportServer is the server API for SalesReport service.
// All implementations should embed UnimplementedSalesReportServer
// for forward compatibility.
type SalesReportServer interface {
	GetReport(context.Context, *SalesReportRequest) (*SalesReportResponse, error)
}

// UnimplementedSalesReportServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSalesReportServer struct{}

func (UnimplementedSalesReportServer) GetReport(context.Context, *SalesReportRequest) (*SalesReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (UnimplementedSalesReportServer) testEmbeddedByValue() {}

// UnsafeSalesReportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SalesReportServer will
// result in compilation errors.
type UnsafeSalesReportServer interface {
	mustEmbedUnimplementedSalesReportServer()
}

func RegisterSalesReportServer(s grpc.ServiceRegistrar, srv SalesReportServer) {
	// If the following call pancis, it indicates UnimplementedSalesReportServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SalesReport_ServiceDesc, srv)
}

func _SalesReport_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SalesReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesReportServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SalesReport_GetReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesReportServer).GetReport(ctx, req.(*SalesReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SalesReport_ServiceDesc is the grpc.ServiceDesc for SalesReport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SalesReport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_store.admin_service.SalesReport",
	HandlerType: (*SalesReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReport",
			Handler:    _SalesReport_GetReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}