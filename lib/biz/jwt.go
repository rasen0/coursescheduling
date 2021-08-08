package biz

import (
	"errors"
	"fmt"
	"time"

	//"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
)

func MakeToken(userName string) (tokenString string, err error){
	var hmacSampleSecret []byte
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userName,
		"iss": "schedulingcourse",
		"exp": time.Now().Add(3* time.Hour).Unix(),
	})
	tokenString, err = token.SignedString(hmacSampleSecret)
	return
}

func ParseToken(tokenString string) (claims jwt.MapClaims,err error) {
	if tokenString == ""{
		return nil, errors.New("token is nil")
	}
	var hmacSampleSecret []byte
	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	claims ,_ = token.Claims.(jwt.MapClaims)
	return claims, err
}