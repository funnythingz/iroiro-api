package iro

import (
	"../../db"
	"../../domain"
	_ "github.com/k0kubun/pp"
)

type Repository struct{}

func (r *Repository) Commit(iro domain.Iro) domain.Iro {
	db.Dbmap.NewRecord(iro)
	db.Dbmap.Create(&iro)
	return iro
}

func (r *Repository) Fetch(id int) domain.Iro {
	iro := domain.Iro{}
	db.Dbmap.Where(&domain.Iro{Id: id}).First(&iro)
	if iro.Id == 0 {
		return iro
	}
	db.Dbmap.Model(&iro).Related(&iro.Color)
	return iro
}

func (r *Repository) FetchList(permit int, page int) domain.IroIro {
	iroiro := []domain.Iro{}
	db.Dbmap.Order("id desc").Offset((page - 1) * permit).Limit(permit).Find(&iroiro).Offset(page * permit).Limit(permit)
	for i, iro := range iroiro {
		db.Dbmap.Model(&iro).Related(&iroiro[i].Color)
	}
	return domain.IroIro{IroIro: iroiro}
}

var (
	repo = &Repository{}
)

func Commit(iro domain.Iro) domain.Iro {
	return repo.Commit(iro)
}

func Fetch(id int) domain.Iro {
	return repo.Fetch(id)
}

func FetchList(permit int, page int) domain.IroIro {
	return repo.FetchList(permit, page)
}
