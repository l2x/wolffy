package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/l2x/wolffy/utils/git"
	"github.com/martini-contrib/render"

	"github.com/l2x/wolffy/master/config"
	"github.com/l2x/wolffy/master/models"
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

	repo := git.NewRepository(config.RepoPath, project.Path)
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
	projectClusters, _ := models.ProjectClusterModel.GetAll(project.Id)

	project.ProjectClusters = projectClusters

	RenderRes(r, res, project)
}

func (c Project) GetAll(r render.Render, req *http.Request) {
	res := NewRes()

	project, err := models.ProjectModel.GetAll()
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, project)
}

func (c Project) Add(r render.Render, req *http.Request) {
	res := NewRes()
	var err error
	var idint int = 0

	//path := "git@123.57.75.209:leiyonglin/wolffy.git"
	id := req.URL.Query().Get("id")
	name := req.URL.Query().Get("name")
	path := req.URL.Query().Get("path")
	pushPath := req.URL.Query().Get("pushPath")
	tags := req.URL.Query().Get("tags")
	note := req.URL.Query().Get("note")
	projectClusters := req.URL.Query().Get("projectClusters")
	if projectClusters == "" {
		projectClusters = "[]"
	}

	if id != "" {
		idint, err = strconv.Atoi(id)
		if err = RenderError(r, res, err); err != nil {
			return
		}
	}

	var clusters []models.ProjectCluster
	err = json.Unmarshal([]byte(projectClusters), &clusters)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	var project *models.Project
	var add bool

	if idint == 0 {
		project, err = models.ProjectModel.Add(name, path, pushPath, tags, note)
		add = true
	} else {
		project, err = models.ProjectModel.Update(idint, name, path, pushPath, tags, note)
	}
	if err = RenderError(r, res, err); err != nil {
		return
	}

	// clone project
	if add {
		repo := git.NewRepository(config.RepoPath, path)
		_, err = repo.Clone()
		if err = RenderError(r, res, err); err != nil {
			models.ProjectModel.Del(project.Id)
			return
		}
	}

	err = models.ProjectClusterModel.DelProject(project.Id)
	for _, v := range clusters {
		_, err = models.ProjectClusterModel.Add(project.Id, v.Cid, "", v.Bshell, v.Eshell, v.Note)
	}

	RenderRes(r, res, project)
}

func (c Project) Delete(r render.Render, req *http.Request) {
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

	err = models.ProjectModel.Del(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	os.Remove(fmt.Sprintf("%s/%s", config.RepoPath, project.Path))

	RenderRes(r, res, map[string]string{})
}

func (c Project) Update(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	name := req.URL.Query().Get("name")
	path := req.URL.Query().Get("path")
	pushpath := req.URL.Query().Get("pushpath")
	tags := req.URL.Query().Get("tags")
	note := req.URL.Query().Get("note")

	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	project, err := models.ProjectModel.Update(idint, name, path, pushpath, tags, note)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, project)
}
