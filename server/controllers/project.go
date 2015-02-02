package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/l2x/wolffy/utils/git"
	"github.com/martini-contrib/render"

	"github.com/l2x/wolffy/server/config"
	"github.com/l2x/wolffy/server/models"
)

type Project struct{}

func (c Project) Add(r render.Render, req *http.Request, q martini.Context) {
	res := NewRes()
	fmt.Println(q)

	//remotePath := "git@123.57.75.209:leiyonglin/wolffy.git"
	remotePath := req.URL.Query().Get("remotepath")
	pid := req.URL.Query().Get("pid")
	name := req.URL.Query().Get("name")
	path := req.URL.Query().Get("path")
	note := req.URL.Query().Get("note")

	pidint, err := strconv.Atoi(pid)
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	_, err = models.ProjectModel.Add(pidint, name, path, note)
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	repo := git.NewRepository(config.BasePath, remotePath)
	_, err = repo.Clone()
	if err != nil {
		res.Errmsg = err.Error()
	}

	r.JSON(200, res)
}
