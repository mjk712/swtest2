package service

import (
	"swtest2/internal/database"
	"swtest2/internal/models"
)

func AddColleague(colleague *models.Colleague) error {
	err := database.CreateColleague(colleague)
	if err != nil {
		return err
	}
	return nil
}

func BlockColleague(id string) error {
	err := database.BlockColleague(id)
	if err != nil {
		return err
	}
	return nil
}
