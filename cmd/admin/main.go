package main

import (
	"net"

	"github.com/Surya-7890/book_store/admin/config"
	"github.com/Surya-7890/book_store/admin/gen"
	"github.com/Surya-7890/book_store/admin/routes"
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

	server := grpc.NewServer()
	gen.RegisterAdminAuthServer(server, &routes.AdminAuthService{})
	gen.RegisterAdminBooksServer(server, &routes.AdminBooksService{})
	gen.RegisterBookCategoriesServer(server, &routes.BookCategoryService{})
	gen.RegisterSalesReportServer(server, &routes.SalesReportService{})

	reflection.Register(server)

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}