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

func (p ProductDBRepository) GetProductByID(id string) (models.Products, error) {
	var product models.Products
	err := p.conn.QueryRow(
		"SELECT id, name, price, description FROM products WHERE id = ?",
		id).Scan(&product.Id, &product.Name, product.Price, product.Description)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p ProductDBRepository) InsertToProducts(mp *models.Products) (int, error) {
	var id int

	if p.TX != nil {
		err := p.TX.QueryRow("INSERT products(name, price, description) VALUES(?, ?, ?) RETURNING id",
			mp.Name, mp.Price, mp.Description).Scan(&id)
		if err != nil {
			_ = p.TX.Rollback()
		}
		err = p.TX.Commit()
		if err != nil {
			_ = p.TX.Rollback()
		}
		return id, err
	}
	err := p.conn.QueryRow(
		"INSERT products(name, price, description) VALUES(?, ?, ?) RETURNING id",
		mp.Name, mp.Price, mp.Description).Scan(&id)

	return id, err
}

func (p ProductDBRepository) UpdateProductById(mp *models.Products) int64 {
	rows, err := p.conn.Prepare("UPDATE  products SET name , price, description VALUES(?,?,?) WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := rows.Exec(mp.Name, mp.Price, mp.Description, mp.Id)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Rows affected = %d", rowCnt)
	return rowCnt
}

func NewProductRepo(conn *sql.DB) ProductDBRepository {
	return ProductDBRepository{conn: conn}
}
