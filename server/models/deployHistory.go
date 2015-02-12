package models

import "time"

var (
	DeployHistoryModel = &DeployHistory{}
)

// status
// - 0 未开始
// - 1 发布中
// - 2 发布完成
// - 3 发布失败
type DeployHistory struct {
	Id       int
	Did      int
	Status   int
	Note     string
	Created  time.Time
	Modified time.Time
}

func (m DeployHistory) TableName() string {
	return "deploy_history"
}

func (m DeployHistory) TableIndex() [][]string {
	return [][]string{
		[]string{"Did"},
	}
}

func (m DeployHistory) GetAll(did int) ([]*DeployHistory, error) {
	var deployHistorys []*DeployHistory

	_, err := DB.QueryTable(m.TableName()).Filter("did", did).All(&deployHistorys)
	if err != nil {
		return nil, err
	}
	return deployHistorys, nil
}

func (m DeployHistory) GetOne(id int) (*DeployHistory, error) {
	deployHistory := &DeployHistory{}
	err := DB.QueryTable(m.TableName()).Filter("Id", id).One(deployHistory)
	if err != nil {
		return nil, err
	}
	return deployHistory, nil
}

func (m DeployHistory) Add(did int) (*DeployHistory, error) {
	deployHistory := &DeployHistory{
		Did:      did,
		Status:   1,
		Created:  time.Now(),
		Modified: time.Now(),
	}
	id, err := DB.Insert(deployHistory)
	if err != nil {
		return nil, err
	}

	deployHistory, err = m.GetOne(int(id))
	if err != nil {
		return nil, err
	}

	return deployHistory, nil
}

func (m DeployHistory) Update(id, status int, note string) error {
	deployHistory := &DeployHistory{
		Id:       id,
		Status:   status,
		Note:     note,
		Modified: time.Now(),
	}
	_, err := DB.Update(deployHistory, "Status", "Note", "Modified")
	if err != nil {
		return err
	}

	return nil
}