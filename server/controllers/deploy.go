package controllers

import (
	"errors"
	"fmt"
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
	pid := req.URL.Query().Get("pid")
	commit := req.URL.Query().Get("commit")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}
	pidint, err := strconv.Atoi(pid)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	deploy, err := models.DeployModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	if deploy.Status == 1 {
		RenderError(r, res, errors.New(config.ERR[config.ERR_PROJECT_DEPLOYING]))
	}

	project, err := models.ProjectModel.GetOne(pidint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	repo := git.NewRepository(config.RepoPath, project.Path)
	archiveFile, err := repo.Archive(commit, repo.Path)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	err = c.pushCluster(project, deploy.Id, archiveFile)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	err = models.DeployModel.UpdateStatus(deploy.Id, 1)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, deploy)
}

func (c Deploy) pushCluster(project *models.Project, did int, archiveFile string) error {
	projectClusters, err := models.ProjectClusterModel.GetAll(project.Id)
	if err != nil {
		return errors.New(config.ERR[config.ERR_PROJECT_CLUSTER_EMPTY])
	}

	for _, v1 := range projectClusters {
		for _, v2 := range v1.Cluster.Machines {
			deployHistory, err := models.DeployHistoryModel.Add(did, v2.Ip)
			if err != nil {
				continue
			}

			go c.pushFile(deployHistory.Id, archiveFile, project.PushPath, v1.Bshell, v1.Eshell)
		}
	}

	return nil
}

func (c Deploy) pushFile(dhid int, archiveFile, pushPath, bshell, eshell string) {
	fmt.Println(dhid, archiveFile, pushPath, bshell, eshell)
	status := 1

	status = 2
	models.DeployHistoryModel.Update(dhid, status)
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

	deploys, err := models.DeployModel.GetAll(pidint, 15)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	// get deploy history
	for k, v := range deploys {
		deployHistory, _ := models.DeployHistoryModel.GetAll(v.Id)
		deploys[k].DeployHistory = deployHistory
	}

	RenderRes(r, res, deploys)
}
