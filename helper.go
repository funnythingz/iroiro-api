package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func resultJSON(w http.ResponseWriter, m string) {
	response, _ := json.Marshal(map[string]string{"Message": m})
	io.WriteString(w, string(response))
}
