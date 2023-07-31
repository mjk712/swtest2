package database

import "swtest2/internal/models"

type Storage interface {
	CreateApplication(application *models.Application, clientId string) error
	ChangeApplicationStatus(id, status string) error

	AuthClient(client *models.Client, loginpassword string) error
	CreateClient(client *models.Client) error
	ShowClientInfo(id string) ([]*models.ClientVacations, error)
	ChangeClientFio(client *models.Client, id string) error
	ChangeClientPassport(client *models.Client, id string) error
	ChangeClientEmailTelephone(client *models.Client, id string) error
	ChangeClientLoginPassword(client *models.Client, id string) error
	GetClient(client *models.Client, id string) error
	ShowClientApplicationsInfo() ([]*models.ClientsApplications, error)

	AuthColleague(colleague *models.Colleague, loginpassword string) error
	CreateColleague(colleague *models.Colleague) error
	BlockColleague(id string) error

	CreateCompany(company *models.Company) (*models.Company, error)
	AddClientIntoCompany(clientFio, companyName string) error
	ChangeCompanyName(company *models.Company, id string) error
	ChangeCompanyInn(company *models.Company, id string) error
	ChangeCompanyLegalAddress(company *models.Company, id string) error
	ChangeCompanyClients(company *models.Company, id string) error
	GetCompany(company *models.Company, id string) error
	ShowClientCompanysInfo(id string) ([]*models.ClientCompanys, error)
}
