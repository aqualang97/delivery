package repositories

import (
	"database/sql"
	"delivery/internal/models"
	"github.com/aqualang97/logger/v4"
	"log"
)

type SuppliersCategoriesRepository struct {
	conn   *sql.DB
	TX     *sql.Tx
	logger *logger.Logger
}

func NewSuppliersCategoriesRepo(conn *sql.DB, TX *sql.Tx, logger *logger.Logger) *SuppliersCategoriesRepository {
	return &SuppliersCategoriesRepository{conn: conn, TX: TX, logger: logger}
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

func (s SuppliersCategoriesRepository) GetSupplierCategoryID(name string) (int, error) {
	var id int
	err := s.conn.QueryRow(
		"SELECT id FROM suppliers_categories WHERE name = ?",
		name).Scan(&id)
	if err != nil {
		log.Println(err)
	}
	return id, err
}
