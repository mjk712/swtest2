package service

import (
	"errors"
	"strings"
	"swtest2/internal/models"
)

func (s *service) AddClient(client *models.Client) error {
	logPass := client.LoginPassword
	authData := strings.Split(logPass, "/")
	for _, v := range authData {
		if strings.Contains(v, "/") {
			return errors.New("found fobbiden symbol '/' in the loginPassword data")
		}
	}
	err := s.travelAgencyRepo.CreateClient(client)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ShowClientInfo(id string) ([]*models.ClientVacations, error) {
	client, err := s.travelAgencyRepo.ShowClientInfo(id)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *service) ChangeClient(client *models.Client, id string) error {

	if client.Fio != "" {
		err := s.travelAgencyRepo.ChangeClientFio(client, id)
		if err != nil {
			return err
		}
	}
	if client.Passport != "" {
		err := s.travelAgencyRepo.ChangeClientPassport(client, id)
		if err != nil {
			return err
		}
	}
	if client.EmailTelephone != "" {
		err := s.travelAgencyRepo.ChangeClientEmailTelephone(client, id)
		if err != nil {
			return err
		}
	}
	if client.LoginPassword != "" {
		logPass := client.LoginPassword
		authData := strings.Split(logPass, "/")
		for _, v := range authData {
			if strings.Contains(v, "/") {
				return errors.New("found fobbiden symbol '/' in the loginPassword data")
			}
		}
		err := s.travelAgencyRepo.ChangeClientLoginPassword(client, id)
		if err != nil {
			return err
		}
	}
	err := s.travelAgencyRepo.GetClient(client, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ShowClientsApplicationsInfo() ([]*models.ClientsApplications, error) {
	clientApplications, err := s.travelAgencyRepo.ShowClientApplicationsInfo()
	if err != nil {
		return nil, err
	}
	return clientApplications, nil
}
