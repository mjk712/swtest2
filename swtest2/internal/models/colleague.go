package models

type Colleague struct {
	Id            uint   `json:"id"`
	Fio           string `json:"fio"`
	LoginPassword string `json:"loginpassword"`
	Status        string `json:"status"`
}
