package models

import "time"

var (
	ClusterModel = &Cluster{}
)

type Cluster struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Tags     string    `json:"tags"`
	Machines string    `json:"machines"`
	Note     string    `json:"note"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func (m Cluster) TableName() string {
	return "cluster"
}

func (m Cluster) TableUnique() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

func (m Cluster) Search(name string) ([]*Cluster, error) {
	var clusters []*Cluster

	_, err := DB.QueryTable(m.TableName()).Filter("Name__icontains", name).All(&clusters)
	if err != nil {
		return nil, err
	}

	return clusters, nil
}

func (m Cluster) GetAll() ([]*Cluster, error) {
	clusters := []*Cluster{}
	if _, err := DB.QueryTable(m.TableName()).All(&clusters); err != nil {
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

	return cluster, nil
}

func (m Cluster) Add(name, tags, machines, note string) (*Cluster, error) {
	cluster := &Cluster{
		Name:     name,
		Tags:     tags,
		Machines: machines,
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

func (m Cluster) Update(id int, name, tags, machines, note string) (*Cluster, error) {
	cluster := &Cluster{
		Id:       id,
		Name:     name,
		Tags:     tags,
		Machines: machines,
		Note:     note,
		Modified: time.Now(),
	}
	_, err := DB.Update(cluster, "Name", "Tags", "Machines", "Note", "Modified")
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (m Cluster) Del(id int) error {
	cluster := &Cluster{
		Id: id,
	}
	if _, err := DB.Delete(cluster); err != nil {
		return err
	}

	return nil
}
