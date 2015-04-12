package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/martini-contrib/render"

	"github.com/l2x/wolffy/server/config"
	"github.com/l2x/wolffy/server/models"
	"github.com/l2x/wolffy/utils"
)

type User struct{}

func (c User) Login(r render.Render, w http.ResponseWriter, req *http.Request) {
	res := NewRes()

	username := strings.Trim(req.URL.Query().Get("username"), " ")
	password := strings.Trim(req.URL.Query().Get("password"), " ")

	user, err := models.UserModel.GetViaUsername(username)
	if err == sql.ErrNoRows {
		res.Errno = config.ERR_USER_NOT_FOUND
		err = config.GetErr(res.Errno)
	}
	if err = RenderError(r, res, err); err != nil {
		return
	}

	password = utils.SignPassword(password, user.Id)
	user, err = models.UserModel.CheckPassword(username, password)
	if err == sql.ErrNoRows {
		res.Errno = config.ERR_USER_PASSWORD_INCORRECT
		err = config.GetErr(res.Errno)
	}
	if err = RenderError(r, res, err); err != nil {
		return
	}

	ip := utils.ClientIp(req)
	Sessions.Add(w, user.Id, user.Username, ip)

	// 如果长时间没有登录，需要修改密码
	if user.LastLogin.Before(time.Now().AddDate(0, -6, 0)) {
		res.Errno = config.ERR_USER_NEED_CHANGE_PWD
		err = config.GetErr(res.Errno)
		RenderError(r, res, err)
		return
	}

	err = models.UserModel.UpdateLastLogin(user.Id, ip)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, user)
}

func (c User) Logout(r render.Render, req *http.Request) {
	res := NewRes()

	cookie, err := req.Cookie(config.CookieName)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	sid := cookie.Value
	Sessions.Del(sid)

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

	user, err := models.UserModel.GetAll()
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, user)
}

func (c User) Edit(r render.Render, req *http.Request) {
	res := NewRes()

	username := req.URL.Query().Get("username")
	name := req.URL.Query().Get("name")
	password := req.URL.Query().Get("password")
	administrator := req.URL.Query().Get("administrator")
	administratorint := 0
	if administrator == "true" {
		administratorint = 1
	}
	err := checkAdministrator(req)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	user, err := models.UserModel.Add(username, name, "", administratorint)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	signPassword := utils.SignPassword(password, user.Id)
	err = models.UserModel.UpdatePassword(user.Id, signPassword)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	//user.Password = password

	RenderRes(r, res, user)
}

func (c User) Del(r render.Render, req *http.Request) {
	res := NewRes()
	id := req.URL.Query().Get("id")
	idint, err := strconv.Atoi(id)
	if err = RenderError(r, res, err); err != nil {
		return
	}
	err = checkAdministrator(req)
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

	err := checkAdministrator(req)
	if err = RenderError(r, res, err); err != nil {
		return
	}

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

	oldpassword := req.URL.Query().Get("newpassword")
	newpassword := req.URL.Query().Get("oldpassword")

	user, err := Sessions.GetUser(req)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	oldpassword = utils.SignPassword(oldpassword, user.Id)
	user, err = models.UserModel.CheckPassword(user.Username, oldpassword)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	newpassword = utils.SignPassword(newpassword, user.Id)
	err = models.UserModel.UpdatePassword(user.Id, newpassword)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	err = models.UserModel.UpdateLastLogin(user.Id, utils.ClientIp(req))
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, user)
}

func (c User) GetUserInfo(r render.Render, req *http.Request) {
	res := NewRes()

	user, err := Sessions.GetUser(req)
	if err = RenderError(r, res, err); err != nil {
		return
	}

	RenderRes(r, res, user)
}

func checkAdministrator(req *http.Request) error {
	user, err := Sessions.GetUser(req)
	if err != nil {
		return err
	}

	if user.Administrator != 1 {
		return config.GetErr(config.ERR_USER_NOT_ADMIN)
	}

	return nil
}
