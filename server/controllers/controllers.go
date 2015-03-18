package controllers

import "github.com/martini-contrib/render"

func InitControllers() error {
	NewSession()

	return nil
}

//返回数据格式
type Res struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func NewRes() Res {
	return Res{
		Errno:  1,
		Errmsg: "",
		Data:   []int{},
	}
}

//标准输出
func RenderError(r render.Render, errno int, data interface{}) {
	res := NewRes()
	res.Errno = errno
	res.Errmsg = config.GetErr(errno)
	res.Data = data
	r.JSON(200, res)
}

func RenderRes(r render.Render, data interface{}) {
	res.Errno = 0
	res.Data = data
	r.JSON(200, res)
}
