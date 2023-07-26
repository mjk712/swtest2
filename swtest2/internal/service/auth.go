package service

import (
	"fmt"
	"swtest2/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
)

var MySecretKey = []byte("GregorySwtest2")

func GetUserToken(login, password string) (string, error) {

	loginpassword := login + "/" + password
	colleague := &models.Colleague{}
	client := &models.Client{}
	err := db.Get(colleague, "SELECT id FROM colleague WHERE loginpassword = $1", loginpassword)
	if err == nil {
		role := "colleague"
		id := colleague.Id
		token, err := generateJWT(id, role)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	err = db.Get(client, "SELECT id FROM client WHERE loginpassword = $1", loginpassword)
	if err == nil {
		role := "client"
		id := client.Id
		token, err := generateJWT(id, role)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", err
}

func generateJWT(id uint, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(MySecretKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}