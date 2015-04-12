package models

import (
	"database/sql"
	"time"
)

var (
	ClusterModel = &Cluster{}
)

type Cluster struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Tags     string     `json:"tags"`
	Note     string     `json:"note"`
	Created  time.Time  `json:"created"`
	Modified time.Time  `json:"modified"`
	Machines []*Machine `orm:"-" json:"machines"`
}

func (m Cluster) TableName() string {
	return "cluster"
}

func (m Cluster) TableUnique() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

func (m Cluster) GetAll() ([]*Cluster, error) {
	clusters := []*Cluster{}
	_, err := DB.QueryTable(m.TableName()).All(&clusters)
	if err == sql.ErrNoRows {
		return clusters, nil
	}

	if err != nil {
		return nil, err
	}

	return clusters, nil
}

func (m Cluster) GetOne(id int) (*Cluster, error) {
	cluster := &Cluster{
		Id: id,
	}

	if err := DB.Read(cluster); err != nil {
		return nil, err
	}

	cluster.Machines, _ = ClusterMachineModel.GetAll(cluster.Id)

	return cluster, nil
}

func (m Cluster) Add(name, tags, note string) (*Cluster, error) {
	cluster := &Cluster{
		Name:     name,
		Tags:     tags,
		Note:     note,
		Created:  time.Now(),
		Modified: time.Now(),
	}
	id, err := DB.Insert(cluster)
	if err != nil {
		return nil, err
	}

	cluster, err = m.GetOne(int(id))
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (m Cluster) Update(id int, name, tags, note string) (*Cluster, error) {
	cluster := &Cluster{
		Id:       id,
		Name:     name,
		Tags:     tags,
		Note:     note,
		Modified: time.Now(),
	}
	_, err := DB.Update(cluster, "Name", "Tags", "Note", "Modified")
	if err != nil {
		return nil, err
	}
	cluster, err = m.GetOne(id)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (m Cluster) Delete(id int) error {
	cluster := &Cluster{
		Id: id,
	}
	if _, err := DB.Delete(cluster); err != nil {
		return err
	}

	return nil
}
