package database

import (
	"swtest2/internal/database/query"
	"swtest2/internal/models"
)

func CreateColleague(colleague *models.Colleague) error {
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
