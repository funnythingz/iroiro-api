package main

type Color struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

var (
	purple500 = Color{Name: "Purple500", Code: "#9C27B0"}
	pink500   = Color{Name: "Pink500", Code: "#E91E63"}
)
