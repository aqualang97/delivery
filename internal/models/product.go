package models

import "time"

type Products struct {
	Id          int
	SupplierId  int
	Category    int
	Name        string
	Price       float32
	Description string
	//Discount   int
	CreatedAt time.Time
}
