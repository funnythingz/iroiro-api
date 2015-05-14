package domain

import (
	"time"
)

type Iro struct {
	Id        int       `json:"id"`
	ColorId   int       `json:"color_id"`
	Color     Color     `json:"color"`
	Content   string    `json:"content"`
	ReIroIro  []Iro     `json:"re_iroiro"`
	ReIroId   int       `json:"re_iro_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IroIro struct {
	IroIro []Iro `json:"iroiro"`
}
