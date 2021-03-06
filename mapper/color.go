package mapper

import (
	"github.com/funnythingz/iroiro-api/db"
	"github.com/funnythingz/iroiro-api/ddd"
	"github.com/funnythingz/iroiro-api/domain"
	_ "github.com/k0kubun/pp"
)

type Color struct {
	ddd.EntityMapper
	Name     string
	Code     string
	TextCode string
}

func (m *Color) Map(color domain.Color) {
	m.Name = color.Name
	m.Code = color.Code
	m.TextCode = color.TextCode
}

func (m *Color) Commit() {
	db.Dbmap.NewRecord(m)
	db.Dbmap.Create(&m)
}

func (m *Color) Fetch(id int) {
	db.Dbmap.Find(&m, id).First(&m)
}

type ColorList struct {
	ColorList []Color
}

func (m *ColorList) Fetch(permit int, page int) {
	db.Dbmap.Order("id desc").Offset((page - 1) * permit).Limit(permit).Find(&m.ColorList).Offset(page * permit).Limit(permit)
}
