package auth

import (
	"encoding/json"
	"net/http"

	"github.com/RipulHandoo/blogx/authentication/pkg"
	"github.com/RipulHandoo/blogx/authentication/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func HandlerLoginuser(w http.ResponseWriter, req *http.Request) {
	type reqBody struct{
		Email string `json: "email"`
		Password string `json: "password"`
	}

	decoder := json.NewDecoder(req.Body)
	body := reqBody{}

	err := decoder.Decode(&body)

	if err != nil {
		utils.ErrorResponse(w,http.StatusBadRequest, err)
		return
	}

	// get the user from the database
	apiConfig := pkg.DbInstance()

	user, err := apiConfig.GetUserByEmail(req.Context(),body.Email)

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadGateway, err)
		return
	}

	// check if the password is right if the user exist in the database
	err = bcrypt.CompareHashAndPassword(([]byte(user.Password)), []byte(body.Password))

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	// generate the jwt token
	token, expiryTime, err := utils.GetJwt(
		utils.Credentials{
			Email: user.Email,
			Name: user.FirstName + " " + user.LastName,
		},
	)
	var expiry = expiryTime
	http.SetCookie(w, &http.Cookie{
		Name: "auth_token",
		Value: token,
		Expires: expiry,
		Path: "/",
	})

	utils.ResponseJson(w, http.StatusOK, utils.MapLoginUser(user))
}