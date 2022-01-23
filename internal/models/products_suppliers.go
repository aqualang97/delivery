package models

import "time"

type ProductsSuppliers struct {
	ProductID          int
	SupplierID         int
	ExternalProductID  int
	ExternalSupplierID int
	Price              int
	Image              string
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}
