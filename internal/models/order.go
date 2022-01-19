package models

import "time"

type Order struct {
	Id        int
	UserId    int
	Price     float32
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
