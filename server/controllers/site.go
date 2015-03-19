package controllers

import "github.com/martini-contrib/render"

type Site struct{}

func (c Site) Index(r render.Render) {
	res := NewRes()
	RenderRes(r, res, map[string]string{})
}
