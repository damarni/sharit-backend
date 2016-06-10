package controllers

import (
	"encoding/base64"
	"fmt"
	"hash/fnv"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

func DecodeToken(myToken string) (string, error) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("privateKey")), nil
	})
	if err == nil && token.Valid {
		fmt.Println("token valid")
		return token.Claims["userid"].(string), nil
	}
	fmt.Println(err.Error())

	return "Invalid token", err
}

func EncodeToken(userID, pass string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["userid"] = userID

	fmt.Println(beego.AppConfig.String("privateKey"))
	key := []byte(beego.AppConfig.String("privateKey"))
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println("token")
		fmt.Println(err.Error())

	} else {
		fmt.Println("token ok")

	}
	fmt.Println(tokenString)
	return tokenString, err
}

func EncodeID64(email, name, surname string) string {
	msg := email + name + surname
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	return encoded
}

func EncodeMsg(msg string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	return encoded
}

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
