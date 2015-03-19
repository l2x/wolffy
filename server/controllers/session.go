package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/l2x/wolffy/server/config"
	"github.com/l2x/wolffy/server/models"
	"github.com/l2x/wolffy/utils"
)

var (
	Sessions *Session
)

type Session struct {
	cache    map[string]*Cache
	interval int
}

type Cache struct {
	Id        int
	LastLogin time.Time
	Expire    time.Time
	Ip        string
}

func NewSession() {
	Sessions = &Session{
		cache:    map[string]*Cache{},
		interval: config.SessionInterval,
	}
	go Sessions.trash()
}

func CheckSession(res http.ResponseWriter, req *http.Request) error {
	cookie, err := req.Cookie(config.CookieName)
	if err != nil {
		return err
	}
	sid := cookie.Value
	c, ok := Sessions.cache[sid]
	if !ok {
		return errors.New("session not found")
	}

	if time.Now().Before(c.Expire) {
		delete(Sessions.cache, sid)
		return errors.New("seesion expired")
	}

	cookie.Expires = time.Now().Add(time.Duration(config.SessionExpire) * time.Second)
	http.SetCookie(res, cookie)

	return nil
}

func (s *Session) Add(id int, username, ip string) {
	cache := &Cache{
		Id:        id,
		Ip:        ip,
		LastLogin: time.Now(),
		Expire:    time.Now().Add(time.Duration(config.SessionExpire) * time.Second),
	}

	sid := s.genSid()
	s.cache[sid] = cache
}

func (s *Session) GetUser(req *http.Request) (*models.User, error) {
	cookie, err := req.Cookie(config.CookieName)
	if err != nil {
		return nil, err
	}
	sid := cookie.Value

	session, err := s.Get(sid)
	if err != nil {
		return nil, err
	}

	user, err := models.UserModel.GetOne(session.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Session) Update(sid string) bool {
	cache, ok := s.cache[sid]
	if !ok {
		return false
	}
	cache.Expire = time.Now().Add(time.Duration(config.SessionExpire) * time.Second)

	return true
}

func (s *Session) Get(sid string) (*Cache, error) {
	session, ok := s.cache[sid]
	if !ok {
		return nil, errors.New("session not found")
	}
	return session, nil
}

func (s *Session) Del(sid string) {
	delete(s.cache, sid)
}

func (s *Session) genSid() string {
	t := fmt.Sprintf("%v%v", time.Now().UnixNano(), utils.RandInt(0, 999999))
	sid := utils.Md5(utils.Md5(t))
	return sid
}

func (s *Session) trash() {
	for {
		for k, v := range s.cache {
			if time.Now().Before(v.Expire) {
				delete(s.cache, k)
			}
		}

		time.Sleep(time.Duration(s.interval) * time.Second)
	}
}
