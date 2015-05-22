package domain

import (
	"../ddd"
)

type Color struct {
	ddd.Entity
	Name string `json:"name"`
	Code string `json:"code"`
}

type ColorList struct {
	ColorList []Color `json:"color_list"`
}
