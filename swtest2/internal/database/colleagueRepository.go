package database

import (
	"swtest2/internal/database/query"
	"swtest2/internal/models"
)

func (r RepoDB) CreateColleague(colleague *models.Colleague) error {
	_, err := r.db.Query(query.AddColleague, colleague.Fio, colleague.LoginPassword, colleague.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) BlockColleague(id string) error {
	_, err := r.db.Query(query.BlockColleague, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) AuthColleague(colleague *models.Colleague, loginpassword string) error {
	err := r.db.Get(colleague, "SELECT id FROM colleague WHERE loginpassword = $1", loginpassword)
	if err != nil {
		return err
	}
	return nil
}
