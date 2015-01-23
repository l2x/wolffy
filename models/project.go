package models

import "time"

type Project struct {
	Id      int
	Pid     int
	Name    string
	Path    string
	Note    string
	Created time.Time
}

func (m Project) GetAll() ([]Project, error) {

}

func (m Project) GetOne(id int) (Project, error) {

}

func (m Project) Add(name, path string) (int, error) {
}
