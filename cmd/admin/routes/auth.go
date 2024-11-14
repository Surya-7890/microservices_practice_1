package routes

import (
	"context"

	"github.com/Surya-7890/book_store/admin/db"
	"github.com/Surya-7890/book_store/admin/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AdminAuthService struct {
	gen.UnimplementedAdminAuthServer
	DB *gorm.DB
}

/* POST: /v1/admin/login */
func (a *AdminAuthService) AdminLogin(ctx context.Context, req *gen.AdminLoginRequest) (*gen.AdminLoginResponse, error) {
	res := &gen.AdminLoginResponse{}
	admin := db.Admin{}

	a.DB.Where("username = ?", req.GetUsername()).First(&admin)

	if !admin.IsCorrectPassword(req.GetPassword()) {
		res.Status = RESPONSE_FAILURE
		return res, status.Error(codes.PermissionDenied, "Incorrect Password")
	}
	res.Status = RESPONSE_SUCCESS

	return res, nil
}

/* POST: /v1/admin/create */
func (a *AdminAuthService) AdminCreate(ctx context.Context, req *gen.AdminCreateRequest) (*gen.AdminCreateResponse, error) {
	res := &gen.AdminCreateResponse{}

	admin := db.Admin{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
	if admin.AlreadyExists(a.DB) {
		res.Status = RESPONSE_FAILURE
		return res, status.Error(codes.AlreadyExists, "Admin With The Username Already Exists")
	}
	tx := a.DB.Create(&admin)

	// handle errors while creating
	if tx.Error != nil {
		res.Status = RESPONSE_FAILURE
		return res, status.Errorf(codes.Internal, "Unable To Create Admin: %s", tx.Error.Error())
	}

	// handle rows affected on creation
	if tx.RowsAffected == 0 {
		res.Status = RESPONSE_FAILURE
		return res, status.Error(codes.Unknown, "Unable To Create Admin")
	}
	res.Status = RESPONSE_SUCCESS

	ctx = context.WithValue(ctx, "token", "token")
	ctx = context.WithValue(ctx, "username", admin.Username)
	ctx = context.WithValue(ctx, "user_id", admin.ID)

	return res, nil
}
