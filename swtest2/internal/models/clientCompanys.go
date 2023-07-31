package models

type ClientCompanys struct {
	ClientFullname      string `json:"client_full_name"`
	CompanyName         string `json:"company_name"`
	CompanyInn          string `json:"company_inn"`
	CompanyLegalAddress string `json:"company_legal_address"`
}
