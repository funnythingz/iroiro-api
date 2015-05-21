package mapper

import (
	"../ddd"
)

type IroMapper struct {
	ddd.EntityMapper
	ColorCode string
	Content   string
	ReIroIro  []Iro
	ReIroId   int
}

func (m *IroMapper) New(iro domain.Iro) {
	m.ColorCode = iro.Color.Code
	m.Content = iro.Content
	m.ReIroIro = iro.ReIroIro
	m.ReIroId = iro.ReIroId
}

func (m *IroMapper) Commit() {
	db.Dbmap.NewRecord(m)
	db.Dbmap.Create(&m)
}

func (m *IroMapper) Fetch(id int) {
	db.Dbmap.Find(&m, id).First(&m)
}
