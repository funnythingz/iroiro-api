package mapper

import (
	"../db"
	"../ddd"
	"../domain"
	_ "github.com/k0kubun/pp"
)

type Iro struct {
	ddd.EntityMapper
	ColorId  int
	Color    Color
	Content  string
	ReIroId  int
	ReIroIro []domain.Iro
}

func (m *Iro) Map(iro domain.Iro) {
	m.ColorId = iro.Color.Entity.Id
	m.Content = iro.Content
	m.ReIroId = iro.ReIroId
	m.ReIroIro = iro.ReIroIro
}

func (m *Iro) Commit() {
	db.Dbmap.NewRecord(m)
	db.Dbmap.Create(&m)
}

func (m *Iro) Fetch(id int) {
	db.Dbmap.Find(&m, id).First(&m)
	m.Color.Fetch(m.ColorId)
}

type IroIro struct {
	IroIro []Iro
}

func (m *IroIro) Fetch(permit int, page int) {
	db.Dbmap.Order("id desc").Offset((page - 1) * permit).Limit(permit).Find(&m.IroIro).Offset(page * permit).Limit(permit)
	for i, iro := range m.IroIro {
		m.IroIro[i].Fetch(iro.Id)
	}
}
