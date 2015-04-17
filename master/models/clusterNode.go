package models

import (
	"database/sql"
	"time"
)

var (
	ClusterNodeModel = &ClusterNode{}
)

type ClusterNode struct {
	Id      int       `json:"id"`
	Cid     int       `json:"cid"`
	Mid     int       `json:"mid"`
	Created time.Time `json:"created"`
}

func (m ClusterNode) TableName() string {
	return "cluster_node"
}

func (m ClusterNode) TableUnique() [][]string {
	return [][]string{
		[]string{"Cid", "Mid"},
	}
}

func (m ClusterNode) GetAll(cid int) ([]*Node, error) {
	clusterNode := []*ClusterNode{}
	nodes := []*Node{}
	_, err := DB.QueryTable(m.TableName()).Filter("Cid", cid).All(&clusterNode)
	if err == sql.ErrNoRows {
		return nodes, nil
	}
	if err != nil {
		return nodes, err
	}

	node := &Node{}
	for _, v := range clusterNode {
		node, err = NodeModel.GetOne(v.Mid)
		if err != nil {
			continue
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (m ClusterNode) Add(cid, mid int) error {
	clusterNode := &ClusterNode{
		Cid:     cid,
		Mid:     mid,
		Created: time.Now(),
	}
	_, err := DB.Insert(clusterNode)
	if err != nil {
		return err
	}

	return nil
}

func (m ClusterNode) Delete(cid int) error {
	_, err := DB.QueryTable(m.TableName()).Filter("Cid", cid).Delete()
	if err != nil {
		return err
	}

	return nil
}
