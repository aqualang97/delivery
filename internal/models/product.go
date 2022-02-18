package models

import "time"

type Product struct {
	ID         int    `json:"ID"`
	Name       string `json:"name"`
	Category   int    `json:"category"`
	ExternalID int    `json:"externalID"`
	//Description string
	//Discount   int
	CreatedAt *time.Time `json:"createdAt"`
}

type Position struct {
	ID             int      `json:"ID"`
	ExternalID     int      `json:"externalID"`
	Name           string   `json:"name"`
	Price          float64  `json:"price"`
	Image          string   `json:"image"`
	Type           string   `json:"type"`
	SupplierId     int      `json:"supplierId"`
	ExternalSuppId int      `json:"externalSuppId"`
	Ingredients    []string `json:"ingredients"`
}
