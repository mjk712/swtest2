package models

type Application struct {
	Id            uint   `json:"id"`
	VacationId    uint   `json:"vacationid"`
	Status        string `json:"status"`
	ClientMessage string `json:"client_message"`
}
