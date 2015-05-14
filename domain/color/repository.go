package color

import (
	"../../db"
	"../../domain"
	_ "github.com/k0kubun/pp"
)

type Repository struct{}

func (r *Repository) Commit(color domain.Color) domain.Color {
	db.Dbmap.NewRecord(color)
	db.Dbmap.Create(&color)
	return color
}

func (r *Repository) Fetch(id int) domain.Color {
	color := domain.Color{}
	db.Dbmap.Where(&domain.Color{Id: id}).First(&color)
	return color
}

func (r *Repository) FetchList(permit int, page int) domain.ColorList {
	colorList := []domain.Color{}
	db.Dbmap.Order("id desc").Offset((page - 1) * permit).Limit(permit).Find(&colorList).Offset(page * permit).Limit(permit)
	return domain.ColorList{ColorList: colorList}
}

var (
	repo = &Repository{}
)

func Commit(color domain.Color) domain.Color {
	return repo.Commit(color)
}

func Fetch(id int) domain.Color {
	return repo.Fetch(id)
}

func FetchList(permit int, page int) domain.ColorList {
	return repo.FetchList(permit, page)
}
