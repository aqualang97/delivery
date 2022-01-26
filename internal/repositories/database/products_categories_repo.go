package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

type ProductsCategoriesRepo struct {
	conn *sql.DB
	TX   *sql.Tx
}

func NewProductsCategoriesRepo(conn *sql.DB, TX *sql.Tx) *ProductsCategoriesRepo {
	return &ProductsCategoriesRepo{conn: conn, TX: TX}
}

func (p ProductsCategoriesRepo) CreateCategory(category models.ProductsCategories) (int, error) {
	var exist bool
	var id int64
	err := p.conn.QueryRow("SELECT EXISTS(SELECT * FROM products_categories WHERE name=?)", category.Name).Scan(&exist)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if !exist {
		res, err := p.conn.Exec("INSERT products_categories(name) VALUES(?)ON DUPLICATE KEY UPDATE name=(?)",
			category.Name, category.Name)
		if err != nil {
			log.Println(err)
			return 0, err
		}
		id, err = res.LastInsertId()
	}

	return int(id), err
}

func (p ProductsCategoriesRepo) DeleteCategory(id int) error {
	_, err := p.conn.Exec("DELETE FROM products_categories WHERE id=?", id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p ProductsCategoriesRepo) GetCategoryByID(id int) (*models.ProductsCategories, error) {
	var mo *models.ProductsCategories
	err := p.conn.QueryRow("SELECT id, name FROM products_categories WHERE id=?)", id).Scan(&mo.ID, &mo.Name)
	if err != nil {
		log.Println(err)
	}
	return mo, err
}
