package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/RipulHandoo/blogx/authentication/db/database"
	"github.com/RipulHandoo/blogx/authentication/pkg"
	"github.com/RipulHandoo/blogx/authentication/pkg/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// @title Register a user
// @version 1
// @description Register a user with fist name, last name, email, password and bio given in the body
// @Tags authentication
// @accept json
// @param data body database.User true "User details"
// @produce json
// @success 201 {object} utils.DbUserFullSchema
// @failure 400 {object} string
// @failure 500 {object} string
// @router /auth/register [post]
func HandleRegisterUser(w http.ResponseWriter, req *http.Request) {
	type reqBody struct {
		FirstName string         `json:"first_name"`
		LastName  string         `json:"last_name"`
		Email     string         `json:"email"`
		Password  string         `json:"password"`
		Bio       sql.NullString `json:"bio"`
	}
	decoder := json.NewDecoder(req.Body)

	bodyDecoded := reqBody{}

	if err := decoder.Decode(&bodyDecoded); err != nil {
		utils.ResponseJson(w, 400, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	apiConfig := pkg.DbInstance()
	// fmt.Println(&apiConfig)

	saltValueString := os.Getenv("BCRYPT_SALT_VALUE")

	saltValue, bcryptErr := strconv.Atoi(saltValueString)

	if bcryptErr != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, bcryptErr)
		return
	}
	
	hashedPassword, err2 := bcrypt.GenerateFromPassword([]byte(bodyDecoded.Password), saltValue)
	if err2 != nil {
		hashedPassword = []byte(bodyDecoded.Password)
	}

	user, failedToAddToDb := apiConfig.CreateUser(
		req.Context(),
		database.CreateUserParams{
			ID:        uuid.New(),
			FirstName: bodyDecoded.FirstName,
			LastName:  bodyDecoded.LastName,
			Email:     bodyDecoded.Email,
			Password:  string(hashedPassword),
			Bio:       bodyDecoded.Bio,
		},
	)

	if failedToAddToDb != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, failedToAddToDb)
		return
	}

	// create jwt token
	token, expiryTime, jwtTokenError := utils.GetJwt(utils.Credentials{
		Email: bodyDecoded.Email,
		Name:  user.FirstName + user.LastName,
	})

	if jwtTokenError != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, jwtTokenError)
		return
	}

	expiry := expiryTime

	http.SetCookie(w, &http.Cookie{
		Name:    "auth_token",
		Value:   token,
		Expires: expiry,
		Path:    "/",
	})

	utils.ResponseJson(w, http.StatusCreated, utils.MapRegisteredUser(user))
}