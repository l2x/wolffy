package models

import (
	"database/sql"
	"time"
)

var (
	NodeModel = &Node{}
)

// status
// 0-未启用
// 1-正常
// -1-停用
type Node struct {
	Id         int       `json:"id"`
	Ip         string    `json:"ip"`
	Port       string    `json:"port"`
	Note       string    `json:"note"`
	Status     int       `json:"status"`
	Created    time.Time `json:"created"`
	Modified   time.Time `json:"modified"`
	LastReport time.Time `json:"lastReport"`
}

func (m Node) TableName() string {
	return "node"
}

func (m Node) TableUnique() [][]string {
	return [][]string{
		[]string{"Ip"},
	}
}

func (m Node) GetAll() ([]*Node, error) {
	var nodes []*Node
	_, err := DB.QueryTable(m.TableName()).All(&nodes)
	if err == sql.ErrNoRows {
		return nodes, nil
	}
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func (m Node) GetOne(id int) (*Node, error) {
	node := &Node{
		Id: id,
	}

	if err := DB.Read(node); err != nil {
		return nil, err
	}

	return node, nil
}

func (m Node) GetOneByIp(ip string) (*Node, error) {
	node := &Node{}

	if err := DB.QueryTable(m.TableName()).Filter("Ip", ip).Limit(1).One(node); err != nil {
		return nil, err
	}

	return node, nil
}

func (m Node) Add(ip, port, note string) (*Node, error) {
	node := &Node{
		Ip:         ip,
		Port:       port,
		Note:       note,
		Created:    time.Now(),
		Modified:   time.Now(),
		LastReport: time.Now(),
	}

	id, err := DB.Insert(node)
	if err != nil {
		return nil, err
	}

	node, err = m.GetOne(int(id))
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (m Node) Update(id int, ip, port, note string, status int, lastReport time.Time) (*Node, error) {
	node := &Node{
		Id:         id,
		Ip:         ip,
		Port:       port,
		Note:       note,
		Status:     status,
		Modified:   time.Now(),
		LastReport: lastReport,
	}
	_, err := DB.Update(node, "Ip", "Port", "Note", "Status", "Modified", "LastReport")
	if err != nil {
		return nil, err
	}
	node, err = m.GetOne(id)
	if err != nil {
		return nil, err
	}

	return node, nil
}

// set status
func (m Node) Del(id int) error {
	node := &Node{
		Id: id,
	}
	if _, err := DB.Delete(node); err != nil {
		return err
	}

	return nil
}
