package controllers

import (
	"net/http"
	"strconv"

	"github.com/martini-contrib/render"

	"github.com/l2x/wolffy/server/models"
)

type Cluster struct{}

func (c Cluster) GetAll(r render.Render, req *http.Request) {
	res := NewRes()

	clusters, err := models.ClusterModel.GetAll()
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, clusters)
}

func (c Cluster) Get(r render.Render, req *http.Request) {
	res := NewRes()

	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	cluster, err := models.ClusterModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, cluster)
}

func (c Cluster) Add(r render.Render, req *http.Request) {
	res := NewRes()

	name := req.URL.Query().Get("name")
	tags := req.URL.Query().Get("tags")
	machine := req.URL.Query().Get("machine")
	note := req.URL.Query().Get("note")

	cluster, err := models.ClusterModel.Add(name, tags, machine, note)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, cluster)
}

func (c Cluster) Del(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	err = models.ClusterModel.Del(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, map[string]string{})
}

func (c Cluster) Update(r render.Render, req *http.Request) {
	res := NewRes()
	name := req.URL.Query().Get("name")
	tags := req.URL.Query().Get("tags")
	machine := req.URL.Query().Get("machine")
	note := req.URL.Query().Get("note")
	id := req.URL.Query().Get("id")

	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	cluster, err := models.ClusterModel.Update(idint, name, tags, machine, note)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, cluster)
}
