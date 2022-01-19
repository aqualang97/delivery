package models

import "time"

type SuppliersCategories struct {
	ID        int
	name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
