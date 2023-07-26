package models

type Application struct {
	Id             uint   `json:"id"`
	VacationId     uint   `json:"vacationid"`
	Status         string `json:"status"`
	Client_message string `json:"client_message"`
}
