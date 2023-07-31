package database

import (
	"swtest2/internal/database/query"
	"swtest2/internal/models"
)

func CreateClient(client *models.Client) error {
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

func ChangeClientFio(client *models.Client, id string) error {
	_, err := db.Query(query.ChangeClientFio, client.Fio, id)
	if err != nil {
		return err
	}
	return nil
}

func ChangeClientPassport(client *models.Client, id string) error {
	_, err := db.Query(query.ChangeClientPassport, client.Passport, id)
	if err != nil {
		return err
	}
	return nil
}

func ChangeClientEmailTelephone(client *models.Client, id string) error {
	_, err := db.Query(query.ChangeClientEmailTelephone, client.Email_Telephone, id)
	if err != nil {
		return err
	}
	return nil
}

func ChangeClientLoginPassword(client *models.Client, id string) error {
	_, err := db.Query(query.ChangeClientLoginPassword, client.LoginPassword, id)
	if err != nil {
		return err
	}
	return nil
}

func GetClient(client *models.Client, id string) error {
	err := db.Get(client, query.GetClientById, id)
	if err != nil {
		return err
	}
	return nil
}

func ShowClientApplicationsInfo() ([]*models.Clients_Applications, error) {
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
