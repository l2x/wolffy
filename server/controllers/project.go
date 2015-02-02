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
	if err = RenderError(r, res, err); err != nil {
		return
	}

	project, err := models.ProjectModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	repo := git.NewRepository(config.BasePath, project.Path)
	err = repo.PullTags()
	if err = RenderError(r, res, err); err != nil {
		return
	}

	tags, err := repo.GetTags()
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, tags)
}

func (c Project) Get(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	project, err := models.ProjectModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, project)
}

func (c Project) Add(r render.Render, req *http.Request) {
	res := NewRes()

	//path := "git@123.57.75.209:leiyonglin/wolffy.git"
	pid := req.URL.Query().Get("pid")
	name := req.URL.Query().Get("name")
	path := req.URL.Query().Get("path")
	note := req.URL.Query().Get("note")

	pidint, err := strconv.Atoi(pid)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	project, err := models.ProjectModel.Add(pidint, name, path, note)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	repo := git.NewRepository(config.BasePath, path)
	_, err = repo.Clone()
	if err = RenderError(r, res, err); err != nil {
		models.ProjectModel.Del(project.Id)
		return
	}

	RenderRes(r, res, project)
}

func (c Project) Del(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	err = models.ProjectModel.Del(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, map[string]string{})
}

func (c Project) Update(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	pid := req.URL.Query().Get("pid")
	name := req.URL.Query().Get("name")
	path := req.URL.Query().Get("path")
	note := req.URL.Query().Get("note")

	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	pidint, err := strconv.Atoi(pid)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	err = models.ProjectModel.Update(idint, pidint, name, path, note)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, map[string]string{})
}
