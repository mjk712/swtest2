package database

import (
	"fmt"
	"swtest2/internal/database/query"
	"swtest2/internal/models"
)

func CreateCompany(company *models.Company) (*models.Company, error) {
	_, err := db.NamedExec(query.InsertCompany, company)
	if err != nil {
		return nil, err
	}
	companyId := &models.CompanyId{}
	clientId := &models.ClientId{}
	err = db.Get(companyId, query.GetCompanyIdByName, company.Name)
	if err != nil {
		return nil, err
	}
	for _, v := range company.Clients {
		err = db.Get(clientId, query.GetClientIdByFio, v)
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
	err := db.Get(companyId, query.GetCompanyIdByName, companyName)
	if err != nil {
		return err
	}

	err = db.Get(clientId, query.GetClientIdByFio, clientFio)
	if err != nil {
		return err
	}
	_, err = db.Query(query.InsertCompanyClients, int(companyId.Id), int(clientId.Id))
	if err != nil {
		return err
	}
	return nil
}

func ChangeCompanyName(company *models.Company, id string) error {
	_, err := db.Query(query.ChangeCompanyName, company.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func ChangeCompanyInn(company *models.Company, id string) error {
	_, err := db.Query(query.ChangeCompanyInn, company.Inn, id)
	if err != nil {
		return err
	}
	return nil
}

func ChangeCompanyLegalAddress(company *models.Company, id string) error {
	_, err := db.Query(query.ChangeCompanyLegalAddress, company.Legal_address, id)
	if err != nil {
		return err
	}
	return nil
}

func ChangeCompanyClients(company *models.Company, id string) error {
	clientId := &models.ClientId{}

	_, err := db.Query(query.DeleteCompanyClients, id)
	if err != nil {
		return err
	}

	for _, v := range company.Clients {
		err := db.Get(clientId, query.GetClientIdByFio, v)
		if err != nil {
			fmt.Println(err)
		}
		_, err = db.Query(query.InsertCompanyClients, id, int(clientId.Id))
		if err != nil {
			return err
		}
	}
	return nil
}

func GetCompany(company *models.Company, id string) error {
	err := db.Get(company, query.GetCompanyById, id)
	if err != nil {
		return err
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
