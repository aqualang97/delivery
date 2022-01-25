package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

type UserContactRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func NewUserContactRepo(conn *sql.DB, TX *sql.Tx) *UserContactRepository {
	return &UserContactRepository{conn: conn, TX: TX}
}

func (c UserContactRepository) CreateUserInfo(data models.UserContactData) error {
	_, err := c.conn.Exec("INSERT user_id, first_name, last_name, phohe_number, address INTO users_contact_data VALUES(?,?,?,?,?)",
		data.UserID, data.FirstName, data.LastName, data.PhoneNumber, data.Address)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (c UserContactRepository) GetUserInfoByUserID(userID int) ([]models.UserContactData, error) {
	// Допускаю что юзер привяжет доп телефон
	var data models.UserContactData
	var listOfData []models.UserContactData

	rows, err := c.conn.Query(
		"SELECT * FROM users_contact_data WHERE user_id = ?",
		userID)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.ID, &data.UserID, &data.FirstName, &data.LastName, &data.PhoneNumber, &data.Address, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
		listOfData = append(listOfData, data)
	}
	return listOfData, err
}
func (c UserContactRepository) GetUserAddressByUSerID(userID int) (string, error) {
	var address string
	err := c.conn.QueryRow(
		"SELECT address FROM users_contact_data WHERE user_id=?",
		userID).Scan(&address)
	if err != nil {
		log.Println(err)
	}
	return address, err
}

func (c UserContactRepository) UpdateAddress(userID int, newAddress string) error {
	_, err := c.conn.Exec("UPDATE  users_contact_data SET address VALUES(?) WHERE user_id=?",
		newAddress, userID)
	if err != nil {
		log.Println(err)
	}
	return err
}