package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RipulHandoo/blogx/user/db/database"
	"github.com/RipulHandoo/blogx/user/pkg"
	"github.com/RipulHandoo/blogx/user/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func Auth(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		godotenv.Load()
		var jwtKey string = os.Getenv("JWT_SECRET_KEY")
		jwtToken := req.Header.Get("auth_token")
		if jwtToken == "" {
			utils.ResponseWithError(w, http.StatusUnauthorized, "no auth token")
			fmt.Println("No auth token")
			return
		}
		tknStr := jwtToken
		claims := &utils.Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			fmt.Println("No valid jwt - " + tknStr)
			utils.ResponseWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		if !tkn.Valid {
			utils.ResponseWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		userEmail := claims.Creds.Email
		apiConfig := pkg.DbClient

		user, dbErr2 := apiConfig.GetUserByEmail(req.Context(), userEmail)
		if dbErr2 != nil {
			utils.ResponseWithError(w, http.StatusInternalServerError, dbErr2.Error())
			return
		}

		handler(w, req, user)
	}
}