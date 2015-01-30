package models

import "time"

type Product struct {
	Id      int
	Name    string
	Note    string
	Created time.Time
}

func (m Product) TableName() string {
	return "product"
}

func (m Product) GetAll() ([]*Product, error) {
	var products []*Product
	_, err := DB.QueryTable(m.TableName()).All(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (m Product) GetOne(id int) (*Product, error) {
	var product *Product = &Product{
		Id: id,
	}
	err := DB.Read(&product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (m Product) Add(name, note string) (int, error) {
	var product *Product = &Product{
		Name:    name,
		Note:    note,
		Created: time.Now(),
	}
	id, err := DB.Insert(product)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
