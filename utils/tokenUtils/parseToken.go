package tokenUtils

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	jwt.StandardClaims
	ID int `json:"id"`
}

func ParseToken(reqToken string) (int, error) {

	var jwtKey = []byte(os.Getenv("SECRET"))

	token, err := jwt.ParseWithClaims(reqToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return 0, fmt.Errorf("Invalid token parsed")
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims.ID, nil
	} else {
		return 0, fmt.Errorf("Invalid token claims")
	}
}
