package helpers

import (
	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("AllYourBase")

type MyCustomClaims struct {
	ID uint
	jwt.StandardClaims
}

func GenerateJWT(email string, userId uint) (string, error) {
	claims := MyCustomClaims{
		ID: userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func ParseToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims.ID, nil
	} else {
		return 0, err
	}
}
