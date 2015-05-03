package main

import (
	"./domain"
	"./repos"
	"encoding/json"
	_ "github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"io"
	"log"
	"net/http"
	"strconv"
	"unicode/utf8"
)

var (
	iroRepository = repos.IroRepository{}
)

type IroiroController struct{}

func (_ *IroiroController) iroiro(c web.C, w http.ResponseWriter, r *http.Request) {
	// dummy
	iroiro := domain.IroIro{
		IroIro: []domain.Iro{
			domain.Iro{
				Id:      1,
				Content: "Hello world.",
				Color:   domain.Purple500,
			},
			domain.Iro{
				Id:      2,
				Content: "Hello good day",
				Color:   domain.Pink500,
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
	iro := domain.Iro{Id: id, Content: "hey!", Color: domain.Purple500}
	response, _ := json.Marshal(iro)
	io.WriteString(w, string(response))
}

func (_ *IroiroController) create(c web.C, w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("iro[content]")
	colorName := r.FormValue("iro[color_name]")
	colorCode := r.FormValue("iro[color_code]")

	// Validation
	errors := []string{}

	if utf8.RuneCountInString(content) <= 0 {
		errors = append(errors, "input Content must be blank.")
	}
	if utf8.RuneCountInString(content) < 5 || utf8.RuneCountInString(content) > 1000 {
		errors = append(errors, "input Content minimum is 5 and maximum is 1000 character.")
	}

	if len(errors) > 0 {
		resultJSON(w, errors)
		return
	}

	iro := domain.Iro{
		Color:   domain.Color{Name: colorName, Code: colorCode},
		Content: content,
	}

	message := iroRepository.Store(iro)
	resultJSON(w, []string{message})
}
