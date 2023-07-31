package service

import (
	"errors"
	"regexp"
	"swtest2/internal/models"
)

func (s *service) ChangeApplicationStatus(id, status string) error {
	matchedStatus, _ := regexp.MatchString(`Новая|В работе|Согласована|Оплачена`, status)
	if !matchedStatus {
		return errors.New("invalid status")
	}
	err := s.travelAgencyRepo.ChangeApplicationStatus(id, status)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) AddApplication(application *models.Application, clientId string) error {
	err := s.travelAgencyRepo.CreateApplication(application, clientId)
	if err != nil {
		return err
	}
	return nil
}
