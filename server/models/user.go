package models

import "time"

var (
	UserModel = &User{}
)

type User struct {
	Id            int       `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Name          string    `json:"name"`
	Administrator int       `json:"administrator"`
	Created       time.Time `json:"created"`
	LastLogin     time.Time `json:"last_login"`
	LastLoginIp   string    `json:"last_login_ip"`
}

func (m User) TableName() string {
	return "user"
}

func (m User) TableUnique() [][]string {
	return [][]string{
		[]string{"Username"},
	}
}

func (m User) CheckPassword(username, password string) (*User, error) {
	user := &User{
		Username: username,
		Password: password,
	}

	if err := DB.Read(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (m User) GetOne(id int) (*User, error) {
	user := &User{
		Id: id,
	}

	if err := DB.Read(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (m User) GetViaUsername(username string) (*User, error) {
	user := &User{
		Username: username,
	}
	if err := DB.Read(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (m User) GetAll() ([]*User, error) {
	users := []*User{}
	if _, err := DB.QueryTable(m.TableName()).All(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (m User) Add(username, name, password string, administrator int) (*User, error) {
	user := &User{
		Username:      username,
		Name:          name,
		Password:      password,
		Administrator: administrator,
		Created:       time.Now(),
		LastLogin:     time.Now().AddDate(-1, 0, 0),
	}
	id, err := DB.Insert(user)
	if err != nil {
		return nil, err
	}

	user, err = m.GetOne(int(id))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m User) Update(id int, username, name string, administrator int) (*User, error) {
	user := &User{
		Id:            id,
		Username:      username,
		Name:          name,
		Administrator: administrator,
	}
	_, err := DB.Update(user, "Username", "Name", "Administrator")
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m User) UpdatePassword(id int, password string) error {
	user := &User{
		Id:       id,
		Password: password,
	}
	_, err := DB.Update(user, "Password")
	if err != nil {
		return err
	}

	return nil
}

func (m User) UpdateLastLogin(id int, ip string) error {
	user := &User{
		Id:          id,
		LastLogin:   time.Now(),
		LastLoginIp: ip,
	}
	_, err := DB.Update(user, "LastLogin", "LastLoginIp")
	if err != nil {
		return err
	}

	return nil
}

func (m User) Del(id int) error {
	user := &User{
		Id: id,
	}
	if _, err := DB.Delete(user); err != nil {
		return err
	}

	return nil
}
