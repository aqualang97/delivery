package models

import "time"

type Menu struct {
	ID         int
	SupplierID int
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
