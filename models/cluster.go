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

func (m Cluster) TableName() string {
	return "cluster"
}

func (m Cluster) Search(name string) ([]*Cluster, error) {
	var clusters []*Cluster

	_, err := DB.QueryTable(m.TableName()).Filter("name__contains", name).All(&clusters)
	if err != nil {
		return nil, err
	}

	return clusters, nil
}

func (m Cluster) Get(id int) (*Cluster, error) {
	var cluster *Cluster

	_, err := DB.QueryTable(m.TableName()).Filter("id", id).All(&cluster)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (m Cluster) Add(name, room, machine, note string) (int, error) {
	cluster := &Cluster{
		Name:    name,
		Room:    room,
		Machine: machine,
		Note:    note,
		Created: time.Now(),
	}
	id, err := DB.Insert(&cluster)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
