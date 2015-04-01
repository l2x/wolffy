package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
)

func router() {
	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(render.Renderer())

	server := Server{}
	m.Post("/pull/", server.Pull)

	m.RunOnAddr(":8001")
}
