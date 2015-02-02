package models

import "time"

var (
	DeployModel = &Deploy{}
)

type Deploy struct {
	Id      int
	Pid     int
	Commit  string
	Created time.Time
}

func (m Deploy) TableName() string {
	return "deploy"
}

func (m Deploy) TableIndex() [][]string {
	return [][]string{
		[]string{"Pid"},
	}
}

func (m Deploy) GetAll(pid int) ([]*Deploy, error) {
	var deploys []*Deploy

	_, err := DB.QueryTable(m.TableName()).Filter("pid", pid).All(&deploys)
	if err != nil {
		return nil, err
	}
	return deploys, nil
}

func (m Deploy) GetOne(id int) (*Deploy, error) {
	deploy := &Deploy{}
	err := DB.QueryTable(m.TableName()).Filter("Id", id).One(&deploy)
	if err != nil {
		return nil, err
	}
	return deploy, nil
}

func (m Deploy) Add(pid int, commit string) (*Deploy, error) {
	deploy := &Deploy{
		Pid:     pid,
		Commit:  commit,
		Created: time.Now(),
	}
	id, err := DB.Insert(deploy)
	if err != nil {
		return nil, err
	}

	deploy, err = m.GetOne(int(id))
	if err != nil {
		return nil, err
	}

	return deploy, nil
}
