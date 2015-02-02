package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/l2x/wolffy/server/controllers"
	"github.com/martini-contrib/render"

	"github.com/martini-contrib/gzip"
)

func router() {

	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(martini.Static("web"))
	m.Use(render.Renderer())
	m.Use(func(c martini.Context, req *http.Request) {
		Query := req.URL.Query()
		c.Map(Query)
	})

	site := controllers.Site{}
	cluster := controllers.Cluster{}
	project := controllers.Project{}

	m.Get("/favicon.ico", func() {})
	m.Get("/", site.Index)

	m.Get("/cluster/", cluster.GetAll)
	m.Get("/project/add/", project.Add)

	m.RunOnAddr(":8000")
}
