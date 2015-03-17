package controllers

import "time"

var (
	Sessions map[string]*Session
)

type Session struct {
	id          int
	LastLogin   time.Time
	Expire      time.Time
	LastLoginIp string
}

func NewSession() {
	Sessions = map[string]*Session{}
	go trash()
}

func trash() {
	for {
		for k, v := range Sessions {
			if time.Now().Before(v.Expire) {
				delete(Sessions, k)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
