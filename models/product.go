package models

import "time"

type Product struct {
	Id      int
	Name    string
	Note    string
	Created time.Time
}

func (m Product) GetAll() ([]Product, error) {
}

func (m Product) GetOne(id int) (Product, error) {
}

func (m Product) Add(name, note string) (int, error) {
}
