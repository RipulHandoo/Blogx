package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Credential struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Claims struct {
	Creds Credential
	jwt.RegisteredClaims
}


func GetJwt(signedCred Credential) (string, int64, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET")

	experyTime := time.Now().Add(time.Hour * 24 * 7).Unix()

	claims := Claims {
		Creds: signedCred,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(experyTime, 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", 0, err
	}

	return tokenString, experyTime, nil
}