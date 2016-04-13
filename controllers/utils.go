package controllers

import (
	"encoding/base64"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

// DecodeToken decode token
func DecodeToken(myToken string) (string, error) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("privateKey")), nil
	})
	if err == nil && token.Valid {
		return token.Claims["userid"].(string), nil
	}
	return "Invalid token", err
}

//EncodeToken create token
func EncodeToken(userID, pass string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["userid"] = userID
	token.Claims["pass"] = pass
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(beego.AppConfig.String("privateKey"))
	return tokenString, err
}

//EncodeID64 create id for a user
func EncodeID64(email, name, surname string) string {
	msg := email + name + surname
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	return encoded
}
