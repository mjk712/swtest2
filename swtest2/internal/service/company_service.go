package service

import (
	"swtest2/internal/database"
	"swtest2/internal/models"
)

func InsertCompany(company *models.Company) (*models.Company, error) {
	filledCompany, err := database.CreateCompany(company)
	if err != nil {
		return nil, err
	}
	return filledCompany, nil
}

func AddClientIntoCompany(clientFio, companyName string) error {
	err := database.AddClientIntoCompany(clientFio, companyName)
	if err != nil {
		return err
	}
	return nil
}

func ChangeCompany(company *models.Company, id string) error {

	if company.Name != "" {
		err := database.ChangeCompanyName(company, id)
		if err != nil {
			return err
		}
	}
	if company.Inn != "" {
		err := database.ChangeCompanyInn(company, id)
		if err != nil {
			return err
		}
	}
	if company.Legal_address != "" {
		err := database.ChangeCompanyLegalAddress(company, id)
		if err != nil {
			return err
		}
	}
	if company.Clients != nil {
		err := database.ChangeCompanyClients(company, id)
		if err != nil {
			return err
		}
	}
	database.GetCompany(company, id)
	return nil
}

func ShowClientCompanysInfo(id string) ([]*models.Client_Companys, error) {

	clientCompanysInfo, err := database.ShowClientCompanysInfo(id)
	if err != nil {
		return nil, err
	}

	return clientCompanysInfo, nil
}
