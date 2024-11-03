package routes

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/admin/gen"
)

type AdminService struct {
	gen.UnimplementedAdminAuthServer
	gen.UnimplementedAdminBooksServer
}

func (a *AdminService) AdminLogin(ctx context.Context, req *gen.AdminLoginRequest) (*gen.AdminLoginResponse, error) {
	res := &gen.AdminLoginResponse{}
	fmt.Println("hello")
	return res, nil
}