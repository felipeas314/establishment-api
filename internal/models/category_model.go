package models

type Category struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	EstablishmentID int    `json:"establishment_id"`
}
