package server

import (
	"fmt"
	"net/http"

	"github.com/RipulHandoo/blogx/authentication/pkg"
	"github.com/RipulHandoo/blogx/authentication/pkg/utils"
)

type error struct {
	Error string `json:"error"`	
}

type resp struct{
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// check if the database is setup properly or not
	db := pkg.DbClient

	if db == nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("Database not setup properly"))
		return
	}

	utils.ResponseJson(w,http.StatusAccepted, resp{
		Status: "ok",
	})
}