package db

import (
	"database/sql"
	"delivery/internal/models"
	"fmt"
	"log"
)

var id int
var name, email, login, password string

type UserDBRepository struct {
	conn *sql.DB
}

func NewUserRepo(conn *sql.DB) UserDBRepository {
	return UserDBRepository{conn: conn}
}
func (udbr UserDBRepository) GetByEmail(email string) *models.User {
	// SELECT email, password_hash, created_at FROM users WHERE email = email

	rows, err := udbr.conn.Query("select * from users where email = ?", email)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		err := rows.Scan(&id, &login, &email, &password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, login, email)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &models.User{}
}

func (udbr UserDBRepository) GetUserById(id int) *models.User {
	rows, err := udbr.conn.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		err := rows.Scan(&id, &login, &email, &password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &models.User{}
}

func (udbr UserDBRepository) InsertToUsers(user *models.User) (int64, error) {
	rows, err := udbr.conn.Prepare("INSERT INTO users(login, email, password) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	res, err := rows.Exec(user.Name, user.Email, user.PasswordHash)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	log.Printf("ID = %d", lastId)
	return lastId, err
}
func (udbr UserDBRepository) UpdateById(user *models.User) int64 {
	rows, err := udbr.conn.Prepare("UPDATE  users(login, email, password) SET login, email, password VALUES(?,?,?) WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := rows.Exec(user.Name, user.Email, user.PasswordHash, user.Id)
	if err != nil {
		log.Fatal(err)
	}

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
