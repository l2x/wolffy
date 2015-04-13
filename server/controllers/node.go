package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/martini-contrib/render"

	"github.com/l2x/wolffy/server/models"
)

type Node struct{}

func (c Node) Ping(r render.Render, req *http.Request) {
	res := NewRes()

	ip := req.URL.Query().Get("ip")
	token := req.URL.Query().Get("token")

	node, err := models.NodeModel.GetOneByIp(ip)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	if node.Status == -1 {
		RenderError(r, res, errors.New("server disable."))
		return
	}

	node, err = models.NodeModel.Update(node.Id, node.Ip, node.Port, node.Note, token, 1, time.Now())
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, node)
}

func (c Node) Add(r render.Render, req *http.Request) {
	res := NewRes()

	ip := req.URL.Query().Get("ip")
	port := req.URL.Query().Get("port")
	note := req.URL.Query().Get("note")

	node, err := models.NodeModel.Add(ip, port, note)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, node)
}

func (c Node) GetAll(r render.Render, req *http.Request) {
	res := NewRes()

	node, err := models.NodeModel.GetAll()
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, node)
}

func (c Node) Update(r render.Render, req *http.Request) {
	res := NewRes()

	ip := req.URL.Query().Get("ip")
	note := req.URL.Query().Get("note")
	status := req.URL.Query().Get("status")
	statusInt, _ := strconv.Atoi(status)

	node, err := models.NodeModel.GetOneByIp(ip)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	node, err = models.NodeModel.Update(node.Id, node.Ip, node.Port, note, node.Token, statusInt, time.Now())
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, node)
}

func (c Node) Del(r render.Render, req *http.Request) {
	res := NewRes()

	id := req.URL.Query().Get("id")
	idint, _ := strconv.Atoi(id)
	err := models.NodeModel.Del(idint)

	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, map[string]string{})
}
