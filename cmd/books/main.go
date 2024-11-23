package main

import (
	"net"

	"github.com/Surya-7890/book_store/books/config"
	"github.com/Surya-7890/book_store/books/db"
	"github.com/Surya-7890/book_store/books/gen"
	"github.com/Surya-7890/book_store/books/kafka"
	"github.com/Surya-7890/book_store/books/routes"
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
	gen.RegisterBooksServer(server, &routes.BooksService{DB: DB})
	gen.RegisterModifyBooksServer(server, &routes.ModifyBooksService{DB: DB})

	reflection.Register(server)

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
