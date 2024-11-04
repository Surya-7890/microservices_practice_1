package routes

import (
	"context"
	"fmt"

	"github.com/Surya-7890/book_store/user/gen"
)

type UserService struct {
	gen.UnimplementedUserAuthServer
	gen.UnimplementedUserCommentsServer
	gen.UnimplementedUserProfileServer
}

/* POST: /v1/user/login */
func (u *UserService) UserLogin(ctx context.Context, req *gen.UserLoginRequest) (*gen.UserLoginResponse, error) {
	res := &gen.UserLoginResponse{}
	fmt.Println("user login")
	return res, nil
}

/* POST: /v1/user/signup */
func (u *UserService) UserSignup(ctx context.Context, req *gen.UserSignupRequest) (*gen.UserSignupResponse, error) {
	res := &gen.UserSignupResponse{}
	fmt.Println("user signup")
	return res, nil
}