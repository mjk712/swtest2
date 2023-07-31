package service

import (
	"swtest2/internal/models"
)

func (s *service) InsertCompany(company *models.Company) (*models.Company, error) {
	filledCompany, err := s.travelAgencyRepo.CreateCompany(company)
	if err != nil {
		return nil, err
	}
	return filledCompany, nil
}

func (s *service) AddClientIntoCompany(clientFio, companyName string) error {
	err := s.travelAgencyRepo.AddClientIntoCompany(clientFio, companyName)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ChangeCompany(company *models.Company, id string) error {

	if company.Name != "" {
		err := s.travelAgencyRepo.ChangeCompanyName(company, id)
		if err != nil {
			return err
		}
	}
	if company.Inn != "" {
		err := s.travelAgencyRepo.ChangeCompanyInn(company, id)
		if err != nil {
			return err
		}
	}
	if company.Legal_address != "" {
		err := s.travelAgencyRepo.ChangeCompanyLegalAddress(company, id)
		if err != nil {
			return err
		}
	}
	if company.Clients != nil {
		err := s.travelAgencyRepo.ChangeCompanyClients(company, id)
		if err != nil {
			return err
		}
	}
	s.travelAgencyRepo.GetCompany(company, id)
	return nil
}

func (s *service) ShowClientCompanysInfo(id string) ([]*models.Client_Companys, error) {

	clientCompanysInfo, err := s.travelAgencyRepo.ShowClientCompanysInfo(id)
	if err != nil {
		return nil, err
	}

	return clientCompanysInfo, nil
}
