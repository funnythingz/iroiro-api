package domain

import (
	"../ddd"
)

type Iro struct {
	ddd.Entity
	ColorId  int    `json:"color_id"`
	Color    Color  `json:"color"`
	Content  string `json:"content"`
	ReIroIro []Iro  `json:"re_iroiro"`
	ReIroId  int    `json:"re_iro_id"`
}

type IroIro struct {
	IroIro []Iro `json:"iroiro"`
}
