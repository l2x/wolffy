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
	Id      int       `json:"id"`
	Did     int       `json:"did"`
	Ip      string    `json:"ip"`
	Status  int       `json:"status"`
	Note    string    `json:"note"`
	Created time.Time `json:"created"`
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
	deployHistorys := []*DeployHistory{}

	_, err := DB.QueryTable(m.TableName()).Filter("did", did).All(&deployHistorys)
	if err != nil {
		return deployHistorys, err
	}
	return deployHistorys, nil
}

func (m DeployHistory) GetOne(id int) (*DeployHistory, error) {
	deployHistory := &DeployHistory{}
	err := DB.QueryTable(m.TableName()).Filter("Id", id).Limit(1).One(deployHistory)
	if err != nil {
		return nil, err
	}
	return deployHistory, nil
}

func (m DeployHistory) Add(did int, ip string) (*DeployHistory, error) {
	deployHistory := &DeployHistory{
		Did:     did,
		Status:  1,
		Ip:      ip,
		Created: time.Now(),
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

func (m DeployHistory) Update(id, status int) error {
	deployHistory := &DeployHistory{
		Id:     id,
		Status: status,
	}
	_, err := DB.Update(deployHistory, "Status")
	if err != nil {
		return err
	}

	return nil
}
