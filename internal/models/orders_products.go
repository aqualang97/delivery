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
