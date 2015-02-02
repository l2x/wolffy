package controllers

import (
	"net/http"
	"strconv"

	"github.com/l2x/wolffy/utils/git"
	"github.com/martini-contrib/render"

	"github.com/l2x/wolffy/server/config"
	"github.com/l2x/wolffy/server/models"
)

type Project struct{}

func (c Project) GetTags(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	project, err := models.ProjectModel.GetOne(idint)
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	repo := git.NewRepository(config.BasePath, project.Path)
	err = repo.PullTags()
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	tags, err := repo.GetTags()
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}
	res.Errno = 0
	res.Data = tags
	r.JSON(200, res)
}

func (c Project) Get(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	project, err := models.ProjectModel.GetOne(idint)
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	res.Errno = 0
	res.Data = project

	r.JSON(200, res)
	return
}

func (c Project) Add(r render.Render, req *http.Request) {
	res := NewRes()

	//path := "git@123.57.75.209:leiyonglin/wolffy.git"
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

	project, err := models.ProjectModel.Add(pidint, name, path, note)
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	repo := git.NewRepository(config.BasePath, path)
	if _, err = repo.Clone(); err != nil {
		res.Errmsg = err.Error()
		models.ProjectModel.Del(project.Id)
	}

	res.Errno = 0
	res.Data = project
	r.JSON(200, res)
}

func (c Project) Del(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	if err = models.ProjectModel.Del(idint); err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	res.Errno = 0
	r.JSON(200, res)
}

func (c Project) Update(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	pid := req.URL.Query().Get("pid")
	name := req.URL.Query().Get("name")
	path := req.URL.Query().Get("path")
	note := req.URL.Query().Get("note")

	idint, err := strconv.Atoi(id)
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	pidint, err := strconv.Atoi(pid)
	if err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	if err := models.ProjectModel.Update(idint, pidint, name, path, note); err != nil {
		res.Errmsg = err.Error()
		r.JSON(200, res)
		return
	}

	res.Errno = 0
	r.JSON(200, res)
}
