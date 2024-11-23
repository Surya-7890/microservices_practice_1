package main

import (
	"fmt"
	"net"

	"github.com/Surya-7890/book_store/user/config"
	"github.com/Surya-7890/book_store/user/db"
	"github.com/Surya-7890/book_store/user/gen"
	"github.com/Surya-7890/book_store/user/kafka"
	"github.com/Surya-7890/book_store/user/routes"
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

	fmt.Println(App.DatabaseConfig.DBName)

	DB := db.ConnectToPostgres(App.Kafka, &App.DatabaseConfig)

	server := grpc.NewServer()
	gen.RegisterUserAuthServer(server, &routes.UserService{
		DB:    DB,
		Kafka: App.Kafka,
	})
	gen.RegisterUserProfileServer(server, &routes.UserProfileService{
		DB:    DB,
		Kafka: App.Kafka,
	})

	reflection.Register(server)

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
