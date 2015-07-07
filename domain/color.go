package domain

import (
	"github.com/funnythingz/iroiro-api/ddd"
)

type Color struct {
	ddd.Entity
	Name     string `json:"name"`
	Code     string `json:"code"`
	TextCode string `json:"text_code"`
}
