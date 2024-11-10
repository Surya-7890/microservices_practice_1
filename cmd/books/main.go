package main

import (
	"net"

	"github.com/Surya-7890/book_store/books/config"
	"github.com/Surya-7890/book_store/books/gen"
	"github.com/Surya-7890/book_store/books/routes"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	config.LoadConfig()

	port := viper.GetString("port")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	gen.RegisterBooksServer(server, &routes.BooksService{})
	gen.RegisterModifyBooksServer(server, &routes.ModifyBooksService{})

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
