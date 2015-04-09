package main

type Iro struct {
	Id      int     `json:"id"`
	Comment Comment `json:"comment"`
	Color   Color   `json:color`
}

type Iroiro struct {
	Iroiro []Iro `json:"iroiro"`
}
