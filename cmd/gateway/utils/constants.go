package utils

import "github.com/Surya-7890/book_store/gateway/gen"

type AUTH_ROUTES string
type INFO string

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

type AuthResponse interface {
	*gen.Admin | *gen.User
}

type ERRORS string

const (
	HANDLER_REGISTRATION_ERROR ERRORS = "handler registration error (auth service)"
	JWT_AUTHORIZATION_ERROR    ERRORS = "jwt authorization error (auth service)"
	REDIS_ERROR                ERRORS = "redis error (auth service)"

	SERVER_INFO INFO = "server info (auth service)"
)
