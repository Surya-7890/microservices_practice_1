package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Surya-7890/book_store/books/config"
	"github.com/Surya-7890/book_store/books/db"
	"github.com/Surya-7890/book_store/books/gen"
	"github.com/Surya-7890/book_store/books/kafka"
	"github.com/Surya-7890/book_store/books/routes"
	"github.com/Surya-7890/book_store/books/utils"
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
	gen.RegisterBooksServer(server, &routes.BooksService{
		DB:    DB,
		Kafka: App.Kafka,
	})
	gen.RegisterModifyBooksServer(server, &routes.ModifyBooksService{
		DB:    DB,
		Kafka: App.Kafka,
	})

	reflection.Register(server)

	err_ := App.Kafka.Info.WriteMessages(context.Background(), _kafka.Message{
		Key:   []byte(utils.SERVER_INFO),
		Value: []byte("[books-service]: running server... on port: " + App.Port),
	})
	if err_ != nil {
		fmt.Println(err_.Error())
	}

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
