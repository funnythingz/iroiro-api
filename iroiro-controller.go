package main

import (
	"encoding/json"
	_ "github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	purple500 = Color{Name: "Purple500", Code: "#9C27B0"}
	pink500   = Color{Name: "Pink500", Code: "#E91E63"}
)

type IroiroController struct{}

func (_ *IroiroController) iroiro(c web.C, w http.ResponseWriter, r *http.Request) {
	// dummy
	iroiro := Iroiro{
		Iroiro: []Iro{
			Iro{
				Id:      1,
				Content: "Hello world.",
				Color:   purple500,
			},
			Iro{
				Id:      2,
				Content: "Hello good day",
				Color:   pink500,
			},
		},
	}

	response, err := json.Marshal(iroiro)
	if err != nil {
		log.Println(w, err)
	}

	io.WriteString(w, string(response))
}

func (_ *IroiroController) iro(c web.C, w http.ResponseWriter, r *http.Request) {
	// dummy
	id, _ := strconv.Atoi(c.URLParams["id"])
	iro := Iro{Id: id, Content: "hey!", Color: purple500}
	response, _ := json.Marshal(iro)
	io.WriteString(w, string(response))
}

func (_ *IroiroController) create(c web.C, w http.ResponseWriter, r *http.Request) {
	// TODO: createする
}
