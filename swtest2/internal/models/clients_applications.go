package models

type Clients_Applications struct {
	Client_full_name   string `json:"client_full_name"`
	Application_status string `json:"application_status"`
	Application_id     string `json:"application_id"`
}
