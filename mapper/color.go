package mapper

import (
	"../db"
	"../ddd"
	"../domain"
)

type Color struct {
	ddd.EntityMapper
	Name string
	Code string
}

func (m *Color) Map(color domain.Color) {
	m.Name = color.Name
	m.Code = color.Code
}

func (m *Color) Commit() {
	db.Dbmap.NewRecord(m)
	db.Dbmap.Create(&m)
}

func (m *Color) Fetch(id int) {
	db.Dbmap.Find(&m, id).First(&m)
}
