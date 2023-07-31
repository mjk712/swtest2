package database

import (
	"fmt"
	"swtest2/internal/database/query"
	"swtest2/internal/models"
)

func (r RepoDB) CreateCompany(company *models.Company) (*models.Company, error) {
	_, err := r.db.NamedExec(query.InsertCompany, company)
	if err != nil {
		return nil, err
	}
	companyId := &models.CompanyId{}
	clientId := &models.ClientId{}
	err = r.db.Get(companyId, query.GetCompanyIdByName, company.Name)
	if err != nil {
		return nil, err
	}
	for _, v := range company.Clients {
		err = r.db.Get(clientId, query.GetClientIdByFio, v)
		if err != nil {
			fmt.Println(err)
		}
		company.Id = companyId.Id
		_, err := r.db.Query(query.InsertCompanyClients, int(companyId.Id), int(clientId.Id))
		if err != nil {
			return nil, err
		}
	}
	return company, nil
}

func (r RepoDB) AddClientIntoCompany(clientFio, companyName string) error {
	companyId := &models.CompanyId{}
	clientId := &models.ClientId{}
	err := r.db.Get(companyId, query.GetCompanyIdByName, companyName)
	if err != nil {
		return err
	}

	err = r.db.Get(clientId, query.GetClientIdByFio, clientFio)
	if err != nil {
		return err
	}
	_, err = r.db.Query(query.InsertCompanyClients, int(companyId.Id), int(clientId.Id))
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) ChangeCompanyName(company *models.Company, id string) error {
	_, err := r.db.Query(query.ChangeCompanyName, company.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) ChangeCompanyInn(company *models.Company, id string) error {
	_, err := r.db.Query(query.ChangeCompanyInn, company.Inn, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) ChangeCompanyLegalAddress(company *models.Company, id string) error {
	_, err := r.db.Query(query.ChangeCompanyLegalAddress, company.LegalAddress, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) ChangeCompanyClients(company *models.Company, id string) error {
	clientId := &models.ClientId{}

	_, err := r.db.Query(query.DeleteCompanyClients, id)
	if err != nil {
		return err
	}

	for _, v := range company.Clients {
		err := r.db.Get(clientId, query.GetClientIdByFio, v)
		if err != nil {
			fmt.Println(err)
		}
		_, err = r.db.Query(query.InsertCompanyClients, id, int(clientId.Id))
		if err != nil {
			return err
		}
	}
	return nil
}

func (r RepoDB) GetCompany(company *models.Company, id string) error {
	err := r.db.Get(company, query.GetCompanyById, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) ShowClientCompanysInfo(id string) ([]*models.ClientCompanys, error) {
	var clientCompanysInfo = make([]*models.ClientCompanys, 0)
	rows, err := r.db.Queryx(query.ShowClientCompanysInfo, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a models.ClientCompanys
		err = rows.StructScan(&a)
		if err != nil {
			return nil, err
		}
		clientCompanysInfo = append(clientCompanysInfo, &a)
	}
	rows.Close()
	return clientCompanysInfo, nil
}
