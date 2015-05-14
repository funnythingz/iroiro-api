package main

import (
	"./db"
	"./handler"
	_ "github.com/zenazn/goji"
	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
	"regexp"
)

var (
	iroiroHandler    = &handler.IroiroHandler{}
	colorsHandler    = &handler.ColorsHandler{}
	exceptionHandler = &handler.ExceptionHandler{}
)

func main() {
	// Database
	db.Connect()

	// Goji
	m := web.New()

	// Iroiro
	m.Get("/v1/iroiro", iroiroHandler.Iroiro)
	m.Post("/v1/iroiro", iroiroHandler.Create)
	m.Get(regexp.MustCompile(`^/v1/iroiro/(?P<id>\d+)$`), iroiroHandler.Iro)

	// Colors
	m.Get("/v1/colors", colorsHandler.Colors)
	m.Post("/v1/colors", colorsHandler.Create)
	m.Get(regexp.MustCompile(`^/v1/colors/(?P<id>\d+)$`), colorsHandler.Color)

	// Exception
	m.NotFound(exceptionHandler.NotFound)

	// Serve
	graceful.Serve(bind.Default(), m)
}
