package controllers

import (
	"fmt"
	"net/http"
)

type Site struct{}

func (c Site) Index(res http.ResponseWriter) {
	fmt.Fprint(res, "hello world")
}
