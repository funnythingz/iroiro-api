package main

import (
	"time"
)

type Iro struct {
	Id        int       `json:"id"`
	Color     Color     `json:"color"`
	Content   string    `json:"content"`
	Fan       []Iro     `json:"fan"`
	ReIroId   int       `json:"re_iro_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Iroiro struct {
	Iroiro []Iro `json:"iroiro"`
}
