package main

import (
	"github.com/funnythingz/iroiro-api/db"
	"github.com/funnythingz/iroiro-api/handler"
	_ "github.com/zenazn/goji"
	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
	"regexp"
)

var (
	exceptionHandler = &handler.ExceptionHandler{}
	iroiroHandler    = &handler.IroiroHandler{}
	colorsHandler    = &handler.ColorsHandler{}
)

func main() {
	// Database
	db.Connect()

	// Goji
	m := web.New()

	// Iroiro
	m.Get("/v1/iroiro", iroiroHandler.Iroiro)
	m.Post("/v1/iroiro", iroiroHandler.CreateIro)
	m.Get(regexp.MustCompile(`^/v1/iroiro/(?P<id>\d+)$`), iroiroHandler.Iro)
	m.Put(regexp.MustCompile(`^/v1/iroiro/(?P<id>\d+)$`), iroiroHandler.UpdateIro)

	// Colors
	m.Get("/v1/colors", colorsHandler.Colors)
	m.Post("/v1/colors", colorsHandler.CreateColor)
	m.Get(regexp.MustCompile(`^/v1/colors/(?P<id>\d+)$`), colorsHandler.Color)

	// Exception
	m.NotFound(exceptionHandler.NotFound)

	// Serve
	graceful.Serve(bind.Default(), m)
}
