package models

import "time"

type ProductsIngredients struct {
	ProductId    int
	IngredientID int
	CreatedAt    *time.Time
}
