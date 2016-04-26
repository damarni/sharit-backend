package controllers

import (
	"encoding/base64"
	"fmt"
	"hash/fnv"
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
		fmt.Println("token valid")
		return token.Claims["userid"].(string), nil
	}
	fmt.Println(err.Error())

	return "Invalid token", err
}

//EncodeToken create token
func EncodeToken(userID, pass string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["userid"] = userID

	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
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

//EncodeID64 create id for a user
func EncodeID64(email, name, surname string) string {
	msg := email + name + surname
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	return encoded
}

// hash hash string
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
