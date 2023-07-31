package service

import (
	"swtest2/internal/database"
	"swtest2/internal/models"
)

func AddClient(client *models.Client) error {
	err := database.CreateClient(client)
	if err != nil {
		return err
	}
	return nil
}

func ShowClientInfo(id string) ([]*models.Client_Vacations, error) {
	client, err := database.ShowClientInfo(id)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ChangeClient(client *models.Client, id string) error {

	if client.Fio != "" {
		err := database.ChangeClientFio(client, id)
		if err != nil {
			return err
		}
		return nil
	}
	if client.Passport != "" {
		err := database.ChangeClientPassport(client, id)
		if err != nil {
			return err
		}
		return nil
	}
	if client.Email_Telephone != "" {
		err := database.ChangeClientEmailTelephone(client, id)
		if err != nil {
			return err
		}
		return nil
	}
	if client.LoginPassword != "" {
		err := database.ChangeClientLoginPassword(client, id)
		if err != nil {
			return err
		}
		return nil
	}
	err := database.GetClient(client, id)
	if err != nil {
		return err
	}
	return nil
}

func ShowClientsApplicationsInfo() ([]*models.Clients_Applications, error) {
	clientApplications, err := database.ShowClientApplicationsInfo()
	if err != nil {
		return nil, err
	}
	return clientApplications, nil
}
