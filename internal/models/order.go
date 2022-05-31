package models

import "time"

type Order struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	Price         float32 `json:"price"`
	Status        string  `json:"status"`
	PaymentMethod string  `json:"payment_method"`
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
type RequestCardPay struct {
	Card        string `json:"card"`
	Status      string `json:"status"`
	LastOrderId int    `json:"last_order_id"`
}

type OldOrders struct {
	OrderId   int           `json:"order_id"`
	Status    string        `json:"status"`
	UserID    int           `json:"user_id"`
	FullPrice float32       `json:"full_price"`
	UserOrder []OldPosition `json:"user_order"`
}
type OldPosition struct {
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	Image         string  `json:"image"`
	Quantity      int     `json:"quantity"`
	PurchasePrice float32 `json:"purchase_price"`
}
