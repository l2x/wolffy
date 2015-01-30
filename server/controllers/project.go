package controllers

import (
	"github.com/l2x/wolffy/utils/git"
	"github.com/martini-contrib/render"
)

type Project struct{}

func (c Project) Add(r render.Render) {
	res := NewRes()

	path := "/tmp/repo"
	remotePath := "git@123.57.75.209:leiyonglin/wolffy.git"

	repo := git.NewRepository(path, remotePath)
	_, err := repo.Clone()
	if err != nil {
		res.Errmsg = err.Error()
	}

	r.JSON(200, res)
}
