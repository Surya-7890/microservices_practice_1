package main

import (
	"net"

	"github.com/Surya-7890/book_store/admin/config"
	"github.com/Surya-7890/book_store/admin/gen"
	"github.com/Surya-7890/book_store/admin/routes"
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
	gen.RegisterAdminAuthServer(server, &routes.AdminService{})
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}