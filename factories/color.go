package factories

import (
	"github.com/funnythingz/iroiro-api/ddd"
	"github.com/funnythingz/iroiro-api/domain"
	"github.com/funnythingz/iroiro-api/mapper"
	_ "github.com/k0kubun/pp"
)

type ColorFactory struct{}

func (f *ColorFactory) CreateColor(mc mapper.Color) domain.Color {
	return domain.Color{
		Entity: ddd.Entity{
			Id: mc.EntityMapper.Id,
		},
		Name:     mc.Name,
		Code:     mc.Code,
		TextCode: mc.TextCode,
	}
}

func (f *ColorFactory) CreateColorList(mcl mapper.ColorList) []domain.Color {
	colorList := []domain.Color{}
	for _, mc := range mcl.ColorList {
		colorList = append(colorList, f.CreateColor(mc))
	}
	return colorList
}

var (
	ColorFact = &ColorFactory{}
)
