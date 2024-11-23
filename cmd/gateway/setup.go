package main

import (
	"context"
	"os"
	"strings"

	"github.com/Surya-7890/book_store/gateway/config"
	"github.com/Surya-7890/book_store/gateway/gen"
	"github.com/Surya-7890/book_store/gateway/utils"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func setupAdminEndpoints(ctx context.Context, gw *gwruntime.ServeMux, dialOpts []grpc.DialOption, service *config.Service) {
	err := gen.RegisterAdminAuthHandlerFromEndpoint(context.WithoutCancel(ctx), gw, strings.Join([]string{service.Host, service.Port}, ":"), dialOpts)
	if err != nil {
		Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.HANDLER_REGISTRATION_ERROR),
			Value: []byte(err.Error()),
		})
		os.Exit(1)
	}
}

func setupBooksEndpoints(ctx context.Context, gw *gwruntime.ServeMux, dialOpts []grpc.DialOption, service *config.Service) {
	err := gen.RegisterBooksHandlerFromEndpoint(context.WithoutCancel(ctx), gw, strings.Join([]string{service.Host, service.Port}, ":"), dialOpts)
	if err != nil {
		Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.HANDLER_REGISTRATION_ERROR),
			Value: []byte(err.Error()),
		})
		os.Exit(1)
	}

	err = gen.RegisterModifyBooksHandlerFromEndpoint(context.WithoutCancel(ctx), gw, strings.Join([]string{service.Host, service.Port}, ":"), dialOpts)
	if err != nil {
		Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.HANDLER_REGISTRATION_ERROR),
			Value: []byte(err.Error()),
		})
		os.Exit(1)
	}
}

func setupUserEndpoints(ctx context.Context, gw *gwruntime.ServeMux, dialOpts []grpc.DialOption, service *config.Service) {
	err := gen.RegisterUserAuthHandlerFromEndpoint(context.WithoutCancel(ctx), gw, strings.Join([]string{service.Host, service.Port}, ":"), dialOpts)
	if err != nil {
		Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.HANDLER_REGISTRATION_ERROR),
			Value: []byte(err.Error()),
		})
		os.Exit(1)
	}
	err = gen.RegisterUserProfileHandlerFromEndpoint(context.WithoutCancel(ctx), gw, strings.Join([]string{service.Host, service.Port}, ":"), dialOpts)
	if err != nil {
		Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.HANDLER_REGISTRATION_ERROR),
			Value: []byte(err.Error()),
		})
		os.Exit(1)
	}
}

func setup(gw *gwruntime.ServeMux, app *config.Application) {
	ctx := context.Background()
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	setupAdminEndpoints(ctx, gw, dialOpts, &app.Admin)
	setupBooksEndpoints(ctx, gw, dialOpts, &app.Books)
	setupUserEndpoints(ctx, gw, dialOpts, &app.User)
}
