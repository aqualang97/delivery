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

func NewOrderProductRepo(conn *sql.DB, TX *sql.Tx) *OrderProductDBRepository {
	return &OrderProductDBRepository{conn: conn, TX: TX}
}

func (o OrderProductDBRepository) InsertToOrdersProducts(mo models.OrderProducts) (int64, error) {
	res, err := o.conn.Exec(
		"INSERT orders_products(product_id, order_id, numbers_of_product, purchase_price) VALUES(?, ?, ?, ?)",
		mo.ProductId, mo.OrderId, mo.NumbersOfProduct, mo.PurchasePrice)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return id, err
}

func (o OrderProductDBRepository) UpdateNumbersByProductAndOrderID(mo models.OrderProducts) error {
	_, err := o.conn.Exec("UPDATE  orders_products SET (numbers_of_product) VALUES(?) WHERE product_id = ? AND order_id=?",
		mo.NumbersOfProduct, mo.ProductId, mo.OrderId)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func (o OrderDBRepository) GetAllProductsByOrderID(orderID int) ([]models.OrderProducts, error) {
	var mo models.OrderProducts
	var productsInOrder []models.OrderProducts

	rows, err := o.conn.Query("SELECT id, product_id, order_id, numbers_of_product, purchase_price FROM orders_products WHERE order_id=?", orderID)
	if err != nil {
		log.Println(err)
		return productsInOrder, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&mo.Id, &mo.ProductId, &mo.OrderId, &mo.NumbersOfProduct, &mo.PurchasePrice)
		if err != nil {
			log.Fatal(err)
		}
		productsInOrder = append(productsInOrder, mo)
	}
	return productsInOrder, err
}

func (o OrderDBRepository) DeleteProduct(order models.OrderProducts) error {
	_, err := o.conn.Exec("DELETE FROM orders_products WHERE product_id=? AND order_id=?", order.ProductId, order.OrderId)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (o OrderDBRepository) DeleteAll(order models.OrderProducts) error {
	_, err := o.conn.Exec("DELETE FROM orders_products WHERE order_id=?", order.OrderId)
	if err != nil {
		log.Println(err)
	}
	return err
}
