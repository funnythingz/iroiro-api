package iro

import (
	"github.com/funnythingz/iroiro-api/ddd"
	"github.com/funnythingz/iroiro-api/domain"
	"github.com/funnythingz/iroiro-api/mapper"
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
			Name:     mi.Color.Name,
			Code:     mi.Color.Code,
			TextCode: mi.Color.TextCode,
		},
		Content:  mi.Content,
		ReIroId:  mi.ReIroId,
		ReIroIro: mi.ReIroIro,
	}
}

func (f *Factory) CreateIroIro(mii mapper.IroIro) []domain.Iro {
	iroiro := []domain.Iro{}
	for _, mi := range mii.IroIro {
		iroiro = append(iroiro, f.CreateIro(mi))
	}
	return iroiro
}

var (
	factory = Factory{}
)
