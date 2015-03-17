package controllers

import (
	"time"

	"github.com/l2x/wolffy/server/config"
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

	s.cache[username] = cache
}

func (s *Session) Update(username string) bool {
	cache, ok := s.cache[username]
	if !ok {
		return false
	}
	cache.Expire = time.Now().Add(time.Duration(config.SessionExpire) * time.Second)

	return true
}

func (s *Session) del(username string) {
	delete(s.cache, username)
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
