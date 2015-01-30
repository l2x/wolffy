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

func (m Project) TableName() string {
	return "project"
}

func (m Project) GetAll() ([]*Project, error) {
	var project []*Project
	_, err := DB.QueryTable(m.TableName()).All(&project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (m Project) GetOne(id int) (*Project, error) {
	var project *Project = &Project{
		Id: id,
	}
	err := DB.Read(&project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (m Project) Add(pid int, name, path, note string) (int, error) {
	var project *Project = &Project{
		Pid:     pid,
		Name:    name,
		Path:    path,
		Note:    note,
		Created: time.Now(),
	}
	id, err := DB.Insert(project)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
