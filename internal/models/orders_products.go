package models

import "time"

type OrderProducts struct {
	Id               int     `json:"id"`
	ProductId        int     `json:"productId"`
	ProductName      string  `json:"productName"`
	OrderId          int     `json:"orderId"`
	NumbersOfProduct int     `json:"numbersOfProduct"`
	PurchasePrice    float32 `json:"purchasePrice"`
	AccessToken      string  `json:"accessToken"`
	RefreshToken     string  `json:"refreshToken"`
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	//UserId    int
}

type ProductFromCart struct {
	ProductId        int     `json:"product_id"`
	PurchasePrice    float32 `json:"purchase_price"`
	NumbersOfProduct int     `json:"quantity"`
}
type RequestOrderProducts struct {
	ContactData   UserContactData   `json:"contact_data"`
	Cart          []ProductFromCart `json:"cart"`
	PaymentMethod string            `json:"payment_method"`
}
