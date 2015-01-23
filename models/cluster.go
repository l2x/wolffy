package models

import (
	"time"

	"github.com/l2x/wolffy/controllers"
)

type Cluster struct {
	Id      int
	Name    string
	Room    string
	Machine string
	Note    string
	Created time.Time
}

func (m Cluster) Search(name string) ([]*Cluster, error) {
	sqlStr := `SELECT id, name, room, machine, note, created 
			   FROM cluster 
			   WHERE name = ?`
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		return nil, err
	}

	clusters := []*Cluster{}
	for rows.Next() {
		c := &Cluster{}
		err = rows.Scan(&c.Id, &c.Name, &c.Room, &c.Machine, &c.Note, &c.Created)
		if err != nil {
			continue
		}
		clusters = append(clusters, c)
	}

	return clusters, nil
}

func (m Cluster) Get(id int) (*Cluster, error) {
	sqlStr := `SELECT id, name, room, machine, note, created 
			   FROM cluster 
			   WHERE id = ?
			   LIMIT 1`
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	c := &Cluster{}
	err = stmt.QueryRow(id).Scan(&c.Id, &c.Name, &c.Room, &c.Machine, &c.Note, &c.Created)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (m Cluster) Add(name, room, machine, note string) (int, error) {
	sqlStr := `INSERT INTO cluster(name, room, machine, note, created)
			   VALUES(?, ?, ?, ?, ?)`
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	created := time.Now().Format(controllers.DatetimeFormat)
	res, err := stmt.Exec(name, room, machine, note, created)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return int(id), nil
}
