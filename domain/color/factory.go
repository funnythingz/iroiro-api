package color

import (
	"../../ddd"
	"../../domain"
	"../../mapper"
	_ "github.com/k0kubun/pp"
)

type Factory struct{}

func (f *Factory) CreateColor(mc mapper.Color) domain.Color {
	return domain.Color{
		Entity: ddd.Entity{
			Id: mc.EntityMapper.Id,
		},
		Name: mc.Name,
		Code: mc.Code,
	}
}

func (f *Factory) CreateColorList(mcl mapper.ColorList) domain.ColorList {
	colorList := []domain.Color{}
	for _, mc := range mcl.ColorList {
		colorList = append(colorList, f.CreateColor(mc))
	}
	return domain.ColorList{ColorList: colorList}
}

var (
	factory = Factory{}
)
