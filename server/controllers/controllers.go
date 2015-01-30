package controllers

import (
	"github.com/l2x/wolffy/server/config"
)

type Res struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func NewRes() Res {
	return Res{
		Errno:  0,
		Errmsg: "",
		Data:   []int{},
	}
}

func init() {
	config.InitConfig("")
}
