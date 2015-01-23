package models

import "time"

type Project struct {
	Id      int
	Name    string
	Path    string
	Created time.Time
}

func (m Project) GetAll() []Project {

}

func (m Project) GetOne(id int) Project {

}

func (m Project) Add(name, path string) (int, error) {
}
