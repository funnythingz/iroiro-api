package main

import (
	"github.com/zenazn/goji"
	"regexp"
)

var (
	iroiroController    = &IroiroController{}
	exceptionController = &ExceptionController{}
)

func main() {
	//TODO: MySQL
	//dbLoad()

	// Iroiro
	goji.Get("/iroiro", iroiroController.iroiro)
	goji.Get(regexp.MustCompile(`^/iroiro/(?P<id>\d+)$`), iroiroController.iro)
	goji.Post("/iroiro/create", iroiroController.createIro)

	// Exception
	goji.NotFound(exceptionController.NotFound)

	// Serve
	goji.Serve()
}
