package service

import (
	"errors"
	"regexp"
	"swtest2/internal/database"
	"swtest2/internal/models"
)

func ChangeApplicationStatus(id, status string) error {
	matchedStatus, _ := regexp.MatchString(`Новая|В работе|Согласована|Оплачена`, status)
	if !matchedStatus {
		return errors.New("invalid status")
	}
	err := database.ChangeApplicationStatus(id, status)
	if err != nil {
		return err
	}
	return nil
}

func AddApplication(application *models.Application, clientId string) error {
	err := database.CreateApplication(application, clientId)
	if err != nil {
		return err
	}
	return nil
}
