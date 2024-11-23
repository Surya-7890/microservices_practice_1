package utils

type ERRORS string
type INFO string

const (
	AUTH_ERROR     ERRORS = "auth error (user service)"
	DB_ERROR       ERRORS = "db error (user service)"
	REQUEST_ERROR  ERRORS = "invalid request params (user service)"
	INTERNAL_ERROR ERRORS = "internal error (user service)"

	USER_CREATE INFO = "user created successfully (user service)"
	USER_LOGIN  INFO = "user login successful (user service)"
	USER_INFO   INFO = "user info (user service)"
	DB_INFO     INFO = "db info (user service)"
)
