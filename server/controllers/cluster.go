package controllers

import (
	"github.com/martini-contrib/render"
)

type Cluster struct{}

func (c Cluster) GetAll(r render.Render) {
	r.JSON(200, map[string]interface{}{"hello": "world"})
}
