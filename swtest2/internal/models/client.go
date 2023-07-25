package models

type Client struct {
	Id              uint   `json:"id"`
	Fio             string `json:"fio"`
	Passport        string `json:"passport"`
	Email_Telephone string `json:"email_telephone"`
	LoginPassword   string `json:"loginpassword"`
}

type ClientId struct {
	Id uint `json:"id"`
}
