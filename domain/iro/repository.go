package iro

import (
	"../../ddd"
	"../../domain"
	"../../mapper"
	_ "github.com/k0kubun/pp"
)

type Repository struct{}

func (r *Repository) Commit(iro domain.Iro) domain.Iro {
	im := mapper.Iro{}
	im.New(iro)
	im.Commit()
	return r.Fetch(im.Id)
}

func (r *Repository) Fetch(id int) domain.Iro {
	im := mapper.Iro{}
	im.Fetch(id)
	return domain.Iro{
		Entity: ddd.Entity{
			Id: im.EntityMapper.Id,
		},
		Color: domain.Color{
			Entity: ddd.Entity{
				Id: im.Color.EntityMapper.Id,
			},
			Name: im.Color.Name,
			Code: im.Color.Code,
		},
		Content:  im.Content,
		ReIroId:  im.ReIroId,
		ReIroIro: im.ReIroIro,
	}
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
