package models

import "time"

type SuppliersCategories struct {
	ID        int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
