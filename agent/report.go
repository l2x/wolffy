package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"encoding/json"

	"github.com/l2x/wolffy/server/controllers"
	"github.com/l2x/wolffy/utils"
)

var (
	timeout = 10
)

func report() error {
	master := fmt.Sprintf("http://%s/%s", strings.TrimLeft(Master, "http://"), "node/report")

	u, err := url.Parse(master)
	if err != nil {
		return err
	}

	token, sign := utils.GenSign(PrivateKey)
	q := u.Query()
	q.Set("token", token)
	q.Set("sign", sign)
	q.Set("port", Port)
	u.RawQuery = q.Encode()

	resp, err := utils.HttpGet(u.String(), timeout)
	if err != nil {
		return err
	}

	res := controllers.NewRes()
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return errors.New(err.Error() + "\n" + string(resp))
	}
	if res.Errno != 0 {
		return errors.New(res.Errmsg)
	}

	return nil
}
