package models

import "time"

type ProductsIngredients struct {
	ProductId    int
	IngredientID int
	CreatedAt    *time.Time
}
type ProductsIngredientsName struct {
	ProductId  int
	Ingredient string
	CreatedAt  *time.Time
}
