package controllers

import (
	"fmt"
	"time"

	"github.com/l2x/wolffy/server/config"
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

func (s *Session) add(id int, username, ip string) {
	cache := &Cache{
		Id:        id,
		Ip:        ip,
		LastLogin: time.Now(),
		Expire:    time.Now().Add(time.Duration(config.SessionExpire) * time.Second),
	}

	sid := s.genSid()
	s.cache[sid] = cache
}

func (s *Session) Update(sid string) bool {
	cache, ok := s.cache[sid]
	if !ok {
		return false
	}
	cache.Expire = time.Now().Add(time.Duration(config.SessionExpire) * time.Second)

	return true
}

func (s *Session) Del(sid string) {
	delete(s.cache, sid)
}

func (s *Session) genSid() string {
	t := fmt.Sprintf("%v", time.Now().UnixNano())
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
