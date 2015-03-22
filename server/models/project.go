package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	ProjectModel = &Project{}
)

type Project struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	PushPath string    `json:"pushPath"`
	Tags     string    `json:"tags"`
	Note     string    `json:"note"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`

	ProjectClusters []*ProjectCluster `orm:"-", json:"projectClusters"`
}

func (m Project) TableName() string {
	return "project"
}

func (m Project) TableUnique() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

func (m Project) Search(key string) ([]*Project, error) {
	projects := []*Project{}

	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", key).Or("Tags__icontains", key)

	_, err := DB.QueryTable(m.TableName()).SetCond(cond1).All(&projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
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
	if err := DB.QueryTable(m.TableName()).Filter("Id", id).Limit(1).One(project); err != nil {
		return nil, err
	}

	project.ProjectClusters, _ = ProjectClusterModel.GetAll(project.Id)

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
		Modified: time.Now(),
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
		Modified: time.Now(),
	}
	if _, err := DB.Update(project, "Name", "Path", "PushPath", "Tags", "Note", "Modified"); err != nil {
		return nil, err
	}

	project, err := m.GetOne(project.Id)
	if err != nil {
		return nil, err
	}

	return project, nil
}
