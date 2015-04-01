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

	sign := req.URL.Query().Get("sign")
	path := req.URL.Query().Get("path")
	bshell := req.URL.Query().Get("bshell")
	eshell := req.URL.Query().Get("eshell")

	err := checkSign(sign)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}

	fmt.Println(path, bshell, eshell)
	pdir := filepath.Dir(path)
	dir := filepath.Base(path)

	err = utils.Mkdir(pdir)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}

	if bshell != "" {
		err = utils.RunCmd(path, bshell)
		if err = controllers.RenderError(r, res, err); err != nil {
			return
		}
	}

	sfile, err := saveFile(req, pdir)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}
	err = decompress(sfile, pdir, dir)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}

	if eshell != "" {
		err = utils.RunCmd(path, eshell)
		if err = controllers.RenderError(r, res, err); err != nil {
			return
		}
	}

	err = os.Remove(sfile)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}
	//TODO remove old dir

	res.Errno = 0
	controllers.RenderRes(r, res, map[string]string{})
}

func checkSign(sign string) error {
	return nil
}

func decompress(file, pdir, dir string) error {
	fmt.Println("decompress", file, pdir, dir)
	ufile := strings.TrimRight(file, ".tar.gz")
	err := utils.Mkdir(ufile)
	if err != nil {
		return err
	}
	err = utils.UnzipToFolder(pdir, file, ufile)
	if err != nil {
		return err
	}

	cmd := fmt.Sprintf("ln -nsf %s %s", ufile, dir)
	err = utils.RunCmd(pdir, cmd)
	if err != nil {
		return err
	}

	return nil
}

func saveFile(req *http.Request, path string) (string, error) {
	file, handler, err := req.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	filename := strings.Split(handler.Filename, "/")
	save := fmt.Sprintf("%s/%s", strings.TrimRight(path, "/"), filename[len(filename)-1])
	fmt.Println("saveFile", save)

	f, err := os.OpenFile(save, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)

	return save, nil
}
