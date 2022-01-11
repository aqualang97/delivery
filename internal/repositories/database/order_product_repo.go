package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

type OrderProductDBRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func (o OrderProductDBRepository) InsertToOrdersProducts(mo models.OrderProducts) (int, error) {
	var id int

	if o.TX != nil {
		err := o.TX.QueryRow(
			"INSERT orders_products(product_id, order_id, numbers_of_product, purchase_price) VALUES(?, ?, ?, ?) RETURNING id",
			mo.ProductId, mo.OrderId, mo.Numbers, mo.Price).Scan(&id)
		if err != nil {
			_ = o.TX.Rollback()
		}
		err = o.TX.Commit()
		if err != nil {
			_ = o.TX.Rollback()
		}
		return id, err
	}
	err := o.conn.QueryRow(
		"INSERT orders_products(product_id, order_id, numbers_of_product, purchase_price) VALUES(?, ?, ?, ?) RETURNING id",
		mo.ProductId, mo.OrderId, mo.Numbers, mo.Price).Scan(&id)

	return id, err
}

func (o OrderProductDBRepository) UpdateOrdersProductsById(mo *models.OrderProducts) int64 {
	rows, err := o.conn.Prepare(
		"UPDATE  orders_products SET (product_id, order_id, numbers_of_product, purchase_price) VALUES(?, ?, ?, ?) WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := rows.Exec(mo.ProductId, mo.OrderId, mo.Numbers, mo.Price)
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
func (o OrderProductDBRepository) GetOrdersProductsByID(id string) (models.OrderProducts, error) {
	var mo models.OrderProducts
	err := o.conn.QueryRow(
		"SELECT product_id, order_id, numbers_of_product, purchase_price FROM orders_products WHERE id = ?",
		id).Scan(mo.ProductId, mo.OrderId, mo.Numbers, mo.Price)
	if err != nil {
		return mo, err
	}
	return mo, nil
}
func (o OrderProductDBRepository) GetOrdersProductsByProductId(product_id string) (models.OrderProducts, error) {
	var mo models.OrderProducts
	err := o.conn.QueryRow(
		"SELECT id, product_id, order_id, numbers_of_product, purchase_price FROM orders_products WHERE product_id = ?",
		product_id).Scan(mo.Id, mo.ProductId, mo.OrderId, mo.Numbers, mo.Price)
	if err != nil {
		return mo, err
	}
	return mo, nil
}

func (o OrderProductDBRepository) GetOrdersProductsByOrderId(order_id string) (models.OrderProducts, error) {
	var mo models.OrderProducts
	err := o.conn.QueryRow(
		"SELECT id, product_id, order_id, numbers_of_product, purchase_price FROM orders_products WHERE order_id = ?",
		order_id).Scan(mo.Id, mo.ProductId, mo.OrderId, mo.Numbers, mo.Price)
	if err != nil {
		return mo, err
	}
	return mo, nil
}
func NewOrderProductRepo(conn *sql.DB) OrderProductDBRepository {
	return OrderProductDBRepository{conn: conn}
}
