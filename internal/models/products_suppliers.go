package models

import "time"

type ProductsSuppliers struct {
	ProductID  int
	SupplierID int

	CreatedAt *time.Time
	UpdatedAt *time.Time
}
