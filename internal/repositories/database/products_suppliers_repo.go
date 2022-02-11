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
func (p ProductsSuppliersRepository) IsExistProductSupplier(ps models.ProductsSuppliers) bool {
	var exist bool
	err := p.conn.QueryRow("SELECT EXISTS(SELECT * FROM products_suppliers WHERE external_product_id=? AND external_supplier_id=?)",
		ps.ExternalProductID, ps.ExternalSupplierID).Scan(&exist)
	if err != nil {
		log.Println(err)
	}
	return exist
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

func (p ProductsSuppliersRepository) UpdatePriceByExternalData(price float64, extProdID, extSuppID int) error {
	_, err := p.conn.Exec("UPDATE products_suppliers SET price=? WHERE external_product_id=? AND external_supplier_id=?",
		price, extProdID, extSuppID)
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

func (p ProductsSuppliersRepository) GetAllExternalProductIDByExternalSupplierID(extSuppID int) ([]int, error) {
	var extProdId int
	var listOfProdId []int
	rows, err := p.conn.Query(
		"SELECT external_product_id FROM products_suppliers WHERE external_supplier_id = ?",
		extSuppID)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&extProdId)
		if err != nil {
			log.Println(err)
		}
		listOfProdId = append(listOfProdId, extProdId)
	}
	return listOfProdId, err
}
