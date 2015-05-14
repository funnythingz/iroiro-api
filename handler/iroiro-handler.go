package handler

import (
	"../domain"
	"../domain/iro"
	"../domain/service"
	"../helper"
	"encoding/json"
	"fmt"
	_ "github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"unicode/utf8"
)

type IroiroHandler struct{}

func (h *IroiroHandler) Iroiro(c web.C, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	permit := 10
	urlQuery, _ := url.ParseQuery(r.URL.RawQuery)

	var page int
	if len(urlQuery["page"]) == 0 {
		page = 1
	} else {
		page, _ = strconv.Atoi(urlQuery["page"][0])
	}

	iroiro := iro.FetchList(permit, page)
	response, _ := json.Marshal(iroiro)
	io.WriteString(w, string(response))
}

func (h *IroiroHandler) Iro(c web.C, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	id, _ := strconv.Atoi(c.URLParams["id"])
	iro := iro.Fetch(id)
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
	colorId, _ := strconv.Atoi(r.FormValue("iro[color_id]"))
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

	i := domain.Iro{
		ColorId: colorId,
		Content: content,
		ReIroId: reIroId,
	}

	resultIro := iro.Commit(i)
	response, _ := json.Marshal(resultIro)
	io.WriteString(w, string(response))
}
