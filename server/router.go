package main

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-martini/martini"
	"github.com/l2x/wolffy/server/config"
	"github.com/l2x/wolffy/server/controllers"
	"github.com/martini-contrib/render"

	"github.com/martini-contrib/gzip"
)

var (
	NOT_CHECKSESSION = []string{
		"/user/login",
		"/node/report",
	}
)

func router() {
	martini.Env = "production"

	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(martini.Static("web"))
	m.Use(render.Renderer())

	m.Use(func(r render.Render, w http.ResponseWriter, req *http.Request) {
		for _, v := range NOT_CHECKSESSION {
			if req.URL.Path == v {
				return
			}
		}

		err := controllers.CheckSession(w, req)
		if err != nil {
			result := controllers.NewRes()
			result.Errno = 401
			controllers.RenderError(r, result, err)
		}
	})

	//
	m.Group("/debug/pprof", func(r martini.Router) {
		r.Any("/", pprof.Index)
		r.Any("/cmdline", pprof.Cmdline)
		r.Any("/profile", pprof.Profile)
		r.Any("/symbol", pprof.Symbol)
		r.Any("/block", pprof.Handler("block").ServeHTTP)
		r.Any("/heap", pprof.Handler("heap").ServeHTTP)
		r.Any("/goroutine", pprof.Handler("goroutine").ServeHTTP)
		r.Any("/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
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
	m.Get("/node/report", Node.Report)
	m.Get("/node/getall", Node.GetAll)
	m.Get("/node/delete", Node.Delete)
	m.Get("/node/getprivatekey", Node.GetPrivateKey)

	m.RunOnAddr(config.Port)
}
