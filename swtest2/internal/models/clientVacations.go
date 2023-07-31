package models

type ClientVacations struct {
	ClientFullname       string `json:"client_full_name"`
	ClientPassport       string `json:"client_passport"`
	ClientEmailTelephone string `json:"client_email_telephone"`
	VacationName         string `json:"vacation_name"`
}
