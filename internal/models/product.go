package models

import "time"

type Product struct {
	Id          int
	Name        string
	Category    int
	Description string
	//Discount   int
	CreatedAt *time.Time
}
