package main

import (
	"net"

	"github.com/Surya-7890/book_store/user/config"
	"github.com/Surya-7890/book_store/user/gen"
	"github.com/Surya-7890/book_store/user/routes"
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
	gen.RegisterUserAuthServer(server, &routes.UserService{})
	gen.RegisterUserCommentsServer(server, &routes.UserService{})
	gen.RegisterUserProfileServer(server, &routes.UserService{})

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}