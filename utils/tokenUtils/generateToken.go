package tokenUtils

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId int) (string, error) {

	var jwtKey = []byte(os.Getenv("SECRET"))

	claims := jwt.MapClaims{
		"id": userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)

}
