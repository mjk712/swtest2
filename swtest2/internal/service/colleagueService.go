package service

import (
	"errors"
	"strings"
	"swtest2/internal/models"
)

func (s *service) AddColleague(colleague *models.Colleague) error {
	logPass := colleague.LoginPassword
	authData := strings.Split(logPass, "/")
	for _, v := range authData {
		if strings.Contains(v, "/") {
			return errors.New("found fobbiden symbol '/' in the loginPassword data")
		}
	}
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
