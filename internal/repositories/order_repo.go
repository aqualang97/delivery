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
		"INSERT orders(user_id, price, status)VALUES(?, ?, ?) RETURNING id",
		mo.UserID, mo.Price, mo.Status)
	id, _ := res.LastInsertId()
	return int(id), err
}

func (o OrderDBRepository) UpdateOrdersByID(mo *models.Order) error {
	_, err := o.conn.Exec("UPDATE  orders SET (price, status) VALUES( ?, ?) WHERE id = ?",
		mo.Price, mo.Status)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func (o OrderDBRepository) UpdateOrdersByUserID(mo *models.Order) error {
	_, err := o.conn.Exec("UPDATE  orders SET (price, status) VALUES(?, ?) WHERE user_id = ?",
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
