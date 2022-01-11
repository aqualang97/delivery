package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

type OrderDBRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func (o OrderDBRepository) InsertToOrders(mo models.Order) (int, error) {
	var id int

	if o.TX != nil {
		err := o.TX.QueryRow(
			"INSERT orders(user_id, price, status) VALUES(?, ?, ?) RETURNING id",
			mo.UserId, mo.Price, mo.Status).Scan(&id)
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
		"INSERT orders(user_id, price, status)VALUES(?, ?, ?) RETURNING id",
		mo.UserId, mo.Price, mo.Status).Scan(&id)

	return id, err
}

func (o OrderDBRepository) UpdateOrdersById(mo *models.Order) int64 {
	rows, err := o.conn.Prepare(
		"UPDATE  orders SET (user_id, price, status) VALUES(?, ?, ?) WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := rows.Exec(mo.UserId, mo.Price, mo.Status)
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
func (o OrderDBRepository) GetOrderByID(id string) (models.Order, error) {
	var mo models.Order
	err := o.conn.QueryRow(
		"SELECT user_id, price, status FROM users WHERE id = ?",
		id).Scan(mo.UserId, mo.Price, mo.Status)
	if err != nil {
		return mo, err
	}
	return mo, nil
}
func NewOrderRepo(conn *sql.DB) OrderDBRepository {
	return OrderDBRepository{conn: conn}
}
