package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

type ProductDBRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func (p ProductDBRepository) GetProductByID(id int) (models.Product, error) {
	var product models.Product
	err := p.conn.QueryRow(
		"SELECT id, name, category, external_id FROM products WHERE id = ?",
		id).Scan(&product.Id, &product.Name, &product.Category, &product.ExternalID)
	if err != nil {
		log.Println(err)
	}
	return product, err
}

func (p ProductDBRepository) InsertToProducts(mp models.Product) (int, error) {
	res, err := p.conn.Exec(
		"INSERT products(name, price, external_id) VALUES(?, ?, ?)",
		mp.Name, mp.Category, mp.ExternalID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return int(id), err
}

func (p ProductDBRepository) UpdateProductById(mp models.Product) {
	// ?
	// Пока не представляю что можно тут обновлять
	//
	return
}

func (p ProductDBRepository) DeleteProductByExternalID(name string, externalID int) error {
	_, err := p.conn.Exec("DELETE FROM products WHERE name=? AND external_id=?", name, externalID)
	if err != nil {
		log.Println(err)
	}
	return err
}
func NewProductRepo(conn *sql.DB, TX *sql.Tx) *ProductDBRepository {
	return &ProductDBRepository{conn: conn, TX: TX}
}
