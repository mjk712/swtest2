package models

type ClientsApplications struct {
	ClientFullname    string `json:"client_full_name"`
	ApplicationStatus string `json:"application_status"`
	ApplicationId     string `json:"application_id"`
	ClientMessage     string `json:"client_message"`
}
