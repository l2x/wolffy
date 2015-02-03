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
	m.Get("/favicon.ico", func() {})
	m.Get("/", site.Index)

	project := controllers.Project{}
	m.Get("/project/add/", project.Add)
	m.Get("/project/get/", project.Get)
	m.Get("/project/del/", project.Del)
	m.Get("/project/update/", project.Update)
	m.Get("/project/gettags/", project.GetTags)
	m.Get("/project/diff/", project.Diff)

	product := controllers.Product{}
	m.Get("/product/add/", product.Add)
	m.Get("/product/get/", product.Get)
	m.Get("/product/getall/", product.GetAll)
	m.Get("/product/del/", product.Del)
	m.Get("/product/update/", product.Update)

	cluster := controllers.Cluster{}
	m.Get("/cluster/add/", cluster.Add)
	m.Get("/cluster/del/", cluster.Del)
	m.Get("/cluster/get/", cluster.Get)
	m.Get("/cluster/getall/", cluster.GetAll)
	m.Get("/cluster/update/", cluster.Update)

	deploy := controllers.Deploy{}
	m.Get("/deploy/push/", deploy.Push)
	m.Get("/deploy/history/", deploy.History)

	m.RunOnAddr(":8000")
}
