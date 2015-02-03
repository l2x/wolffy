package models

import "time"

var (
	ProjectModel = &Project{}
)

type Project struct {
	Id       int
	Name     string
	Path     string
	PushPath string
	Tags     string
	Note     string
	Created  time.Time
}

func (m Project) TableName() string {
	return "project"
}

func (m Project) TableUnique() [][]string {
	return [][]string{
		[]string{"Pid", "Name"},
	}
}

func (m Project) GetAll() ([]*Project, error) {
	projects := []*Project{}
	if _, err := DB.QueryTable(m.TableName()).All(&projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (m Project) GetOne(id int) (*Project, error) {
	project := &Project{
		Id: id,
	}
	if err := DB.QueryTable(m.TableName()).Filter("Id", id).One(project); err != nil {
		return nil, err
	}

	return project, nil
}

func (m Project) Add(name, path, pushpath, tags, note string) (*Project, error) {
	project := &Project{
		Name:     name,
		Path:     path,
		PushPath: pushpath,
		Tags:     tags,
		Note:     note,
		Created:  time.Now(),
	}
	id, err := DB.Insert(project)
	if err != nil {
		return nil, err
	}

	project, err = m.GetOne(int(id))
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (m Project) Del(id int) error {
	project := &Project{
		Id: id,
	}
	if _, err := DB.Delete(project); err != nil {
		return err
	}

	return nil
}

func (m Project) Update(id int, name, path, pushpath, tags, note string) (*Project, error) {
	project := &Project{
		Id:       id,
		Name:     name,
		Path:     path,
		PushPath: pushpath,
		Tags:     tags,
		Note:     note,
		Created:  time.Now(),
	}
	if _, err := DB.Update(project); err != nil {
		return nil, err
	}

	return project, nil
}
