package service

import (
	"swtest2/internal/database/query"
	"swtest2/internal/models"
)

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
	err := db.Get(client, "SELECT * FROM client WHERE id = $1", id)
	if err != nil {
		return err
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
