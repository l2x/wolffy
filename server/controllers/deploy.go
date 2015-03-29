package controllers

import (
	"net/http"
	"strconv"

	"github.com/l2x/wolffy/server/config"
	"github.com/l2x/wolffy/server/models"
	"github.com/l2x/wolffy/utils/git"
	"github.com/martini-contrib/render"
)

type Deploy struct{}

func (c Deploy) Push(r render.Render, req *http.Request) {
	res := NewRes()

	id := req.URL.Query().Get("id")
	commit := req.URL.Query().Get("commit")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	deploy, err := models.DeployModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	project, err := models.ProjectModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	repo := git.NewRepository(config.RepoPath, project.Path)
	err = repo.Archive(commit, repo.Path)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	//TODO push code to agent

	//finish
	err = models.DeployModel.UpdateStatus(deploy.Id, 2)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, deploy)
}

func (c Deploy) Get(r render.Render, req *http.Request) {
	res := NewRes()

	id := req.URL.Query().Get("id")
	idint, _ := strconv.Atoi(id)

	deploy, err := models.DeployModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, deploy)
}

func (c Deploy) AddTag(r render.Render, req *http.Request) {
	res := NewRes()
	tag := req.URL.Query().Get("tag")
	btag := req.URL.Query().Get("btag")
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	project, err := models.ProjectModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	var diff string
	if btag != "" {
		repo := git.NewRepository(config.RepoPath, project.Path)
		diff, err = repo.Diff(tag, btag)
		if err = RenderError(r, res, err); err != nil {
			return
		}
	}

	deploy, err := models.DeployModel.Add(idint, tag, diff)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, deploy)
}

func (c Deploy) History(r render.Render, req *http.Request) {
	res := NewRes()

	pid := req.URL.Query().Get("id")
	pidint, err := strconv.Atoi(pid)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	deploys, err := models.DeployModel.GetAll(pidint, 50)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, deploys)
}
