package main

import (
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

	site := controllers.Site{}
	cluster := controllers.Cluster{}
	project := controllers.Project{}

	m.Get("/favicon.ico", func() {})
	m.Get("/", site.Index)

	m.Get("/cluster/", cluster.GetAll)
	m.Get("/project/add/", project.Add)

	m.RunOnAddr(":8000")
}
