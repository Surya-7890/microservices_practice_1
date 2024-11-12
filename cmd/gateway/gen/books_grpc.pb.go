// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: books.proto

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
	Books_GetBooks_FullMethodName = "/book_store.books_service.Books/GetBooks"
	Books_GetBook_FullMethodName  = "/book_store.books_service.Books/GetBook"
)

// BooksClient is the client API for Books service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BooksClient interface {
	GetBooks(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (*GetBooksResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
}

type booksClient struct {
	cc grpc.ClientConnInterface
}

func NewBooksClient(cc grpc.ClientConnInterface) BooksClient {
	return &booksClient{cc}
}

func (c *booksClient) GetBooks(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (*GetBooksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBooksResponse)
	err := c.cc.Invoke(ctx, Books_GetBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, Books_GetBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooksServer is the server API for Books service.
// All implementations should embed UnimplementedBooksServer
// for forward compatibility.
type BooksServer interface {
	GetBooks(context.Context, *GetBooksRequest) (*GetBooksResponse, error)
	GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error)
}

// UnimplementedBooksServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBooksServer struct{}

func (UnimplementedBooksServer) GetBooks(context.Context, *GetBooksRequest) (*GetBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (UnimplementedBooksServer) GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBooksServer) testEmbeddedByValue() {}

// UnsafeBooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BooksServer will
// result in compilation errors.
type UnsafeBooksServer interface {
	mustEmbedUnimplementedBooksServer()
}

func RegisterBooksServer(s grpc.ServiceRegistrar, srv BooksServer) {
	// If the following call pancis, it indicates UnimplementedBooksServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Books_ServiceDesc, srv)
}

func _Books_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Books_GetBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServer).GetBooks(ctx, req.(*GetBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Books_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Books_GetBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Books_ServiceDesc is the grpc.ServiceDesc for Books service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Books_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_store.books_service.Books",
	HandlerType: (*BooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBooks",
			Handler:    _Books_GetBooks_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Books_GetBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books.proto",
}

const (
	ModifyBooks_NewBook_FullMethodName     = "/book_store.books_service.ModifyBooks/NewBook"
	ModifyBooks_UpdateBooks_FullMethodName = "/book_store.books_service.ModifyBooks/UpdateBooks"
	ModifyBooks_DeleteBooks_FullMethodName = "/book_store.books_service.ModifyBooks/DeleteBooks"
)

// ModifyBooksClient is the client API for ModifyBooks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModifyBooksClient interface {
	NewBook(ctx context.Context, in *NewBookRequest, opts ...grpc.CallOption) (*NewBookResponse, error)
	UpdateBooks(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error)
	DeleteBooks(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error)
}

type modifyBooksClient struct {
	cc grpc.ClientConnInterface
}

func NewModifyBooksClient(cc grpc.ClientConnInterface) ModifyBooksClient {
	return &modifyBooksClient{cc}
}

func (c *modifyBooksClient) NewBook(ctx context.Context, in *NewBookRequest, opts ...grpc.CallOption) (*NewBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NewBookResponse)
	err := c.cc.Invoke(ctx, ModifyBooks_NewBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modifyBooksClient) UpdateBooks(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBookResponse)
	err := c.cc.Invoke(ctx, ModifyBooks_UpdateBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modifyBooksClient) DeleteBooks(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBookResponse)
	err := c.cc.Invoke(ctx, ModifyBooks_DeleteBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModifyBooksServer is the server API for ModifyBooks service.
// All implementations should embed UnimplementedModifyBooksServer
// for forward compatibility.
type ModifyBooksServer interface {
	NewBook(context.Context, *NewBookRequest) (*NewBookResponse, error)
	UpdateBooks(context.Context, *UpdateBookRequest) (*UpdateBookResponse, error)
	DeleteBooks(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error)
}

// UnimplementedModifyBooksServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedModifyBooksServer struct{}

func (UnimplementedModifyBooksServer) NewBook(context.Context, *NewBookRequest) (*NewBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBook not implemented")
}
func (UnimplementedModifyBooksServer) UpdateBooks(context.Context, *UpdateBookRequest) (*UpdateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBooks not implemented")
}
func (UnimplementedModifyBooksServer) DeleteBooks(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBooks not implemented")
}
func (UnimplementedModifyBooksServer) testEmbeddedByValue() {}

// UnsafeModifyBooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModifyBooksServer will
// result in compilation errors.
type UnsafeModifyBooksServer interface {
	mustEmbedUnimplementedModifyBooksServer()
}

func RegisterModifyBooksServer(s grpc.ServiceRegistrar, srv ModifyBooksServer) {
	// If the following call pancis, it indicates UnimplementedModifyBooksServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ModifyBooks_ServiceDesc, srv)
}

func _ModifyBooks_NewBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModifyBooksServer).NewBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModifyBooks_NewBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModifyBooksServer).NewBook(ctx, req.(*NewBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModifyBooks_UpdateBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModifyBooksServer).UpdateBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModifyBooks_UpdateBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModifyBooksServer).UpdateBooks(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModifyBooks_DeleteBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModifyBooksServer).DeleteBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModifyBooks_DeleteBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModifyBooksServer).DeleteBooks(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModifyBooks_ServiceDesc is the grpc.ServiceDesc for ModifyBooks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModifyBooks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_store.books_service.ModifyBooks",
	HandlerType: (*ModifyBooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewBook",
			Handler:    _ModifyBooks_NewBook_Handler,
		},
		{
			MethodName: "UpdateBooks",
			Handler:    _ModifyBooks_UpdateBooks_Handler,
		},
		{
			MethodName: "DeleteBooks",
			Handler:    _ModifyBooks_DeleteBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books.proto",
}