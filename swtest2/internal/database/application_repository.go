package database

import (
	"swtest2/internal/database/query"
	"swtest2/internal/models"
)

func (r RepoDB) ChangeApplicationStatus(id, status string) error {
	_, err := r.db.Query(query.ChangeApplicationStatus, status, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) CreateApplication(application *models.Application, clientId string) error {
	_, err := r.db.NamedExec(query.AddApplication, application)
	if err != nil {
		return err
	}
	err = r.db.Get(application, query.GetClientByMessage, application.Client_message)
	if err != nil {
		return err
	}
	_, err = r.db.Query(query.InsertIntoClientApplication, clientId, application.Id)
	if err != nil {
		return err
	}
	_, err = r.db.Query(query.GetClientIdFromClientVacations, clientId, application.VacationId)
	if err != nil {
		_, err = r.db.Query(query.InsertIntoClientVacation, clientId, application.VacationId)
		if err != nil {
			return err
		}
	}
	return nil
}
