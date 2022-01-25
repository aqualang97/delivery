package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

type IngredientRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func NewIngredientRepo(conn *sql.DB, TX *sql.Tx) *IngredientRepository {
	return &IngredientRepository{conn: conn, TX: TX}
}

func (i IngredientRepository) IsExistIngredient(ingredient string) bool {
	var exist bool
	err := i.conn.QueryRow("SELECT EXISTS(SELECT * FROM ingredients WHERE name=?)", ingredient).Scan(&exist)
	if err != nil {
		log.Println(err)
		return false
	}
	return exist
}
func (i IngredientRepository) InsertIngredient(ingredient string) error {
	_, err := i.conn.Exec("INSERT ingredients(name) VALUE(?) ON DUPLICATE KEY UPDATE name=(?)",
		ingredient, ingredient)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
	//ingredientID, err = res.LastInsertId()
	//if err != nil {
	//	log.Println(err)
	//
	//	return 0, err
	//}

}
func (i IngredientRepository) DeleteIngredient(ingredient string) error {
	_, err := i.conn.Exec("DELETE FROM ingredients WHERE name=?", ingredient)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (i IngredientRepository) GetIngredientByName(name string) (models.Ingredients, error) {
	var ingredient models.Ingredients
	res, err := i.conn.Query("SELECT * FROM ingredients WHERE name=(?)", name)
	if err != nil {
		log.Fatal(err)
		return ingredient, err
	}
	defer res.Close()
	for res.Next() {
		err := res.Scan(&ingredient.ID, &ingredient.Name, &ingredient.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}
	return ingredient, err
}

func (i IngredientRepository) GetIngredientByID(id string) (models.Ingredients, error) {
	var ingredient models.Ingredients
	res, err := i.conn.Query("SELECT * FROM ingredients WHERE id=(?)", id)
	if err != nil {
		log.Fatal(err)
		return ingredient, err
	}
	defer res.Close()
	for res.Next() {
		err := res.Scan(&ingredient.ID, &ingredient.Name, &ingredient.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}
	return ingredient, err
}
