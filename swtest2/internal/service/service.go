package service

import (
	"swtest2/internal/database"
	"swtest2/internal/models"
)

type service struct {
	travelAgencyRepo database.Storage
}

type Service interface {
	GetUserToken(login, password string) (string, error)

	ChangeApplicationStatus(id, status string) error
	AddApplication(application *models.Application, clientId string) error

	AddClient(client *models.Client) error
	ShowClientInfo(id string) ([]*models.Client_Vacations, error)
	ChangeClient(client *models.Client, id string) error
	ShowClientsApplicationsInfo() ([]*models.Clients_Applications, error)

	AddColleague(colleague *models.Colleague) error
	BlockColleague(id string) error

	InsertCompany(company *models.Company) (*models.Company, error)
	AddClientIntoCompany(clientFio, companyName string) error
	ChangeCompany(company *models.Company, id string) error
	ShowClientCompanysInfo(id string) ([]*models.Client_Companys, error)
}

func NewService(travelAgencyRepo database.Storage) Service {
	return &service{travelAgencyRepo}
}
