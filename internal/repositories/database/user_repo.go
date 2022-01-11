package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

var name string

type UserDBRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func NewUserRepo(conn *sql.DB) UserDBRepository {
	return UserDBRepository{conn: conn}
}
func (udbr UserDBRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := udbr.conn.QueryRow(
		"SELECT id, email, login FROM users WHERE id = ?",
		id).Scan(&user.Id, &user.Email, &user.Name, &user.PasswordHash)
	if err != nil {
		return user, err
	}
	return user, nil

}
func (udbr UserDBRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := udbr.conn.QueryRow(
		"SELECT id, login, email, password FROM users",
		email).Scan(&user.Id, &user.Name, &user.Email, &user.PasswordHash)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (udbr UserDBRepository) CreateUser(u models.User) (int, error) {
	var id int

	if udbr.TX != nil {
		err := udbr.TX.QueryRow("INSERT users(name, email, password_hash) VALUES(?, ?, ?) RETURNING id", u.Name, u.Email, u.PasswordHash).Scan(&id)
		if err != nil {
			_ = udbr.TX.Rollback()
		}
		err = udbr.TX.Commit()
		if err != nil {
			_ = udbr.TX.Rollback()
		}
		return id, err
	}
	err := udbr.conn.QueryRow(
		"INSERT users(name, email, password_hash) VALUES(?, ?, ?) RETURNING id",
		u.Name, u.Email, u.PasswordHash).Scan(&id)

	return id, err
}

func (udbr UserDBRepository) UpdateById(user *models.User) int64 {
	rows, err := udbr.conn.Prepare(
		"UPDATE  users(login, email, password) SET login, email, password VALUES(?,?,?) WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := rows.Exec(user.Name, user.Email, user.PasswordHash, user.Id)
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
