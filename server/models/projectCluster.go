package models

import "time"

var (
	ProjectClusterModel = &ProjectCluster{}
)

type ProjectCluster struct {
	Id            int
	Pid           int
	Cid           int
	CustomMachine string
	Bshell        string
	Eshell        string
	Note          string
	Created       time.Time
}

func (m ProjectCluster) TableName() string {
	return "project_cluster"
}

func (m ProjectCluster) TableIndex() [][]string {
	return [][]string{
		[]string{"Pid"},
	}
}

func (m ProjectCluster) GetAll(pid int) ([]*ProjectCluster, error) {
	var projectClusters []*ProjectCluster

	_, err := DB.QueryTable(m.TableName()).Filter("pid", pid).All(&projectClusters)
	if err != nil {
		return nil, err
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

func (m ProjectCluster) Update(id, pid, cid int, customMachine, bshell, eshell, note string) (*ProjectCluster, error) {
	projectCluster := &ProjectCluster{
		Id:            id,
		Pid:           pid,
		Cid:           cid,
		CustomMachine: customMachine,
		Bshell:        bshell,
		Eshell:        eshell,
		Note:          note,
		Created:       time.Now(),
	}

	_, err := DB.Update(projectCluster)
	if err != nil {
		return nil, err
	}

	return projectCluster, nil
}
