package routes

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/admin/gen"
)

type AdminAuthService struct {
	gen.UnimplementedAdminAuthServer
}

/* POST: /v1/admin/login */
func (a *AdminAuthService) AdminLogin(ctx context.Context, req *gen.AdminLoginRequest) (*gen.AdminLoginResponse, error) {
	res := &gen.AdminLoginResponse{}
	fmt.Println("admin login")
	return res, nil
}

/* POST: /v1/admin/create */
func (a *AdminAuthService) AdminCreate(context.Context, *gen.AdminCreateRequest) (*gen.AdminCreateResponse, error) {
	res := &gen.AdminCreateResponse{}
	fmt.Println("admin create")
	return res, nil
}