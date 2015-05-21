package mapper

import (
	"../db"
	"../ddd"
	"../domain"
	"github.com/k0kubun/pp"
)

type Iro struct {
	ddd.EntityMapper
	ColorId  int
	Color    Color
	Content  string
	ReIroId  int
	ReIroIro []domain.Iro
}

func (m *Iro) New(iro domain.Iro) {
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
	color := Color{}
	db.Dbmap.Find(&m, id).First(&m).Related(&color)
	m.Color = color
	pp.Println(m, color)
}
