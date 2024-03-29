package auth

import (
	"net/http"

	"github.com/RipulHandoo/blogx/authentication/db/database"
	"github.com/RipulHandoo/blogx/authentication/pkg/utils"
)

// @title Logout a user
// @version 1
// @description Logout a user
// @Tags authentication
// @accept json
// @produce json
// @success 202 {object} utils.DBUserResponse
// @failure 400 {object} string
// @failure 500 {object} string
// @router /auth/logout [post]
func HandleUserLogout(w http.ResponseWriter, _ *http.Request, user database.User) {
	//clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "auth_token",
		Value: "",
		Path:  "/",
	})
	utils.ResponseJson(w, http.StatusAccepted, utils.MapLoginUser(user))
}