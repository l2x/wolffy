package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/martini-contrib/render"

	"github.com/l2x/wolffy/server/config"
	"github.com/l2x/wolffy/server/models"
	"github.com/l2x/wolffy/utils"
)

type Node struct{}

func (c Node) Report(r render.Render, req *http.Request) {
	res := NewRes()

	token := req.URL.Query().Get("token")
	sign := req.URL.Query().Get("sign")
	port := req.URL.Query().Get("port")

	err := utils.CheckSign(token, sign, config.PrivateKey)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	ip := utils.ClientIp(req)
	node, err := models.NodeModel.GetOneByIp(ip)
	if err != nil && err != orm.ErrNoRows {
		RenderError(r, res, err)
		return
	}

	if err == orm.ErrNoRows {
		node, err = models.NodeModel.Add(ip, port, "")
	} else {
		node, err = models.NodeModel.Update(node.Id, ip, port, "", 0, time.Now())
	}
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

	node, err = models.NodeModel.Update(node.Id, node.Ip, node.Port, note, statusInt, time.Now())
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, node)
}

func (c Node) Delete(r render.Render, req *http.Request) {
	res := NewRes()

	id := req.URL.Query().Get("id")
	idint, _ := strconv.Atoi(id)
	err := models.NodeModel.Del(idint)

	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, map[string]string{})
}

func (c Node) GetPrivateKey(r render.Render, req *http.Request) {
	res := NewRes()

	RenderRes(r, res, config.PrivateKey)
}
