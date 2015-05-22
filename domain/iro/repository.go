package iro

import (
	"../../domain"
	"../../mapper"
	_ "github.com/k0kubun/pp"
)

type IroRepository struct{}

func (r *IroRepository) Commit(iro domain.Iro) domain.Iro {
	mi := mapper.Iro{}
	mi.Map(iro)
	mi.Commit()
	return r.Fetch(mi.Id)
}

func (r *IroRepository) Fetch(id int) domain.Iro {
	mi := mapper.Iro{}
	mi.Fetch(id)
	return factory.CreateIro(mi)
}

func (r *IroRepository) FetchList(permit int, page int) domain.IroIro {
	mii := mapper.IroIro{}
	mii.Fetch(permit, page)
	return factory.CreateIroIro(mii)
}

var (
	Repository = &IroRepository{}
)
