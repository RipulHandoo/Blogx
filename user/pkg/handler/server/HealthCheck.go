package server

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "Server is running good"}`))
}
