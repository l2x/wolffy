package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/l2x/wolffy/master/controllers"
	"github.com/l2x/wolffy/utils"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
)

func router() {
	martini.Env = "production"

	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(render.Renderer())

	m.Use(func(r render.Render, w http.ResponseWriter, req *http.Request) {
		token := req.URL.Query().Get("token")
		sign := req.URL.Query().Get("sign")
		err := utils.CheckSign(token, sign, PrivateKey)
		if err != nil {
			result := controllers.NewRes()
			controllers.RenderError(r, result, err)
			return
		}
	})

	server := Server{}
	m.Post("/pull", server.Pull)

	m.RunOnAddr(Port)
}
