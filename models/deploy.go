package models

import "time"

type Deploy struct {
	Id      int
	Pid     int
	Commit  string
	Created time.Time
}

func (m Deploy) TableName() string {
	return "deploy"
}

func (m Deploy) Get(pid int) ([]*Deploy, error) {
	var deploys []*Deploy

	_, err := DB.QueryTable(m.TableName()).Filter("pid", pid).All(&deploys)
	if err != nil {
		return nil, err
	}
	return deploys, nil
}

func (m Deploy) Add(pid int, commit string) (int, error) {
	deploy := &Deploy{
		Pid:     pid,
		Commit:  commit,
		Created: time.Now(),
	}
	id, err := DB.Insert(&deploy)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
