package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/l2x/wolffy/server/controllers"
	"github.com/l2x/wolffy/utils"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
)

func router() {
	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(render.Renderer())

	m.Use(func(r render.Render, w http.ResponseWriter, req *http.Request) {
		token := req.URL.Query().Get("token")
		sign := req.URL.Query().Get("sign")
		ok := CheckSign(token, sign)
		if !ok {
			result := controllers.NewRes()
			controllers.RenderError(r, result, errors.New("signature invalid"))
			return
		}
	})

	server := Server{}
	m.Post("/pull/", server.Pull)

	m.RunOnAddr(":8001")
}

func CheckSign(token, sign string) bool {
	return sign == utils.Md5(fmt.Sprintf("%s%s", token, PrivateKey))
}
