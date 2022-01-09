package models

type Order struct {
	Id         string
	SupplierId string
	ProductId  string
	UserId     string
	Number     int
	Price      float32
	Status     string
	CreatedAt  string
}
