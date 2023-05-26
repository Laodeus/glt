package tokenUtils

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(reqToken string) (*jwt.Token, error) {

	var jwtKey = []byte(os.Getenv("SECRET"))

	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token method conform to "SigningMethodHMAC"
		_, isSigningMethodValid := token.Method.(*jwt.SigningMethodHMAC)
		if !isSigningMethodValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	return token, err
}
