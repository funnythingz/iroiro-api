package main

import (
	"encoding/json"
	_ "github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"io"
	"net/http"
	"strconv"
)

var (
	comment = Comment{Value: "hello IROIRO."}
	color   = Color{Name: "purple", Code: "#C55BFF"}
)

type IroiroController struct{}

func (_ *IroiroController) iroiro(c web.C, w http.ResponseWriter, r *http.Request) {
	// dummy
	iroiro := Iroiro{
		Iroiro: []Iro{
			Iro{
				Id:      1,
				Comment: comment,
				Color:   color,
			},
			Iro{
				Id:      2,
				Comment: comment,
				Color:   color,
			},
		},
	}

	response, _ := json.Marshal(iroiro)
	io.WriteString(w, string(response))
}

func (_ *IroiroController) iro(c web.C, w http.ResponseWriter, r *http.Request) {
	// dummy
	id, _ := strconv.Atoi(c.URLParams["id"])
	iro := Iro{Id: id, Comment: comment, Color: color}
	response, _ := json.Marshal(iro)
	io.WriteString(w, string(response))
}

func (_ *IroiroController) create(c web.C, w http.ResponseWriter, r *http.Request) {
	// TODO: createする
}
