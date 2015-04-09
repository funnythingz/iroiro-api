package main

import (
	"net/http"
)

type ExceptionController struct{}

func (_ *ExceptionController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
