package main

import (
	"net"

	"github.com/Surya-7890/book_store/books/config"
	"github.com/Surya-7890/book_store/books/db"
	"github.com/Surya-7890/book_store/books/gen"
	"github.com/Surya-7890/book_store/books/routes"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.LoadConfig()

	port := viper.GetString("port")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	DB := db.ConnectToPostgres()

	server := grpc.NewServer()
	gen.RegisterBooksServer(server, &routes.BooksService{DB: DB})
	gen.RegisterModifyBooksServer(server, &routes.ModifyBooksService{DB: DB})
	gen.RegisterModifyCategoriesServer(server, &routes.ModifyCategoriesService{DB: DB})

	reflection.Register(server)

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
