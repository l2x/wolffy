package config

import "errors"

var (
	ERR_OK      = 0
	ERR_UNKNOWN = 1000

	ERR_SESSION_INVALID = 1001
	ERR_SESSION_EXPIRE  = 1002

	ERR_USER_NEED_CHANGE_PWD    = 2001
	ERR_USER_NOT_ADMIN          = 2002
	ERR_USER_NOT_FOUND          = 2003
	ERR_USER_PASSWORD_INCORRECT = 2004

	ERR_PROJECT_DEPLOYING     = 3001
	ERR_PROJECT_CLUSTER_EMPTY = 3002

	ERR = map[int]string{
		ERR_OK:                      "ok",
		ERR_UNKNOWN:                 "unknown error",
		ERR_SESSION_INVALID:         "session invalid",
		ERR_USER_NEED_CHANGE_PWD:    "need change password",
		ERR_USER_NOT_ADMIN:          "administrator required",
		ERR_USER_NOT_FOUND:          "user not found",
		ERR_USER_PASSWORD_INCORRECT: "password incorrect",

		ERR_PROJECT_DEPLOYING:     "project deploying",
		ERR_PROJECT_CLUSTER_EMPTY: "project clusters empty",
	}
)

func GetErr(errno int) error {
	e, ok := ERR[errno]
	if !ok {
		e = ERR[ERR_UNKNOWN]
	}

	return errors.New(e)
}
