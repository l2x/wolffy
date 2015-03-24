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
	var idint int
	var err error

	id := req.URL.Query().Get("id")
	name := req.URL.Query().Get("name")
	tags := req.URL.Query().Get("tags")
	machines := req.URL.Query().Get("machines")
	note := req.URL.Query().Get("note")
	if id != "" {
		idint, err = strconv.Atoi(id)
		if err = RenderError(r, res, err); err != nil {
			return
		}
	}

	var cluster *models.Cluster
	if idint == 0 {
		cluster, err = models.ClusterModel.Add(name, tags, machines, note)
	} else {
		cluster, err = models.ClusterModel.Update(idint, name, tags, machines, note)
	}

	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, cluster)
}

func (c Cluster) Delete(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	err = models.ClusterModel.Delete(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, map[string]string{})
}

func (c Cluster) Update(r render.Render, req *http.Request) {
	res := NewRes()
	name := req.URL.Query().Get("name")
	tags := req.URL.Query().Get("tags")
	machines := req.URL.Query().Get("machines")
	note := req.URL.Query().Get("note")
	id := req.URL.Query().Get("id")

	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	cluster, err := models.ClusterModel.Update(idint, name, tags, machines, note)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, cluster)
}
