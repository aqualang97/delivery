package models

import "time"

type Product struct {
	ID         int    `json:"ID"`
	Name       string `json:"Name"`
	Category   int    `json:"Category"`
	ExternalID int    `json:"ExternalID"`
	//Description string
	//Discount   int
	CreatedAt *time.Time
}

type Position struct {
	ID          int      `json:"ID"`
	ExternalID  int      `json:"external_id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
}
