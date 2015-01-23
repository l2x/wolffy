package models

import "time"

type Deploy struct {
	Id      int
	Pid     int
	Commit  string
	Created time.Time
}

func (m Deploy) Get(pid int) ([]Deploy, error) {
}

func (m Deploy) Add(pid int, commit string) (int, error) {
}
