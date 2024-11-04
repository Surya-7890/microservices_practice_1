package routes

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/user/gen"
)

/* GET: /v1/user/profile */
func (u *UserService) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.GetUserResponse, error) {
	res := &gen.GetUserResponse{}
	fmt.Println("get user")
	return res, nil
}

/* DELETE: /v1/user/profile */
func (u *UserService) DeleteUser(ctx context.Context, req *gen.DeleteUserRequest) (*gen.DeleteUserResponse, error) {
	res := &gen.DeleteUserResponse{}
	fmt.Println("delete user")
	return res, nil
}

/* PATCH: /v1/user/profile */
func (u *UserService) UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.UpdateUserResponse, error) {
	res := &gen.UpdateUserResponse{}
	fmt.Println("update user")
	return res, nil
}