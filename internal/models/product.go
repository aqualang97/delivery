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
	ExternalID     int      `json:"id"` // надо external id, но в фуд апи это id. в итоге можем ловить его только так
	Name           string   `json:"name"`
	Price          float64  `json:"price"`
	Image          string   `json:"image"`
	Type           string   `json:"type"`
	SupplierId     int      `json:"supplierId"`
	SupplierName   string   `json:"supplierName"`
	ExternalSuppId int      `json:"externalSuppId"`
	Ingredients    []string `json:"ingredients"`
	CategoryNum    int      `json:"categoryNum"`
}
