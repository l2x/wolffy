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

	m.Use(func(r render.Render, req http.ResponseWriter, res *http.Request) {
		err := controllers.CheckSession(req, res)
		if err != nil {
			result := controllers.NewRes()
			result.Errno = 1001
			//controllers.RenderError(r, result, err)
		}
	})

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
	m.Get("/project/search/", project.Search)

	cluster := controllers.Cluster{}
	m.Get("/cluster/add/", cluster.Add)
	m.Get("/cluster/del/", cluster.Del)
	m.Get("/cluster/get/", cluster.Get)
	m.Get("/cluster/getall/", cluster.GetAll)
	m.Get("/cluster/update/", cluster.Update)

	deploy := controllers.Deploy{}
	m.Get("/deploy/push/", deploy.Push)
	m.Get("/deploy/history/", deploy.History)

	user := controllers.User{}
	m.Get("/user/get/", user.Get)
	m.Get("/user/getall/", user.GetAll)
	m.Get("/user/delete/", user.Del)
	m.Get("/user/add/", user.Add)
	m.Get("/user/update/", user.Update)
	m.Get("/user/updatepassword/", user.UpdatePassword)
	m.Get("/user/login/", user.Login)
	m.Get("/user/logout/", user.Logout)

	machine := controllers.Machine{}
	m.Get("/machine/ping/", machine.Ping)

	m.RunOnAddr(":8000")
}
