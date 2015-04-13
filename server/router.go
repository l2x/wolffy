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

	m.Use(func(r render.Render, w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/user/login" {
			return
		}

		err := controllers.CheckSession(w, req)
		if err != nil {
			result := controllers.NewRes()
			result.Errno = 401
			controllers.RenderError(r, result, err)
		}
	})

	site := controllers.Site{}
	m.Get("/favicon.ico", func() {})
	m.Get("/", site.Index)

	project := controllers.Project{}
	m.Get("/project/add", project.Add)
	m.Get("/project/get", project.Get)
	m.Get("/project/getall", project.GetAll)
	m.Get("/project/delete", project.Delete)
	m.Get("/project/update", project.Update)
	m.Get("/project/gettags", project.GetTags)

	cluster := controllers.Cluster{}
	m.Get("/cluster/add", cluster.Add)
	m.Get("/cluster/delete", cluster.Delete)
	m.Get("/cluster/get", cluster.Get)
	m.Get("/cluster/getall", cluster.GetAll)
	m.Get("/cluster/update", cluster.Update)

	deploy := controllers.Deploy{}
	m.Get("/deploy/push", deploy.Push)
	m.Get("/deploy/get", deploy.Get)
	m.Get("/deploy/history", deploy.History)
	m.Get("/deploy/historyDetail", deploy.HistoryDetail)
	m.Get("/deploy/addtag", deploy.AddTag)
	m.Get("/deploy/getdiff", deploy.GetDiff)

	user := controllers.User{}
	m.Get("/user/get", user.Get)
	m.Get("/user/getall", user.GetAll)
	m.Get("/user/delete", user.Del)
	m.Get("/user/edit", user.Edit)
	m.Get("/user/update", user.Update)
	m.Get("/user/updatepassword", user.UpdatePassword)
	m.Get("/user/login", user.Login)
	m.Get("/user/logout", user.Logout)
	m.Get("/user/getuserinfo", user.GetUserInfo)

	Node := controllers.Node{}
	m.Get("/Node/add", Node.Add)
	m.Get("/Node/getall", Node.GetAll)

	m.RunOnAddr(":8000")
}
