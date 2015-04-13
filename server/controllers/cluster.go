package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	nodes := req.URL.Query().Get("nodes")
	note := req.URL.Query().Get("note")
	if id != "" {
		idint, err = strconv.Atoi(id)
		if err = RenderError(r, res, err); err != nil {
			return
		}
	}

	var cluster *models.Cluster
	if idint == 0 {
		cluster, err = models.ClusterModel.Add(name, tags, note)
	} else {
		cluster, err = models.ClusterModel.Update(idint, name, tags, note)
	}

	if err = RenderError(r, res, err); err != nil {
		return
	}
	err = c.Update(cluster.Id, name, tags, note, nodes)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, cluster)
}

func (c Cluster) Update(id int, name, tags, note, nodes string) error {
	err := models.ClusterNodeModel.Delete(id)
	if err != nil {
		return err
	}

	m := strings.Split(nodes, ",")
	for _, v := range m {
		mid, err := strconv.Atoi(strings.Trim(v, " "))
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = models.ClusterNodeModel.Add(id, mid)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
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
