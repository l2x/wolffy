package main

import (
	"net/http"

	"github.com/martini-contrib/render"
)

type Server struct{}

func (s Server) Pull(r render.Render, req *http.Request) {
}
