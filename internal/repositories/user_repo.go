package repositories

import (
	"database/sql"
	"delivery/internal/models"
	"github.com/aqualang97/logger/v4"
	"log"
)

var name string

type UserDBRepository struct {
	conn   *sql.DB
	TX     *sql.Tx
	logger *logger.Logger
}

func NewUserRepo(conn *sql.DB, TX *sql.Tx, logger *logger.Logger) *UserDBRepository {
	return &UserDBRepository{conn: conn, TX: TX, logger: logger}
}

func (udbr UserDBRepository) IsExistUserByLogin(login string) (bool, error) {
	var exist bool
	err := udbr.conn.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE login=? )", login).Scan(&exist)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return exist, err
}

func (udbr UserDBRepository) IsExistUserByEmail(email string) (bool, error) {
	var exist bool
	err := udbr.conn.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE email=? )", email).Scan(&exist)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return exist, err
}
func (udbr UserDBRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := udbr.conn.QueryRow(
		"SELECT id, email, login, password FROM users WHERE id = ?",
		id).Scan(&user.ID, &user.Email, &user.Login, &user.PasswordHash)
	return user, err

}
func (udbr UserDBRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	rows, err := udbr.conn.Query("SELECT id, login, email, password FROM users WHERE email = ?", email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Login, &user.Email, &user.PasswordHash)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		return &user, err
	}

	return &user, nil
}
func (udbr UserDBRepository) GetUserByLogin(login string) (*models.User, error) {
	var user models.User
	rows, err := udbr.conn.Query("SELECT id, login, email, password FROM users WHERE login = ?", login)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Login, &user.Email, &user.PasswordHash)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		return &user, err
	}
	return &user, nil
}
func (udbr UserDBRepository) CreateUser(user *models.RegistrationRequest, passwordHash string) (int, error) {
	res, err := udbr.conn.Exec("INSERT users(login, email, password) VALUES(?, ?, ?)",
		&user.Login, &user.Email, passwordHash)
	id, err := res.LastInsertId()
	// НЕ ПОЛУЧАЮТСЯ ТРАНЗАКЦИИ,НАДО ИХ. ВЕРНУТЬ ID НЕ МОГУ. НО ПОКА ТАК:

	/*
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
	*/
	return int(id), err
}

func (udbr UserDBRepository) UpdateUserById(user *models.User) error {
	_, err := udbr.conn.Exec("UPDATE  users(login, email, password) SET login, email, password VALUES(?,?,?) WHERE id = ?",
		user.Login, user.Email, user.PasswordHash, user.ID)
	if err != nil {
		log.Fatal(err)
		return err

	}
	return err
}
