package repos

import (
	"../db"
	"../domain"
	"../tables"
	"fmt"
	"github.com/k0kubun/pp"
)

type IroRepository struct{}

func (r *IroRepository) Store(iro domain.Iro) string {
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
