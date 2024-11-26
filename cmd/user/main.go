package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Surya-7890/book_store/user/config"
	"github.com/Surya-7890/book_store/user/db"
	"github.com/Surya-7890/book_store/user/gen"
	"github.com/Surya-7890/book_store/user/kafka"
	"github.com/Surya-7890/book_store/user/routes"
	"github.com/Surya-7890/book_store/user/utils"
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

	err_ := App.Kafka.Info.WriteMessages(context.Background(), _kafka.Message{
		Key:   []byte(utils.SERVER_INFO),
		Value: []byte("[user-service]: running server... on port: " + App.Port),
	})
	if err_ != nil {
		fmt.Println(err_.Error())
	}

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
