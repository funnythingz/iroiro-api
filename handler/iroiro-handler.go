package handler

import (
	"encoding/json"
	"fmt"
	"github.com/funnythingz/iroiro-api/ddd"
	"github.com/funnythingz/iroiro-api/domain"
	"github.com/funnythingz/iroiro-api/helper"
	"github.com/funnythingz/iroiro-api/repositories"
	"github.com/funnythingz/iroiro-api/services"
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

	permit := 50
	urlQuery, _ := url.ParseQuery(r.URL.RawQuery)

	var page int
	if len(urlQuery["page"]) == 0 {
		page = 1
	} else {
		page, _ = strconv.Atoi(urlQuery["page"][0])
	}

	iroiro := repositories.IroRepo.ResolveList(permit, page)
	response, _ := json.Marshal(iroiro)
	io.WriteString(w, string(response))
}

func (h *IroiroHandler) Iro(c web.C, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	id, _ := strconv.Atoi(c.URLParams["id"])
	i := repositories.IroRepo.Resolve(id)
	if i.Id == 0 {
		helper.ResultMessageJSON(w, []string{fmt.Sprintf("Not Found: %d", id)})
		return
	}
	response, _ := json.Marshal(i)
	io.WriteString(w, string(response))
}

func (h *IroiroHandler) CreateIro(c web.C, w http.ResponseWriter, r *http.Request) {

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

	iro := domain.Iro{
		Color: domain.Color{
			Entity: ddd.Entity{
				Id: colorId,
			},
		},
		Content: content,
		ReIroId: reIroId,
	}

	repositories.IroRepo.Store(&iro)
	response, _ := json.Marshal(iro)
	io.WriteString(w, string(response))
}

func (h *IroiroHandler) UpdateIro(c web.C, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	content := r.FormValue("iro[content]")
	//TODO
	//colorId, _ := strconv.Atoi(r.FormValue("iro[color_id]"))
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

	// Fetch
	id, _ := strconv.Atoi(c.URLParams["id"])
	iro := repositories.IroRepo.Resolve(id)

	if iro.Id == 0 {
		helper.ResultMessageJSON(w, []string{fmt.Sprintf("Not Found: %s", c.URLParams["id"])})
		return
	}

	//TODO
	iro.ReIroId = reIroId

	repositories.IroRepo.Update(&iro)
	response, _ := json.Marshal(iro)
	io.WriteString(w, string(response))
}
