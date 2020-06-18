package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	secret = "your secret"
)

func SignJwtToken(id string) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(60 * time.Second).Unix(),
		Subject: id,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidJwtToken(token string) (string,bool) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println(err)
		return "",false
	}
	if err = jwtToken.Claims.Valid(); err != nil {
		fmt.Println(err)
		return "",false
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		fmt.Println("token 过期")
		return "",false
	}
	return claims["sub"].(string), true
}