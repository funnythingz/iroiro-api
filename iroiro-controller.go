package main

import (
	"./models"
	"encoding/json"
	_ "github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"io"
	"log"
	"net/http"
	"strconv"
)

type IroiroController struct{}

func (_ *IroiroController) iroiro(c web.C, w http.ResponseWriter, r *http.Request) {
	// dummy
	iroiro := models.IroIro{
		IroIro: []models.Iro{
			models.Iro{
				Id:      1,
				Content: "Hello world.",
				Color:   models.Purple500,
			},
			models.Iro{
				Id:      2,
				Content: "Hello good day",
				Color:   models.Pink500,
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
	iro := models.Iro{Id: id, Content: "hey!", Color: models.Purple500}
	response, _ := json.Marshal(iro)
	io.WriteString(w, string(response))
}

func (_ *IroiroController) create(c web.C, w http.ResponseWriter, r *http.Request) {
	// TODO: createする
}
