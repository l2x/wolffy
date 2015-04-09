package models

import "time"

var (
	MachineModel = &Machine{}
)

// status
// 0-未启用
// 1-正常
// -1-停用
type Machine struct {
	Id         int       `json:"id"`
	Ip         string    `json:"ip"`
	Port       string    `json:"port"`
	Note       string    `json:"note"`
	Token      string    `json:"-"`
	Status     int       `json:"status"`
	Created    time.Time `json:"created"`
	Modified   time.Time `json:"modified"`
	LastReport time.Time `json:"lastReport"`
}

func (m Machine) TableName() string {
	return "machine"
}

func (m Machine) TableUnique() [][]string {
	return [][]string{
		[]string{"Ip"},
	}
}

func (m Machine) Search(ip string) ([]*Machine, error) {
	var machines []*Machine

	_, err := DB.QueryTable(m.TableName()).Filter("Ip__contains", ip).All(&machines)
	if err != nil {
		return nil, err
	}

	return machines, nil
}

func (m Machine) GetAll() ([]*Machine, error) {
	var machines []*Machine
	if _, err := DB.QueryTable(m.TableName()).All(&machines); err != nil {
		return nil, err
	}

	return machines, nil
}

func (m Machine) GetOne(id int) (*Machine, error) {
	machine := &Machine{
		Id: id,
	}

	if err := DB.Read(machine); err != nil {
		return nil, err
	}

	return machine, nil
}

func (m Machine) GetOneByIp(ip string) (*Machine, error) {
	machine := &Machine{}

	if err := DB.QueryTable(m.TableName()).Filter("Ip", ip).Limit(1).One(machine); err != nil {
		return nil, err
	}

	return machine, nil
}

func (m Machine) Add(ip, port, note string) (*Machine, error) {
	machine := &Machine{
		Ip:         ip,
		Note:       note,
		Created:    time.Now(),
		Modified:   time.Now(),
		LastReport: time.Now(),
	}

	id, err := DB.Insert(machine)
	if err != nil {
		return nil, err
	}

	machine, err = m.GetOne(int(id))
	if err != nil {
		return nil, err
	}

	return machine, nil
}

func (m Machine) Update(id int, ip, port, note, token string, status int, lastReport time.Time) (*Machine, error) {
	machine := &Machine{
		Id:         id,
		Ip:         ip,
		Port:       port,
		Note:       note,
		Token:      token,
		Status:     status,
		Modified:   time.Now(),
		LastReport: lastReport,
	}
	_, err := DB.Update(machine, "Ip", "Port", "Note", "Token", "Status", "Modified", "LastReport")
	if err != nil {
		return nil, err
	}

	return machine, nil
}

// set status
func (m Machine) Del(id int) error {
	machine := &Machine{
		Id: id,
	}
	if _, err := DB.Delete(machine); err != nil {
		return err
	}

	return nil
}
