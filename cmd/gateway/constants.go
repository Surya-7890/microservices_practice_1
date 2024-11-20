package main

import "github.com/Surya-7890/book_store/gateway/gen"

type AUTH_ROUTES string

const (
	AUTH_ERROR = "auth-error"
)

var AuthRoutes = [4]AUTH_ROUTES{
	"v1/admin/login",
	"v1/admin/create",
	"v1/user/login",
	"v1/user/signup",
}

const (
	AdminRoutes = "/v1/books"
	ROLE_ADMIN  = "admin"
	ROLE_USER   = "user"
)

type Middleware struct {
	Key string
}

type AuthResponse interface {
	*gen.Admin | *gen.User
}
