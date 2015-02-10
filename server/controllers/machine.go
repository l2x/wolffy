package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/martini-contrib/render"

	"github.com/l2x/wolffy/server/models"
)

type Machine struct{}

func (c Machine) Report(r render.Render, req *http.Request) {
	res := NewRes()

	ip := req.URL.Query().Get("ip")
	token := req.URL.Query().Get("token")

	machine, err := models.MachineModel.GetOneByIp(ip)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	if machine.Status == -1 {
		RenderError(r, res, errors.New("server disable."))
		return
	}

	machine, err = models.MachineModel.Update(machine.Id, machine.Ip, machine.Port, machine.Note, token, 1, time.Now())
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, machine)
}

func (c Machine) Add(r render.Render, req *http.Request) {
	res := NewRes()

	ip := req.URL.Query().Get("ip")
	port := req.URL.Query().Get("port")
	note := req.URL.Query().Get("note")

	machine, err := models.MachineModel.Add(ip, port, note)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, machine)
}
