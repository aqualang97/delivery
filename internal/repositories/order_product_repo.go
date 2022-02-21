package repositories

import (
	"database/sql"
	"delivery/internal/models"
	"github.com/aqualang97/logger/v4"
	"log"
)

type OrderProductDBRepository struct {
	conn   *sql.DB
	TX     *sql.Tx
	logger *logger.Logger
}

func NewOrderProductRepo(conn *sql.DB, TX *sql.Tx, logger *logger.Logger) *OrderProductDBRepository {
	return &OrderProductDBRepository{conn: conn, TX: TX, logger: logger}
}

func (o OrderProductDBRepository) InsertToOrdersProducts(moList []models.OrderProducts) error {
	var err error
	for i := range moList {
		mo := moList[i]
		_, err := o.conn.Exec(
			"INSERT orders_products(product_id, order_id, numbers_of_product, purchase_price) VALUES(?, ?, ?, ?)",
			mo.ProductId, mo.OrderId, mo.NumbersOfProduct, mo.PurchasePrice)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	//id, err := res.LastInsertId()
	//if err != nil {
	//	log.Println(err)
	//}
	return err
}

func (o OrderProductDBRepository) UpdateNumbersByProductAndOrderID(mo models.OrderProducts) error {
	_, err := o.conn.Exec("UPDATE  orders_products SET (numbers_of_product) VALUES(?) WHERE product_id = ? AND order_id=?",
		mo.NumbersOfProduct, mo.ProductId, mo.OrderId)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func (o OrderProductDBRepository) GetAllProductsByOrderID(orderID int) ([]models.OrderProducts, error) {
	var mo models.OrderProducts
	var productsInOrderProd []models.OrderProducts

	rows, err := o.conn.Query("SELECT id, product_id, order_id, numbers_of_product, purchase_price, p.name FROM orders_products LEFT JOIN products as p ON orders_products.product_id=p.id WHERE order_id=?", orderID)
	if err != nil {
		log.Println(err)
		return productsInOrderProd, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&mo.Id, &mo.ProductId, &mo.OrderId, &mo.NumbersOfProduct, &mo.PurchasePrice, &mo.ProductName)
		if err != nil {
			log.Fatal(err)
		}
		productsInOrderProd = append(productsInOrderProd, mo)
	}
	return productsInOrderProd, err
}

func (o OrderProductDBRepository) DeleteProduct(order models.OrderProducts) error {
	_, err := o.conn.Exec("DELETE FROM orders_products WHERE product_id=? AND order_id=?", order.ProductId, order.OrderId)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (o OrderProductDBRepository) DeleteAll(order models.OrderProducts) error {
	_, err := o.conn.Exec("DELETE FROM orders_products WHERE order_id=?", order.OrderId)
	if err != nil {
		log.Println(err)
	}
	return err
}
