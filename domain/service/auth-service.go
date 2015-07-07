package service

import (
	"github.com/funnythingz/iroiro-api/config"
	"net/http"
	"net/url"
)

type Auth struct{}

func (a *Auth) BeforeAuth(w http.ResponseWriter, r *http.Request) bool {
	t := a.ParseAccessKeyFromRequest(r)
	isAuth := a.CheckAuth(t)
	if isAuth == false {
		http.Redirect(w, r, "/", 404)
	}
	return isAuth
}

func (a *Auth) ParseAccessKeyFromRequest(r *http.Request) string {
	u, _ := url.Parse(r.RequestURI)
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return err.Error()
	}
	if q["access_key"] != nil {
		return q["access_key"][0]
	}
	return ""
}

func (a *Auth) CheckAuth(accessKey string) bool {
	return accessKey == config.Params.Auth.AccessKey
}

var AuthService = Auth{}

func BeforeAuth(w http.ResponseWriter, r *http.Request) bool {
	return AuthService.BeforeAuth(w, r)
}

func ParseAccessKeyFromRequest(r *http.Request) string {
	return AuthService.ParseAccessKeyFromRequest(r)
}
