package utils

type ERRORS string
type INFO string

const (
	AUTH_ERROR     ERRORS = "auth error (books service)"
	DB_ERROR       ERRORS = "db error (books service)"
	INTERNAL_ERROR ERRORS = "db internal error (books service)"

	DB_INFO     INFO = "db info (books service)"
	SERVER_INFO INFO = "server info (auth service)"
)
