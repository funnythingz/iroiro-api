package handler

import (
	"../domain"
	"../domain/service"
	"../helper"
	"../repos"
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

type IroiroHandler struct{}

func (h *IroiroHandler) Iroiro(c web.C, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	iroiro := iroRepository.FetchList(20)
	response, err := json.Marshal(iroiro)
	if err != nil {
		log.Println(w, err)
	}

	io.WriteString(w, string(response))
}

func (h *IroiroHandler) Iro(c web.C, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	id, _ := strconv.Atoi(c.URLParams["id"])
	iro := iroRepository.Fetch(id)
	if iro.Id == 0 {
		helper.ResultMessageJSON(w, []string{fmt.Sprintf("Not Found: %d", id)})
		return
	}
	response, _ := json.Marshal(iro)
	io.WriteString(w, string(response))
}

func (h *IroiroHandler) Create(c web.C, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	content := r.FormValue("iro[content]")
	colorName := r.FormValue("iro[color_name]")
	colorCode := r.FormValue("iro[color_code]")
	reIroId, _ := strconv.Atoi(r.FormValue("iro[re_iro_id]"))

	// Validation
	errors := []string{}

	if utf8.RuneCountInString(content) <= 0 {
		errors = append(errors, "input Content must be blank.")
	}
	if utf8.RuneCountInString(content) < 5 || utf8.RuneCountInString(content) > 1000 {
		errors = append(errors, "input Content minimum is 5 and maximum is 1000 character.")
	}

	if len(errors) > 0 {
		helper.ResultMessageJSON(w, errors)
		return
	}

	iro := domain.Iro{
		Color:   domain.Color{Name: colorName, Code: colorCode},
		Content: content,
		ReIroId: reIroId,
	}

	message := iroRepository.Commit(iro)
	helper.ResultMessageJSON(w, []string{message})
}
