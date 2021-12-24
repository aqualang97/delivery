package models

type Order struct {
	Id         string
	SupplierId string
	ProductId  string
	UserId     string
	Number     int
}
