package controllers

import (
	"net/http"
	"strconv"

	"github.com/martini-contrib/render"

	"github.com/l2x/wolffy/server/models"
)

type User struct{}

func (c User) Login(r render.Render, req *http.Request) {
	res := NewRes()

	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")
	password = SignPassword(password)

	user, err := models.UserModel.CheckPassword(username, password)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, user)
}

func (c User) Logout(r render.Render, req *http.Request) {
	res := NewRes()

	RenderRes(r, res, map[string]string{})
}

func (c User) Get(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}
	user, err := models.UserModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}
	user.Password = ""

	RenderRes(r, res, user)
}

func (c User) GetAll(r render.Render, req *http.Request) {
	res := NewRes()

	//TODO check administrator

	user, err := models.UserModel.GetAll()
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, user)
}

func (c User) Add(r render.Render, req *http.Request) {
	res := NewRes()

	username := req.URL.Query().Get("username")
	name := req.URL.Query().Get("name")
	administrator := req.URL.Query().Get("administrator")
	administratorint, _ := strconv.Atoi(administrator)

	password := GenPassword()

	user, err := models.UserModel.Add(username, name, password, administratorint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, user)
}

func (c User) Del(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	err = models.UserModel.Del(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, map[string]string{})
}

func (c User) Update(r render.Render, req *http.Request) {
	res := NewRes()

	id := req.URL.Query().Get("id")
	username := req.URL.Query().Get("username")
	name := req.URL.Query().Get("name")
	administrator := req.URL.Query().Get("administrator")
	administratorint, _ := strconv.Atoi(administrator)

	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	user, err := models.UserModel.Update(idint, username, name, administratorint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, user)
}

func (c User) UpdatePassword(r render.Render, req *http.Request) {
	res := NewRes()

	id := req.URL.Query().Get("id")
	oldpassword := req.URL.Query().Get("newpassword")
	newpassword := req.URL.Query().Get("oldpassword")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	user, err := models.UserModel.GetOne(idint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	user, err = models.UserModel.CheckPassword(user.Username, oldpassword)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	newpassword = SignPassword(newpassword)
	user, err = models.UserModel.UpdatePassword(idint, newpassword)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, user)
}

func SignPassword(password string) string {
	return password
}

func GenPassword() string {
	return "123456"
}
