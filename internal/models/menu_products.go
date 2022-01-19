package models

import "time"

type MenuProducts struct {
	ProductId int
	MenuID    int
	Price     float32
	Image     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
