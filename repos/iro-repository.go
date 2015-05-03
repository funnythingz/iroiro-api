package repos

import (
	"../db"
	"../domain"
	"../tables"
	"fmt"
	"github.com/k0kubun/pp"
)

type IroRepository struct{}

func (r *IroRepository) Commit(iro domain.Iro) string {
	iroTable := tables.Iro{
		ColorName: iro.Color.Name,
		ColorCode: iro.Color.Code,
		Content:   iro.Content,
	}

	db.Dbmap.NewRecord(iroTable)
	db.Dbmap.Create(&iroTable)

	pp.Println(iroTable)

	return fmt.Sprintf("Commited ID: %d", iroTable.Id)
}

func (r *IroRepository) Fetch(id int) domain.Iro {
	var iroTable = tables.Iro{}
	db.Dbmap.Where(&tables.Iro{Id: id}).First(&iroTable)

	iro := domain.Iro{
		Id: iroTable.Id,
		Color: domain.Color{
			Name: iroTable.ColorName,
			Code: iroTable.ColorCode,
		},
		Content:   iroTable.Content,
		CreatedAt: iroTable.CreatedAt,
		UpdatedAt: iroTable.UpdatedAt,
	}

	return iro
}

func (r *IroRepository) FetchList(permit int) domain.IroIro {
	var iroTables = []tables.Iro{}
	db.Dbmap.Order("id desc").Limit(permit).Find(&iroTables)

	iroiro := []domain.Iro{}
	for _, iroTable := range iroTables {
		iroiro = append(iroiro, domain.Iro{
			Id: iroTable.Id,
			Color: domain.Color{
				Name: iroTable.ColorName,
				Code: iroTable.ColorCode,
			},
			Content:   iroTable.Content,
			CreatedAt: iroTable.CreatedAt,
			UpdatedAt: iroTable.UpdatedAt,
		})
	}

	return domain.IroIro{IroIro: iroiro}
}
