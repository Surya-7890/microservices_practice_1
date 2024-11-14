package routes

import (
	"context"

	"github.com/Surya-7890/book_store/user/db"
	"github.com/Surya-7890/book_store/user/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserProfileService struct {
	gen.UnimplementedUserProfileServer
	DB *gorm.DB
}

/* GET: /v1/user/profile */
func (u *UserService) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.GetUserResponse, error) {
	res := &gen.GetUserResponse{}
	user := ctx.Value("user").(*db.User)

	res.Username = user.Username
	res.Name = user.Name
	res.Age = user.Age

	return res, nil
}

/* DELETE: /v1/user/profile */
func (u *UserService) DeleteUser(ctx context.Context, req *gen.DeleteUserRequest) (*gen.DeleteUserResponse, error) {
	res := &gen.DeleteUserResponse{}
	user := ctx.Value("user").(*db.User)

	if err := u.DB.Where("username = ?", user.Username).Delete(&db.User{}).Error; err != nil {
		res.Status = RESPONSE_FAILURE
		return res, status.Errorf(codes.Internal, "error while deleting user %s", err.Error())
	}

	res.Status = RESPONSE_SUCCESS
	return res, nil
}

/* PATCH: /v1/user/profile */
func (u *UserService) UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.UpdateUserResponse, error) {
	res := &gen.UpdateUserResponse{}
	user := ctx.Value("user").(*db.User)

	username := req.GetUsername()
	name := req.GetName()
	age := req.GetAge()

	updates := map[string]interface{}{}

	if len(username) > 0 {
		updates["username"] = username
	}

	if len(name) > 0 {
		updates["name"] = name
	}

	if age != 0 {
		updates["age"] = age
	}

	if err := u.DB.Where("username = ?", user.Username).Updates(updates).Error; err != nil {
		res.Status = RESPONSE_FAILURE
		return res, status.Errorf(codes.Internal, "error while updating user %s", err.Error())
	}

	res.Status = RESPONSE_SUCCESS

	return res, nil
}
