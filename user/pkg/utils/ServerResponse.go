package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func ResponseWithError(w http.ResponseWriter, code int, message string) {
	ResponseJSON(w, code, struct {
		Error string `json:"error"`
	}{
		Error: message,
	})
}
