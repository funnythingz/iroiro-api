package repos

import (
	"../domain"
	"../tables"
)

type IroRepository struct{}

func (r *IroRepository) Store(iro domain.Iro) {
	iroTable := tables.Iro{
		Id:        iro.Id,
		ColorName: iro.Color.Name,
		ColorCode: iro.Color.Code,
		Content:   iro.Content,
		UpdatedAt: iro.UpdatedAt,
	}
}
