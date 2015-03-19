package config

import "errors"

var (
	ERR_OK      = 0
	ERR_UNKNOWN = 1000

	ERR_SESSION_INVALID = 1001

	ERR_USER_NEED_CHANGE_PWD = 2001
	ERR_USER_NOT_ADMIN       = 2002

	ERR = map[int]string{
		ERR_OK:                   "ok",
		ERR_UNKNOWN:              "unknown error",
		ERR_SESSION_INVALID:      "session invalid",
		ERR_USER_NEED_CHANGE_PWD: "need change password",
	}
)

func GetErr(errno int) error {
	e, ok := ERR[errno]
	if !ok {
		e = ERR[ERR_UNKNOWN]
	}

	return errors.New(e)
}
