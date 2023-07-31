package database

import (
	"swtest2/internal/database/query"
	"swtest2/internal/models"
)

func (r RepoDB) AuthClient(client *models.Client, loginpassword string) error {
	err := r.db.Get(client, "SELECT id FROM client WHERE loginpassword = $1", loginpassword)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) CreateClient(client *models.Client) error {
	_, err := r.db.NamedExec(query.AddClient, client)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) ShowClientInfo(id string) ([]*models.ClientVacations, error) {
	var client = make([]*models.ClientVacations, 0)
	rows, err := r.db.Queryx(query.ShowClientInfo, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a models.ClientVacations
		err = rows.StructScan(&a)
		if err != nil {
			return nil, err
		}
		client = append(client, &a)
	}
	rows.Close()
	return client, nil
}

func (r RepoDB) ChangeClientFio(client *models.Client, id string) error {
	_, err := r.db.Query(query.ChangeClientFio, client.Fio, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) ChangeClientPassport(client *models.Client, id string) error {
	_, err := r.db.Query(query.ChangeClientPassport, client.Passport, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) ChangeClientEmailTelephone(client *models.Client, id string) error {
	_, err := r.db.Query(query.ChangeClientEmailTelephone, client.EmailTelephone, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) ChangeClientLoginPassword(client *models.Client, id string) error {
	_, err := r.db.Query(query.ChangeClientLoginPassword, client.LoginPassword, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) GetClient(client *models.Client, id string) error {
	err := r.db.Get(client, query.GetClientById, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) ShowClientApplicationsInfo() ([]*models.ClientsApplications, error) {
	var client = make([]*models.ClientsApplications, 0)
	rows, err := r.db.Queryx(query.ShowClientsApplication)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a models.ClientsApplications
		err = rows.StructScan(&a)
		if err != nil {
			return nil, err
		}
		client = append(client, &a)
	}
	rows.Close()
	return client, nil
}
