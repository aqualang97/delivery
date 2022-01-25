package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

type ProductsSuppliersRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func NewProductsSuppliersRepo(conn *sql.DB, TX *sql.Tx) *ProductsSuppliersRepository {
	return &ProductsSuppliersRepository{conn: conn, TX: TX}
}

func (p ProductsSuppliersRepository) InsertProductSupplier(ps models.ProductsSuppliers) error {
	_, err := p.conn.Exec(
		"INSERT products_suppliers(product_id, supplier_id, external_product_id, external_supplier_id, price, image)VALUES(?, ?, ?, ?, ?, ?)",
		ps.ProductID, ps.SupplierID, ps.ExternalProductID, ps.ExternalSupplierID, ps.Price, ps.Image)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p ProductsSuppliersRepository) UpdatePriceByExternalData(ps models.ProductsSuppliers) error {
	_, err := p.conn.Exec("UPDATE  products_suppliers SET (price) VALUES(?) WHERE external_product_id=? AND external_supplier_id=?",
		ps.Price, ps.ExternalProductID, ps.ExternalSupplierID)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p ProductsSuppliersRepository) DeleteProductBySupplier(ps models.ProductsSuppliers) error {
	_, err := p.conn.Exec("DELETE FROM products_suppliers WHERE supplier_id=? AND external_supplier_id=?",
		ps.SupplierID, ps.ExternalSupplierID)
	if err != nil {
		log.Println(err)
	}
	return err
}
