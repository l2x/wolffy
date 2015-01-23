package models

import "time"

type Cluster struct {
	Id      int
	Name    string
	Room    string
	Machine string
	Note    string
	Created time.Time
}

func (m Cluster) Search(name string) ([]Project, error) {
}

func (m Cluster) Get(id int) (Project, error) {
}

func (m Cluster) Add(name, room, machine, note string) (int, error) {
}
