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

	token := req.URL.Query().Get("token")
	sign := req.URL.Query().Get("sign")
	path := req.URL.Query().Get("path")
	bshell := req.URL.Query().Get("bshell")
	eshell := req.URL.Query().Get("eshell")

	err := utils.CheckSign(token, sign, PrivateKey)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}

	pdir := filepath.Dir(path)
	dir := filepath.Base(path)

	err = utils.Mkdir(pdir)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}

	if bshell != "" {
		err = execCmd(path, bshell)
		if err = controllers.RenderError(r, res, err); err != nil {
			return
		}
	}

	sfile, err := saveFile(req, pdir)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}
	defer os.Remove(sfile)
	err = decompress(sfile, pdir, dir)
	if err = controllers.RenderError(r, res, err); err != nil {
		return
	}

	if eshell != "" {
		err = execCmd(path, eshell)
		if err = controllers.RenderError(r, res, err); err != nil {
			return
		}
	}

	res.Errno = 0
	controllers.RenderRes(r, res, map[string]string{})
}

func execCmd(path, c string) error {
	arr := strings.Split(c, "\n")
	for _, v := range arr {
		if strings.Trim(v, " ") == "" {
			continue
		}
		err := utils.RunCmd(path, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func decompress(file, pdir, dir string) error {
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

	//删除上个版本目录
	buf, err := readLog(pdir)
	if err != nil {
	}
	err = os.Remove(string(buf))
	if err != nil {
	}
	//记录这个个版本目录
	err = addLog(ufile, pdir)
	if err != nil {
	}

	return nil
}

func readLog(pdir string) ([]byte, error) {
	f, err := os.OpenFile(fmt.Sprintf("%s.log", pdir), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 1024)
	_, err = f.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func addLog(ufile, pdir string) error {
	f, err := os.OpenFile(fmt.Sprintf("%s.log", pdir), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(ufile)
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

	f, err := os.OpenFile(save, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)

	return save, nil
}
