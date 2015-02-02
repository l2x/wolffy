package controllers

import (
	"net/http"
	"strconv"

	"github.com/martini-contrib/render"

	"github.com/l2x/wolffy/server/models"
)

type Product struct{}

func (c Product) Get(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	product, err := models.ProductModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, product)
}

func (c Product) GetAll(r render.Render, req *http.Request) {
	res := NewRes()

	product, err := models.ProductModel.GetAll()
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, product)
}

func (c Product) Add(r render.Render, req *http.Request) {
	res := NewRes()

	name := req.URL.Query().Get("name")
	note := req.URL.Query().Get("note")

	product, err := models.ProductModel.Add(name, note)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, product)
}

func (c Product) Update(r render.Render, req *http.Request) {
	res := NewRes()

	name := req.URL.Query().Get("name")
	note := req.URL.Query().Get("note")
	id := req.URL.Query().Get("id")

	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	err = models.ProductModel.Update(idint, name, note)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, map[string]string{})
}

func (c Product) Del(r render.Render, req *http.Request) {
	res := NewRes()

	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	err = models.ProductModel.Del(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, map[string]string{})
}
