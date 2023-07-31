package service

import (
	"swtest2/internal/models"
)

func (s *service) AddClient(client *models.Client) error {
	err := s.travelAgencyRepo.CreateClient(client)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ShowClientInfo(id string) ([]*models.Client_Vacations, error) {
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
		return nil
	}
	if client.Passport != "" {
		err := s.travelAgencyRepo.ChangeClientPassport(client, id)
		if err != nil {
			return err
		}
		return nil
	}
	if client.Email_Telephone != "" {
		err := s.travelAgencyRepo.ChangeClientEmailTelephone(client, id)
		if err != nil {
			return err
		}
		return nil
	}
	if client.LoginPassword != "" {
		err := s.travelAgencyRepo.ChangeClientLoginPassword(client, id)
		if err != nil {
			return err
		}
		return nil
	}
	err := s.travelAgencyRepo.GetClient(client, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ShowClientsApplicationsInfo() ([]*models.Clients_Applications, error) {
	clientApplications, err := s.travelAgencyRepo.ShowClientApplicationsInfo()
	if err != nil {
		return nil, err
	}
	return clientApplications, nil
}
