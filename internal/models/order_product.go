package models

type OrderProducts struct {
	Id        int
	ProductId int
	OrderId   int
	Numbers   int
	Price     float32
	//UserId    int
}
