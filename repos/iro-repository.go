package repos

import (
	"../db"
	"../domain"
	"../tables"
	"fmt"
	_ "github.com/k0kubun/pp"
)

type IroRepository struct{}

func (r *IroRepository) Commit(iro domain.Iro) string {
	iroTable := tables.Iro{
		ColorName: iro.Color.Name,
		ColorCode: iro.Color.Code,
		Content:   iro.Content,
		ReIroId:   iro.ReIroId,
	}

	db.Dbmap.NewRecord(iroTable)
	db.Dbmap.Create(&iroTable)

	return fmt.Sprintf("Commited ID: %d", iroTable.Id)
}

func (r *IroRepository) Fetch(id int) domain.Iro {
	var iroTable = tables.Iro{}
	var iroTables = []tables.Iro{}

	db.Dbmap.Where(&tables.Iro{Id: id}).First(&iroTable)
	db.Dbmap.Where(&tables.Iro{ReIroId: id}).Find(&iroTables)
	iroTable.Iros = iroTables
	return r.mapIro(iroTable)
}

func (r *IroRepository) FetchList(permit int) domain.IroIro {
	var iroTables = []tables.Iro{}
	db.Dbmap.Order("id desc").Limit(permit).Find(&iroTables)
	return domain.IroIro{IroIro: r.mapIros(iroTables)}
}

func (r *IroRepository) mapIro(iroTable tables.Iro) domain.Iro {
	return domain.Iro{
		Id: iroTable.Id,
		Color: domain.Color{
			Name: iroTable.ColorName,
			Code: iroTable.ColorCode,
		},
		Fan:       r.mapIros(iroTable.Iros),
		ReIroId:   iroTable.ReIroId,
		Content:   iroTable.Content,
		CreatedAt: iroTable.CreatedAt,
		UpdatedAt: iroTable.UpdatedAt,
	}
}

func (r *IroRepository) mapIros(iroTables []tables.Iro) []domain.Iro {
	iroiro := []domain.Iro{}
	for _, iroTable := range iroTables {
		iroiro = append(iroiro, domain.Iro{
			Id: iroTable.Id,
			Color: domain.Color{
				Name: iroTable.ColorName,
				Code: iroTable.ColorCode,
			},
			ReIroId:   iroTable.ReIroId,
			Content:   iroTable.Content,
			CreatedAt: iroTable.CreatedAt,
			UpdatedAt: iroTable.UpdatedAt,
		})
	}
	return iroiro
}
