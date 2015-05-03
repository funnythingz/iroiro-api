package main

import (
	"./domain"
	"./repos"
	"encoding/json"
	"fmt"
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
	iroiro := iroRepository.FetchList(10)
	response, err := json.Marshal(iroiro)
	if err != nil {
		log.Println(w, err)
	}

	io.WriteString(w, string(response))
}

func (_ *IroiroController) iro(c web.C, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(c.URLParams["id"])
	iro := iroRepository.Fetch(id)
	if iro.Id == 0 {
		resultJSON(w, []string{fmt.Sprintf("Not Found: %d", id)})
		return
	}
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

	message := iroRepository.Commit(iro)
	resultJSON(w, []string{message})
}
