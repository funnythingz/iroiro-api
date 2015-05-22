package domain

import (
	"../ddd"
)

type Iro struct {
	ddd.Entity
	Color    Color  `json:"color"`
	Content  string `json:"content"`
	ReIroId  int    `json:"re_iro_id"`
	ReIroIro []Iro  `json:"re_iroiro"`
}

type IroIro struct {
	IroIro []Iro `json:"iroiro"`
}
