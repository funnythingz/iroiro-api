package iro

import (
	"../../ddd"
	"../../domain"
	"../../mapper"
	_ "github.com/k0kubun/pp"
)

type Factory struct{}

func (f *Factory) CreateIro(im mapper.Iro) domain.Iro {
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

func (f *Factory) CreateIroIro(iim mapper.IroIro) domain.IroIro {
	iroiro := []domain.Iro{}
	for _, im := range iim.IroIro {
		iroiro = append(iroiro, f.CreateIro(im))
	}
	return domain.IroIro{IroIro: iroiro}
}

var (
	factory = Factory{}
)
