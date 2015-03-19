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

func (c Project) Diff(r render.Render, req *http.Request) {
	res := NewRes()

	id := req.URL.Query().Get("id")
	commita := req.URL.Query().Get("commita")
	commitb := req.URL.Query().Get("commitb")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	project, err := models.ProjectModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	repo := git.NewRepository(config.RepoPath, project.Path)
	diff, err := repo.Diff(commita, commitb)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, diff)
}

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

func (c Project) Search(r render.Render, req *http.Request) {
	res := NewRes()
	key := req.URL.Query().Get("key")
	projects, err := models.ProjectModel.Search(key)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, projects)
}

func (c Project) Add(r render.Render, req *http.Request) {
	res := NewRes()

	//path := "git@123.57.75.209:leiyonglin/wolffy.git"
	name := req.URL.Query().Get("name")
	path := req.URL.Query().Get("path")
	pushpath := req.URL.Query().Get("pushpath")
	tags := req.URL.Query().Get("tags")
	note := req.URL.Query().Get("note")

	project, err := models.ProjectModel.Add(name, path, pushpath, tags, note)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	repo := git.NewRepository(config.RepoPath, project.Path)
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
