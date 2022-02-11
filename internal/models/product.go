package models

import "time"

type Product struct {
	Id         int
	Name       string
	Category   int
	ExternalID int
	//Description string
	//Discount   int
	CreatedAt *time.Time
}

type Position struct {
	ExternalID  int      `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
}
