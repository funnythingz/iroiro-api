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
	goji.Get("/v1/iroiro", iroiroController.iroiro)
	goji.Post("/v1/iroiro", iroiroController.create)
	goji.Get(regexp.MustCompile(`^/v1/iroiro/(?P<id>\d+)$`), iroiroController.iro)

	// Exception
	goji.NotFound(exceptionController.NotFound)

	// Serve
	goji.Serve()
}
