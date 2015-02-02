package models

import "time"

var (
	ProductModel = &Product{}
)

type Product struct {
	Id      int
	Name    string
	Note    string
	Created time.Time
}

func (m Product) TableName() string {
	return "product"
}

func (m Product) TableUnique() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

func (m Product) GetAll() ([]*Product, error) {
	products := []*Product{}
	if _, err := DB.QueryTable(m.TableName()).All(&products); err != nil {
		return nil, err
	}

	return products, nil
}

func (m Product) GetOne(id int) (*Product, error) {
	product := &Product{
		Id: id,
	}
	if err := DB.Read(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (m Product) Add(name, note string) (*Product, error) {
	product := &Product{
		Name:    name,
		Note:    note,
		Created: time.Now(),
	}
	id, err := DB.Insert(product)
	if err != nil {
		return nil, err
	}

	if product, err = m.GetOne(int(id)); err != nil {
		return nil, err
	}

	return product, nil
}

func (m Product) Update(id int, name, note string) error {
	product := &Product{
		Id:      id,
		Name:    name,
		Note:    note,
		Created: time.Now(),
	}

	if _, err := DB.Update(product); err != nil {
		return err
	}

	return nil
}

func (m Product) Del(id int) error {
	product := &Product{
		Id: id,
	}
	if _, err := DB.Delete(product); err != nil {
		return err
	}

	return nil
}
