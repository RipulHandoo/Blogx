package utils

import (
	"net/http"
	"encoding/json"
)
func ResponseJson(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func ErrorResponse(w http.ResponseWriter, code int, err error){
	ResponseJson(w, code, 
	struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}