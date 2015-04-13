package main

import (
	"fmt"
	"net/url"

	"github.com/l2x/wolffy/utils"
)

var (
	timeout = 10
)

func report() error {
	master := fmt.Sprintf("%s/%s", Master, "report")

	u, err := url.Parse(master)
	if err != nil {
		return err
	}

	token, sign := utils.GenSign(PrivateKey)
	q := u.Query()
	q.Set("token", token)
	q.Set("sign", sign)
	u.RawQuery = q.Encode()

	_, err = utils.HttpGet(u.String(), timeout)
	if err != nil {
		return err
	}

	//TODO check resp error

	return nil
}
