package models

import "time"

type ProductsSuppliers struct {
	ProductID          int
	SupplierID         int
	ExternalProductID  int
	ExternalSupplierID int
	Price              float64
	Image              string
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}
type AllMenu struct {
	Menu []Position `json:"menu"`
}
