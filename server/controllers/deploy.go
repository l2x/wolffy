package controllers

import (
	"net/http"
	"strconv"

	"github.com/l2x/wolffy/server/models"
	"github.com/martini-contrib/render"
)

type Deploy struct{}

func (c Deploy) Push(r render.Render, req *http.Request) {
	res := NewRes()

	pid := req.URL.Query().Get("pid")
	commit := req.URL.Query().Get("commit")
	pidint, err := strconv.Atoi(pid)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	deploy, err := models.DeployModel.Add(pidint, commit)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	//TODO push code to agent

	RenderRes(r, res, deploy)
}

func (c Deploy) History(r render.Render, req *http.Request) {
	res := NewRes()

	pid := req.URL.Query().Get("pid")
	pidint, err := strconv.Atoi(pid)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	deploys, err := models.DeployModel.GetAll(pidint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, deploys)
}
