package dto

type EventDTO struct {
	Description string `json:"description"`
	Address     string `json:"address"`
	Date        string `json:"date"`
	UserID      string `json:"userId"`
}

type EventUpdateDTO struct {
	Description string `json:"description"`
	Address     string `json:"address"`
	Date        string `json:"date"`
}
