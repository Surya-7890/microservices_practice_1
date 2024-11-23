package utils

type ERRORS string

const (
	AUTH_ERROR   ERRORS = "auth error (auth service)"
	ADMIN_CREATE ERRORS = "admin created successfully (auth service)"
	ADMIN_LOGIN  ERRORS = "admin login successful (auth service)"

	DB_ERROR ERRORS = "db error (auth service)"
)
