package service

import (
	"swtest2/internal/models"
)

func (s *service) AddColleague(colleague *models.Colleague) error {
	err := s.travelAgencyRepo.CreateColleague(colleague)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) BlockColleague(id string) error {
	err := s.travelAgencyRepo.BlockColleague(id)
	if err != nil {
		return err
	}
	return nil
}
