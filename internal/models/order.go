package models

import "time"

type Order struct {
	ID        int
	UserID    int
	Price     float32
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
