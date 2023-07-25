package models

type Company struct {
	Id            uint     `json:"id"`
	Name          string   `json:"name"`
	Inn           string   `json:"inn"`
	Legal_address string   `json:"legal_address"`
	Clients       []string `json:"clients"`
}

type CompanyId struct {
	Id uint `json:"id"`
}
