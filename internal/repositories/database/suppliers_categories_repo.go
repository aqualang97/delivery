package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

type SuppliersCategoriesRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func NewSuppliersCategoriesRepo(conn *sql.DB, TX *sql.Tx) *SuppliersCategoriesRepository {
	return &SuppliersCategoriesRepository{conn: conn, TX: TX}
}

func (s SuppliersCategoriesRepository) CreateCategory(categories models.SuppliersCategories) error {
	var exist bool
	err := s.conn.QueryRow("SELECT EXISTS(SELECT * FROM suppliers_categories WHERE name=?)", categories.Name).Scan(&exist)
	if err != nil {
		log.Println(err)
		return err
	}

	if !exist {
		_, err := s.conn.Exec("INSERT suppliers_categories(name) VALUES(?) ON DUPLICATE KEY UPDATE name=(?)",
			categories.Name, categories.Name)
		if err != nil {
			log.Println(err)
		}
	}
	return err
}
