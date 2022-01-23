package database

import (
	"database/sql"
	repositories "delivery/internal/auth"
	"delivery/internal/models"
	"log"
)

var name string

type UserDBRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func NewUserRepo(conn *sql.DB, TX *sql.Tx) *UserDBRepository {
	return &UserDBRepository{conn: conn, TX: TX}
}
func (udbr UserDBRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := udbr.conn.QueryRow(
		"SELECT id, email, login, password FROM users WHERE id = ?",
		id).Scan(&user.ID, &user.Email, &user.Login, &user.PasswordHash)
	return user, err

}
func (udbr UserDBRepository) GetUserByEmail(email string) (models.User, error) {
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
		println(user.ID, user.Login, user.Email, user.PasswordHash)
	}
	if err != nil {
		return user, err
	}
	return user, nil
}
func (udbr UserDBRepository) GetUserByLogin(login string) (models.User, error) {
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
		println(user.ID, user.Login, user.Email, user.PasswordHash)
	}
	if err != nil {
		return user, err
	}
	return user, nil
}
func (udbr UserDBRepository) CreateUser(user *repositories.RegistrationRequest, passwordHash string) (int, error) {
	res, err := udbr.conn.Exec("INSERT users(login, email, password) VALUES(?, ?, ?)",
		&user.Name, &user.Email, passwordHash)
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

func (udbr UserDBRepository) UpdateUserById(user *models.User) int64 {
	rows, err := udbr.conn.Prepare(
		"UPDATE  users(login, email, password) SET login, email, password VALUES(?,?,?) WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := rows.Exec(user.Login, user.Email, user.PasswordHash, user.ID)
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

func (udbr UserDBRepository) InsertAccessAndRefreshTokens(userToken *models.UserToken) error {
	/*var userId int
	a := userToken.AccessExpiredAt.Format("2006-01-02 15:04:05")

	fmt.Println(userToken.UserId, userToken.AccessToken, a, userToken.Expired)

	if udbr.TX != nil {
		err := udbr.TX.QueryRow("INSERT users_access_tokens(user_id, token, expired_at, expired) VALUES(?, ?, ?, ?) RETURNING user_id",
			2, "123", "1123", "true").Scan(&userId)
		if err != nil {
			_ = udbr.TX.Rollback()
		}
		err = udbr.TX.Commit()
		if err != nil {
			_ = udbr.TX.Rollback()
		}
		return err
	}

	err := udbr.conn.QueryRow(
		"INSERT users_access_tokens(user_id, token, expired_at, expired) VALUES(?, ?, ?, ?) RETURNING user_id",
		2, "123", "1123", "true").Scan(&userId)
	println(userToken.RefreshToken, userId)

	return err*/
	//println(userToken.UserID)
	_, err := udbr.conn.Exec("INSERT users_access_tokens(user_id, token, expired_at, expired) VALUES(?, ?, ?, ?)",
		userToken.UserID, userToken.AccessToken, userToken.AccessExpiredAt, userToken.Expired)
	if err != nil {
		return err
	}
	_, err = udbr.conn.Exec("INSERT users_refresh_tokens(user_id, token, expired_at, expired) VALUES(?, ?, ?, ?)",
		userToken.UserID, userToken.RefreshToken, userToken.AccessExpiredAt, userToken.Expired)
	return err
}

func (udbr UserDBRepository) GetByRefreshToken(refreshToken string) (models.UserToken, error) {
	var userToken models.UserToken
	rows, err := udbr.conn.Query("SELECT user_id, token, expired FROM users_refresh_tokens WHERE token = ?", refreshToken)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&userToken.UserID, &userToken.RefreshToken, &userToken.Expired)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		return userToken, err
	}
	return userToken, nil
}
func (udbr UserDBRepository) GetByAccessToken(accessToken string) (models.UserToken, error) {
	var userToken models.UserToken
	println(accessToken)
	rows, err := udbr.conn.Query("SELECT user_id, token, expired FROM users_access_tokens WHERE token = ?", accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&userToken.UserID, &userToken.AccessToken, &userToken.Expired)
		if err != nil {
			log.Fatal(err)
		}
	}

	println(userToken.UserID, userToken.AccessToken)
	if err != nil {
		return userToken, err
	}
	return userToken, nil
}

func (udbr UserDBRepository) UpdateAccessAndRefreshTokens(userToken models.UserToken,
	response repositories.RefreshResponse) error {

	println(userToken.AccessToken, "\n", userToken.RefreshToken)
	_, err := udbr.conn.Exec(
		"UPDATE users_access_tokens SET expired = ? WHERE token = ?", "true", userToken.AccessToken)
	if err != nil {
		return err
	}
	_, err = udbr.conn.Exec(
		"UPDATE users_refresh_tokens SET expired = ? WHERE token = ?", "true", userToken.RefreshToken)
	if err != nil {
		return err
	}
	_, err = udbr.conn.Exec("INSERT users_access_tokens(user_id, token, expired_at, expired) VALUES(?, ?, ?, ?)",
		response.UserID, response.NewAccessToken, response.AccessExpiredAt, response.Expired)
	if err != nil {
		return err
	}
	_, err = udbr.conn.Exec("INSERT users_refresh_tokens(user_id, token, expired_at, expired) VALUES(?, ?, ?, ?)",
		response.UserID, response.NewRefreshToken, response.AccessExpiredAt, response.Expired)
	return err
}
