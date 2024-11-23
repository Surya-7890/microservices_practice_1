package main

import (
	"fmt"
	"net"

	"github.com/Surya-7890/book_store/admin/config"
	"github.com/Surya-7890/book_store/admin/db"
	"github.com/Surya-7890/book_store/admin/gen"
	"github.com/Surya-7890/book_store/admin/kafka"
	"github.com/Surya-7890/book_store/admin/routes"
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

	DB := db.ConnectToPostgres(&App.DatabaseConfig)
	fmt.Println(App.Kafka.Error.Addr)
	fmt.Println(App.Kafka.Info.Addr)
	fmt.Println(App.Kafka.Warning.Addr)

	server := grpc.NewServer()
	gen.RegisterAdminAuthServer(server, &routes.AdminAuthService{DB: DB})

	reflection.Register(server)

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
