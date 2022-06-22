package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}
