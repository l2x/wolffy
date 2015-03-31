package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/l2x/wolffy/server/controllers"
	"github.com/l2x/wolffy/utils"
	"github.com/martini-contrib/render"
)

type Server struct{}

func (s Server) Pull(r render.Render, req *http.Request) {
	res := controllers.NewRes()

	file := []byte{}
	sign := req.URL.Query().Get("sign")
	path := req.URL.Query().Get("path")
	bShell := req.URL.Query().Get("bshell")
	eShell := req.URL.Query().Get("eshell")

	err := checkSign(sign)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}

	err = utils.Mkdir(path)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}

	if bShell != "" {
		err = runCmd(path, bShell)
		if err = controllers.RenderError(r, res, err); err != nil {
			return
		}
	}

	pdir := filepath.Dir(path)
	sfile, err := saveFile(req, pdir)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}
	err = decompress(sfile, path)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}

	if eShell != "" {
		err = utils.RunCmd(path, eShell)
		if err = controllers.RenderError(r, res, err); err != nil {
			return
		}
	}

	err = os.Remove(sfile)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}
	//TODO remove old dir

	res.Code = 0
	controllers.RenderRes(r, res, map[string]string{})
}

func checkSign(sign string) error {
	return nil
}

func decompress(file, path string) error {
	err := utils.Unzip(file)
	if err != nil {
		return err
	}

	ufile := strings.TrimRight(file, ".tar.gz")
	cmd := fmt.Sprintf("ln -s %s %s", ufile, path)
	err := utils.RunCmd(path, cmd)
	if err != nil {
		return err
	}
}

func saveFile(req *http.Request, path string) (string, error) {
	file, handler, err := req.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	save := fmt.Sprintf("%s/%s", strings.TrimRight(path, "/"), handler.Filename)
	f, err := os.OpenFile(save, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)

	return save, nil
}
