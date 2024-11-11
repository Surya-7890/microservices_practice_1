package main

import (
	"context"
	"strings"

	"github.com/Surya-7890/book_store/gateway/gen"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func setupAdminEndpoints(ctx context.Context, gw *gwruntime.ServeMux, dialOpts []grpc.DialOption) {
	admin_host := viper.GetString("admin.host")
	admin_port := viper.GetString("admin.port")

	err := gen.RegisterAdminAuthHandlerFromEndpoint(context.WithoutCancel(ctx), gw, strings.Join([]string{admin_host, admin_port}, ":"), dialOpts)
	if err != nil {
		panic(err)
	}
}

func setupBooksEndpoints(ctx context.Context, gw *gwruntime.ServeMux, dialOpts []grpc.DialOption) {
	books_host := viper.GetString("books.host")
	books_port := viper.GetString("books.port")

	err := gen.RegisterBooksHandlerFromEndpoint(context.WithoutCancel(ctx), gw, strings.Join([]string{books_host, books_port}, ":"), dialOpts)
	if err != nil {
		panic(err)
	}

	err = gen.RegisterModifyBooksHandlerFromEndpoint(context.WithoutCancel(ctx), gw, strings.Join([]string{books_host, books_port}, ":"), dialOpts)
	if err != nil {
		panic(err)
	}
}

func setupUserEndpoints(ctx context.Context, gw *gwruntime.ServeMux, dialOpts []grpc.DialOption) {
	user_host := viper.GetString("user.host")
	user_port := viper.GetString("user.port")

	err := gen.RegisterUserAuthHandlerFromEndpoint(context.WithoutCancel(ctx), gw, strings.Join([]string{user_host, user_port}, ":"), dialOpts)
	if err != nil {
		panic(err)
	}
	err = gen.RegisterUserCommentsHandlerFromEndpoint(context.WithoutCancel(ctx), gw, strings.Join([]string{user_host, user_port}, ":"), dialOpts)
	if err != nil {
		panic(err)
	}
	err = gen.RegisterUserProfileHandlerFromEndpoint(context.WithoutCancel(ctx), gw, strings.Join([]string{user_host, user_port}, ":"), dialOpts)
	if err != nil {
		panic(err)
	}
}

func setup(gw *gwruntime.ServeMux) {
	ctx := context.Background()
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	setupAdminEndpoints(ctx, gw, dialOpts)
	setupBooksEndpoints(ctx, gw, dialOpts)
	setupUserEndpoints(ctx, gw, dialOpts)
}
