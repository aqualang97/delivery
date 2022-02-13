package models

import "time"

type OrderProducts struct {
	Id               int
	ProductId        int
	OrderId          int
	NumbersOfProduct int
	PurchasePrice    float32
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	//UserId    int
}
