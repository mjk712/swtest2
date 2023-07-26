package service

import (
	"fmt"
	"swtest2/internal/database/query"
	"swtest2/internal/models"
)

func AddColleague(colleague *models.Colleague) error {
	_, err := db.NamedExec(query.AddColleague, colleague)
	if err != nil {
		return err
	}
	return nil
}

func BlockColleague(id string) error {

	_, err := db.Query(query.BlockColleague, id)
	if err != nil {
		return err
	}
	return nil
}

func AddClient(client *models.Client) error {
	_, err := db.NamedExec(query.AddClient, client)
	if err != nil {
		return err
	}
	return nil
}

func ShowClientInfo(id string) ([]*models.Client_Vacations, error) {

	var client = make([]*models.Client_Vacations, 0)
	rows, err := db.Queryx(query.ShowClientInfo, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a models.Client_Vacations
		err = rows.StructScan(&a)
		if err != nil {
			return nil, err
		}
		client = append(client, &a)
	}
	rows.Close()

	return client, nil
}

func ChangeClient(client *models.Client, id string) error {

	if client.Fio != "" {
		_, err := db.Query(query.ChangeClientFio, client.Fio, id)
		if err != nil {
			return err
		}
	}
	if client.Passport != "" {
		_, err := db.Query(query.ChangeClientPassport, client.Passport, id)
		if err != nil {
			return err
		}
	}
	if client.Email_Telephone != "" {
		_, err := db.Query(query.ChangeClientEmailTelephone, client.Email_Telephone, id)
		if err != nil {
			return err
		}
	}
	if client.LoginPassword != "" {
		_, err := db.Query(query.ChangeClientLoginPassword, client.LoginPassword, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func ShowClientsApplicationsInfo() ([]*models.Clients_Applications, error) {

	var client = make([]*models.Clients_Applications, 0)
	rows, err := db.Queryx(query.ShowClientsApplication)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a models.Clients_Applications
		err = rows.StructScan(&a)
		if err != nil {
			return nil, err
		}
		client = append(client, &a)
	}
	rows.Close()

	return client, nil
}

func ChangeApplicationStatus(id, status string) error {

	_, err := db.Query(query.ChangeApplicationStatus, status, id)
	if err != nil {
		return err
	}
	return nil
}

func InsertCompany(company *models.Company) (*models.Company, error) {

	_, err := db.NamedExec(query.InsertCompany, company)
	if err != nil {
		return nil, err
	}
	companyId := &models.CompanyId{}
	clientId := &models.ClientId{}
	err = db.Get(companyId, "SELECT id FROM company WHERE name = $1", company.Name)
	if err != nil {
		return nil, err
	}
	for _, v := range company.Clients {
		err = db.Get(clientId, "SELECT id FROM client WHERE fio = $1", v)
		if err != nil {
			fmt.Println(err)
		}
		company.Id = companyId.Id
		_, err := db.Query(query.InsertCompanyClients, int(companyId.Id), int(clientId.Id))
		if err != nil {
			return nil, err
		}
	}
	return company, nil
}

func AddClientIntoCompany(clientFio, companyName string) error {

	companyId := &models.CompanyId{}
	clientId := &models.ClientId{}
	err := db.Get(companyId, "SELECT id FROM company WHERE name = $1", companyName)
	if err != nil {
		return err
	}

	err = db.Get(clientId, "SELECT id FROM client WHERE fio = $1", clientFio)
	if err != nil {
		return err
	}
	_, err = db.Query(query.InsertCompanyClients, int(companyId.Id), int(clientId.Id))
	if err != nil {
		return err
	}
	return nil
}

func ChangeCompany(company *models.Company, id string) error {

	if company.Name != "" {
		_, err := db.Query(query.ChangeCompanyName, company.Name, id)
		if err != nil {
			return err
		}
	}
	if company.Inn != "" {
		_, err := db.Query(query.ChangeCompanyInn, company.Inn, id)
		if err != nil {
			return err
		}
	}
	if company.Legal_address != "" {
		_, err := db.Query(query.ChangeCompanyLegalAddress, company.Legal_address, id)
		if err != nil {
			return err
		}
	}
	if company.Clients != nil {
		clientId := &models.ClientId{}

		_, err := db.Query(query.DeleteCompanyClients, id)
		if err != nil {
			return err
		}

		for _, v := range company.Clients {
			err := db.Get(clientId, "SELECT id FROM client WHERE fio = $1", v)
			if err != nil {
				fmt.Println(err)
			}
			_, err = db.Query(query.InsertCompanyClients, id, int(clientId.Id))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ShowClientCompanysInfo(id string) ([]*models.Client_Companys, error) {

	var clientCompanysInfo = make([]*models.Client_Companys, 0)
	rows, err := db.Queryx(query.ShowClientCompanysInfo, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a models.Client_Companys
		err = rows.StructScan(&a)
		if err != nil {
			return nil, err
		}
		clientCompanysInfo = append(clientCompanysInfo, &a)
	}
	rows.Close()

	return clientCompanysInfo, nil
}

func AddApplication(application *models.Application) error {
	_, err := db.NamedExec(query.AddApplication, application)
	if err != nil {
		return err
	}
	return nil
}
