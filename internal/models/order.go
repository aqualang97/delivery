package models

import "time"

type Order struct {
	Id         int
	SupplierId int
	ProductId  int
	UserId     int
	Number     int
	Price      float32
	Status     string
	CreatedAt  time.Time
}
