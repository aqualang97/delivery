package models

import "time"

type ProductsCategories struct {
	ID        int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
