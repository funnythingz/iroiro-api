package color

import (
	"../../domain"
	"../../mapper"
	_ "github.com/k0kubun/pp"
)

type ColorRepository struct{}

func (r *ColorRepository) Commit(color domain.Color) domain.Color {
	mc := mapper.Color{}
	mc.Map(color)
	mc.Commit()
	return r.Fetch(mc.Id)
}

func (r *ColorRepository) Fetch(id int) domain.Color {
	mc := mapper.Color{}
	mc.Fetch(id)
	return factory.CreateColor(mc)
}

func (r *ColorRepository) FetchList(permit int, page int) domain.ColorList {
	mcl := mapper.ColorList{}
	mcl.Fetch(permit, page)
	return factory.CreateColorList(mcl)
}

var (
	Repository = &ColorRepository{}
)
