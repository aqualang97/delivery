package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

type ProductsIngredientsRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func NewProductsIngredientsRepo(conn *sql.DB, TX *sql.Tx) *ProductsIngredientsRepository {
	return &ProductsIngredientsRepository{conn: conn, TX: TX}
}

func (i ProductsIngredientsRepository) InsertProductIngredient(productID, ingredientID int) error {
	_, err := i.conn.Exec(
		"INSERT products_ingredients(product_id, ingredient_id)VALUES(?, ?)",
		productID, ingredientID)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (i ProductsIngredientsRepository) GetIngredientsByProductID(id int) ([]models.ProductsIngredients, error) {
	var prIn models.ProductsIngredients
	var listPrIn []models.ProductsIngredients
	rows, err := i.conn.Query(
		"SELECT product_id, ingredient_id FROM products_ingredients WHERE product_id = ?",
		id)
	if err != nil {
		log.Println(err)
		return listPrIn, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&prIn.ProductId, &prIn.IngredientID)
		if err != nil {
			log.Println(err)
			return listPrIn, err
		}
		listPrIn = append(listPrIn, prIn)
	}
	return listPrIn, err
}

func (i ProductsIngredientsRepository) DeleteIngredientByProduct(productID int) error {
	_, err := i.conn.Exec("DELETE * FROM products_ingredients WHERE product_id=?", productID)
	if err != nil {
		log.Println(err)
	}
	return err
}
