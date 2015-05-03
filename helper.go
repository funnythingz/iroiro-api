package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func resultJSON(w http.ResponseWriter, messages []string) {
	response, _ := json.Marshal(map[string][]string{"Message": messages})
	io.WriteString(w, string(response))
}
