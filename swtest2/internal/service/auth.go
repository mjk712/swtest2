package service

import (
	"swtest2/internal/config"
	"swtest2/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func (s *service) GetUserToken(login, password string) (string, error) {

	loginPassword := login + "/" + password
	colleague := &models.Colleague{}
	client := &models.Client{}
	err := s.travelAgencyRepo.AuthColleague(colleague, loginPassword)
	if err == nil {
		role := "colleague"
		id := colleague.Id
		token, err := generateJWT(id, role)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	err = s.travelAgencyRepo.AuthClient(client, loginPassword)
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
	cfg := config.NewStorageConfig()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(cfg.TokenSecretKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
