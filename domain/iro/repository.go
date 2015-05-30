package iro

import (
	"../../domain"
	"../../mapper"
	_ "github.com/k0kubun/pp"
)

type IroRepository struct{}

func (r *IroRepository) Store(iro domain.Iro) domain.Iro {
	mi := mapper.Iro{}
	mi.Map(iro)
	mi.Commit()
	return r.Resolve(mi.Id)
}

func (r *IroRepository) Resolve(id int) domain.Iro {
	mi := mapper.Iro{}
	mi.Fetch(id)
	return factory.CreateIro(mi)
}

func (r *IroRepository) ResolveList(permit int, page int) domain.IroIro {
	mii := mapper.IroIro{}
	mii.Fetch(permit, page)
	return factory.CreateIroIro(mii)
}

var (
	Repository = IroRepository{}
)
