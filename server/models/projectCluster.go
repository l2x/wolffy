package models

import (
	"database/sql"
	"time"
)

var (
	ProjectClusterModel = &ProjectCluster{}
)

type ProjectCluster struct {
	Id            int       `json:"id"`
	Pid           int       `json:"pid"`
	Cid           int       `json:"cid"`
	CustomMachine string    `json:"customMachine"`
	Bshell        string    `json:"bshell"`
	Eshell        string    `json:"eshell"`
	Note          string    `json:"note"`
	Created       time.Time `json:"created"`
	Modified      time.Time `json:"modified"`
	Cluster       *Cluster  `orm:"-" json:"cluster"`
}

func (m ProjectCluster) TableName() string {
	return "project_cluster"
}

func (m ProjectCluster) TableUnique() [][]string {
	return [][]string{
		[]string{"Pid", "Cid"},
	}
}

func (m ProjectCluster) GetAll(pid int) ([]*ProjectCluster, error) {
	var projectClusters []*ProjectCluster

	_, err := DB.QueryTable(m.TableName()).Filter("pid", pid).All(&projectClusters)
	if err == sql.ErrNoRows {
		return projectClusters, nil
	}

	if err != nil {
		return nil, err
	}

	for k, v := range projectClusters {
		cluster, err := ClusterModel.GetOne(v.Cid)
		if err != nil {
			continue
		}
		projectClusters[k].Cluster = cluster
	}

	return projectClusters, nil
}

func (m ProjectCluster) Del(id int) error {
	projectCluster := &ProjectCluster{
		Id: id,
	}
	if _, err := DB.Delete(projectCluster); err != nil {
		return err
	}

	return nil
}

func (m ProjectCluster) DelProject(pid int) error {
	_, err := DB.QueryTable(m.TableName()).Filter("pid", pid).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (m ProjectCluster) Add(pid, cid int, customMachine, bshell, eshell, note string) (*ProjectCluster, error) {
	projectCluster := &ProjectCluster{
		Pid:           pid,
		Cid:           cid,
		CustomMachine: customMachine,
		Bshell:        bshell,
		Eshell:        eshell,
		Note:          note,
		Created:       time.Now(),
		Modified:      time.Now(),
	}

	_, err := DB.Insert(projectCluster)
	if err != nil {
		return nil, err
	}

	return projectCluster, nil
}

func (m ProjectCluster) Update(id, pid, cid int, customMachine, bshell, eshell, note string) (*ProjectCluster, error) {
	projectCluster := &ProjectCluster{
		Id:            id,
		Pid:           pid,
		Cid:           cid,
		CustomMachine: customMachine,
		Bshell:        bshell,
		Eshell:        eshell,
		Note:          note,
		Modified:      time.Now(),
	}

	_, err := DB.Update(projectCluster, "Pid", "Cid", "CustomMachine", "Bshell", "Eshell", "Note", "Modified")
	if err != nil {
		return nil, err
	}

	return projectCluster, nil
}
