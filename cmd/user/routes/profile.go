package routes

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/Surya-7890/book_store/user/db"
	"github.com/Surya-7890/book_store/user/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return res, status.Error(codes.InvalidArgument, "invalid header")
	}

	errors := md.Get("auth-error")
	if len(errors) != 0 {
		return res, status.Error(codes.PermissionDenied, strings.Join(errors, ", "))
	}

	role := md.Get("role")
	if role[0] != "admin" {
		return res, status.Error(codes.Unauthenticated, "operation not permitted")
	}

	user_string := md.Get("user")[0]
	if user_string == "" {
		return res, status.Error(codes.Unauthenticated, "invalid user")
	}

	var user db.User

	err := json.Unmarshal([]byte(user_string), &user)
	if err != nil {
		return res, status.Error(codes.Unauthenticated, "invalid user found")
	}

	res.Username = user.Username
	res.Name = user.Name
	res.Age = user.Age

	return res, nil
}

/* DELETE: /v1/user/profile */
func (u *UserService) DeleteUser(ctx context.Context, req *gen.DeleteUserRequest) (*gen.DeleteUserResponse, error) {
	res := &gen.DeleteUserResponse{}
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return res, status.Error(codes.InvalidArgument, "invalid header")
	}

	errors := md.Get("auth-error")
	if len(errors) != 0 {
		return res, status.Error(codes.PermissionDenied, strings.Join(errors, ", "))
	}

	role := md.Get("role")
	if role[0] != "admin" {
		return res, status.Error(codes.Unauthenticated, "operation not permitted")
	}

	user_string := md.Get("user")[0]
	if user_string == "" {
		return res, status.Error(codes.Unauthenticated, "invalid user")
	}

	var user db.User

	err := json.Unmarshal([]byte(user_string), &user)
	if err != nil {
		return res, status.Errorf(codes.Unauthenticated, "invalid user found %s", err.Error())
	}

	if user.Username == "" {
		return res, status.Error(codes.Unauthenticated, "error with the user found inside token")
	}

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
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return res, status.Error(codes.InvalidArgument, "invalid header")
	}

	errors := md.Get("auth-error")
	if len(errors) != 0 {
		return res, status.Error(codes.PermissionDenied, strings.Join(errors, ", "))
	}

	role := md.Get("role")
	if role[0] != "admin" {
		return res, status.Error(codes.Unauthenticated, "operation not permitted")
	}

	user_string := md.Get("user")[0]
	if user_string == "" {
		return res, status.Error(codes.Unauthenticated, "invalid user")
	}

	var user db.User

	err := json.Unmarshal([]byte(user_string), &user)
	if err != nil {
		return res, status.Error(codes.Unauthenticated, "invalid user found")
	}

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
