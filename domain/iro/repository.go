package iro

import (
	"../../domain"
	"../../mapper"
	_ "github.com/k0kubun/pp"
)

type Repository struct{}

func (r *Repository) Commit(iro domain.Iro) domain.Iro {
	im := mapper.Iro{}
	im.Map(iro)
	im.Commit()
	return r.Fetch(im.Id)
}

func (r *Repository) Fetch(id int) domain.Iro {
	im := mapper.Iro{}
	im.Fetch(id)
	return factory.CreateIro(im)
}

func (r *Repository) FetchList(permit int, page int) domain.IroIro {
	iim := mapper.IroIro{}
	iim.Fetch(permit, page)
	return factory.CreateIroIro(iim)
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
