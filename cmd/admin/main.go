package main

import (
	"context"
	"net"

	"github.com/Surya-7890/book_store/admin/config"
	"github.com/Surya-7890/book_store/admin/db"
	"github.com/Surya-7890/book_store/admin/gen"
	"github.com/Surya-7890/book_store/admin/kafka"
	"github.com/Surya-7890/book_store/admin/routes"
	"github.com/Surya-7890/book_store/admin/utils"
	_kafka "github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	App *config.Application
)

func init() {
	App = config.LoadConfig()
	App.Kafka = kafka.CreateWriters(&App.KafkaConfig)
}

func main() {
	listener, err := net.Listen("tcp", App.Port)
	if err != nil {
		panic(err)
	}

	DB := db.ConnectToPostgres(App.Kafka, &App.DatabaseConfig)

	server := grpc.NewServer()
	gen.RegisterAdminAuthServer(server, &routes.AdminAuthService{DB: DB})

	reflection.Register(server)
	App.Kafka.Info.WriteMessages(context.Background(), _kafka.Message{
		Key:   []byte(utils.SERVER_INFO),
		Value: []byte("[admin-service]: running server... on port: " + App.Port),
	})
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
