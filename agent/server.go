package main

import (
	"errors"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/l2x/wolffy/server/controllers"
	"github.com/martini-contrib/render"
)

type Server struct{}

func (s Server) Pull(r render.Render, req *http.Request) {
	res := controllers.NewRes()

	file := []byte{}
	sign := req.URL.Query().Get("sign")
	path := req.URL.Query().Get("path")
	bShell := req.URL.Query().Get("bShell")
	eShell := req.URL.Query().Get("eShell")

	err := checkSign(sign)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}

	if bShell != "" {
		err = runCmd(path, bShell)
		if err = controllers.RenderError(r, res, err); err != nil {
			return
		}
	}

	err = saveFile(path, file)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}

	if eShell != "" {
		err = runCmd(path, eShell)
		if err = controllers.RenderError(r, res, err); err != nil {
			return
		}
	}

	controllers.RenderRes(r, res, map[string]string{})
}

func checkSign(sign string) error {
	return nil
}

func unzip(path, file string) error {
	_, stderr, err := com.ExecCmdDir(path, "tar", "xvf", file)
	if err != nil {
		return errors.New(err.Error() + "\n" + stderr)
	}
	return nil
}

func runCmd(path, cmd string) error {
	_, stderr, err := com.ExecCmdDir(path, "bash", "-c", cmd)
	if err != nil {
		return errors.New(err.Error() + "\n" + stderr)
	}
	return nil
}

func saveFile(path string, file []byte) error {
	return nil
}
