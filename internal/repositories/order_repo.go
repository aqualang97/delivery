package repositories

import (
	"database/sql"
	"delivery/internal/models"
	"github.com/aqualang97/logger/v4"
	"log"
)

type OrderDBRepository struct {
	conn   *sql.DB
	TX     *sql.Tx
	logger *logger.Logger
}

func NewOrderRepo(conn *sql.DB, TX *sql.Tx, logger *logger.Logger) *OrderDBRepository {
	return &OrderDBRepository{conn: conn, TX: TX, logger: logger}
}
func (o OrderDBRepository) InsertToOrders(mo models.Order) (int, error) {
	//if o.TX != nil {
	//	err := o.TX.QueryRow(
	//		"INSERT orders(user_id, price, status) VALUES(?, ?, ?) RETURNING id",
	//		mo.UserId, mo.Price, mo.Status).Scan(&id)
	//	if err != nil {
	//		_ = o.TX.Rollback()
	//	}
	//	err = o.TX.Commit()
	//	if err != nil {
	//		_ = o.TX.Rollback()
	//	}
	//	return id, err
	//}
	res, err := o.conn.Exec(
		"INSERT orders(user_id, price, payment_method, status)VALUES(?, ?, ?, ?)",
		mo.UserID, mo.Price, mo.PaymentMethod, mo.Status)
	if err != nil {
		log.Println(err)
	}
	id, _ := res.LastInsertId()
	return int(id), err
}

func (o OrderDBRepository) UpdateOrdersByID(mo *models.RequestCardPay) error {
	_, err := o.conn.Exec("UPDATE orders SET status=? WHERE id=?", mo.Status, mo.LastOrderId)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func (o OrderDBRepository) UpdateOrdersByUserID(mo *models.Order) error {
	_, err := o.conn.Exec("UPDATE  orders SET (price, status,) VALUES(?, ?) WHERE user_id = ?",
		mo.Price, mo.Status, mo.UserID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
func (o OrderDBRepository) GetOrderByID(id int) (models.Order, error) {
	var mo models.Order
	err := o.conn.QueryRow(
		"SELECT * FROM users WHERE id = ?",
		id).Scan(&mo.ID, &mo.UserID, &mo.Price, &mo.Status, &mo.CreatedAt, &mo.UpdatedAt)
	return mo, err
}

func (o OrderDBRepository) GetOldOrdersByUserID(UserID int) ([]models.OldOrders, error) {
	var orders models.OldOrders
	var ordersList []models.OldOrders
	var position models.OldPosition
	var positionList []models.OldPosition
	rowsO, err := o.conn.Query("SELECT p.name, o.id, o.status, o.price, o.user_id, op.numbers_of_product, op.purchase_price,pc.name, ps.image FROM orders AS o LEFT JOIN orders_products op ON o.id = op.order_id LEFT JOIN products p on p.id = op.product_id LEFT JOIN products_suppliers ps on p.id = ps.product_id LEFT JOIN products_categories pc on pc.id = p.category WHERE o.user_id=?", UserID)
	if err != nil {
		o.logger.Error("GetListOfProdBySupplier \n", err)
	}
	defer rowsO.Close()
	tempOrderId := 0
	tempIndex := 0

	for rowsO.Next() {
		err = rowsO.Scan(&position.Name, &orders.OrderId,
			&orders.Status, &orders.FullPrice, &orders.UserID, &position.Quantity,
			&position.PurchasePrice, &position.Category, &position.Image)
		if err != nil {
			log.Println(err)
			return ordersList, err
		}
		if tempOrderId == orders.OrderId {
			positionList = append(positionList, position)
			orders.UserOrder = positionList

			ordersList[tempIndex].UserOrder = orders.UserOrder
		} else {
			if len(ordersList) != 0 {
				tempIndex += 1
			}

			positionList = nil
			positionList = append(positionList, position)
			orders.UserOrder = positionList
			ordersList = append(ordersList, orders)
			tempOrderId = orders.OrderId
		}
		//oneOrderList = append(oneOrderList, oneOrder)
		//orders.UserOrder = oneOrderList

	}
	return ordersList, err
}

func (o OrderDBRepository) GetOrderByUserIDNotPaidNotCompleted(userID int) (models.Order, error) {
	var mo models.Order
	err := o.conn.QueryRow(
		"SELECT * FROM users WHERE user_id = ? AND status != 'completed' AND status!='paid'",
		userID).Scan(&mo.ID, &mo.UserID, &mo.Price, &mo.Status, &mo.CreatedAt, &mo.UpdatedAt)
	return mo, err
}

func (o OrderDBRepository) DeleteOrderByUserID(userID int) error {
	_, err := o.conn.Exec("DELETE FROM orders WHERE user_id=?", userID)
	if err != nil {
		log.Println(err)
	}
	return err
}
