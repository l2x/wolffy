package models

import (
	"database/sql"
	"time"
)

var (
	ClusterMachineModel = &ClusterMachine{}
)

type ClusterMachine struct {
	Id      int       `json:"id"`
	Cid     int       `json:"cid"`
	Mid     int       `json:"mid"`
	Created time.Time `json:"created"`
}

func (m ClusterMachine) TableName() string {
	return "cluster_machine"
}

func (m ClusterMachine) TableUnique() [][]string {
	return [][]string{
		[]string{"Cid", "Mid"},
	}
}

func (m ClusterMachine) GetAll(cid int) ([]*Machine, error) {
	clusterMachine := []*ClusterMachine{}
	machines := []*Machine{}
	_, err := DB.QueryTable(m.TableName()).Filter("Cid", cid).All(&clusterMachine)
	if err == sql.ErrNoRows {
		return machines, nil
	}
	if err != nil {
		return machines, err
	}

	machine := &Machine{}
	for _, v := range clusterMachine {
		machine, err = MachineModel.GetOne(v.Mid)
		if err != nil {
			continue
		}
		machines = append(machines, machine)
	}

	return machines, nil
}

func (m ClusterMachine) Add(cid, mid int) error {
	clusterMachine := &ClusterMachine{
		Cid:     cid,
		Mid:     mid,
		Created: time.Now(),
	}
	_, err := DB.Insert(clusterMachine)
	if err != nil {
		return err
	}

	return nil
}

func (m ClusterMachine) Delete(cid int) error {
	_, err := DB.QueryTable(m.TableName()).Filter("Cid", cid).Delete()
	if err != nil {
		return err
	}

	return nil
}
