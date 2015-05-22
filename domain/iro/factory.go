package iro

import (
	"../../ddd"
	"../../domain"
	"../../mapper"
	_ "github.com/k0kubun/pp"
)

type Factory struct{}

func (f *Factory) CreateIro(mi mapper.Iro) domain.Iro {
	return domain.Iro{
		Entity: ddd.Entity{
			Id: mi.EntityMapper.Id,
		},
		Color: domain.Color{
			Entity: ddd.Entity{
				Id: mi.Color.EntityMapper.Id,
			},
			Name: mi.Color.Name,
			Code: mi.Color.Code,
		},
		Content:  mi.Content,
		ReIroId:  mi.ReIroId,
		ReIroIro: mi.ReIroIro,
	}
}

func (f *Factory) CreateIroIro(mii mapper.IroIro) domain.IroIro {
	iroiro := []domain.Iro{}
	for _, mi := range mii.IroIro {
		iroiro = append(iroiro, f.CreateIro(mi))
	}
	return domain.IroIro{IroIro: iroiro}
}

var (
	factory = Factory{}
)
