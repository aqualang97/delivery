package repositories

import (
	"database/sql"
	"delivery/internal/models"
	"github.com/aqualang97/logger/v4"
	"log"
	"time"
)

type UserRefreshTokenRepository struct {
	conn   *sql.DB
	TX     *sql.Tx
	logger *logger.Logger
}

func NewRefreshTokenRepo(conn *sql.DB, TX *sql.Tx, logger *logger.Logger) *UserRefreshTokenRepository {
	return &UserRefreshTokenRepository{conn: conn, TX: TX, logger: logger}
}

func (t UserRefreshTokenRepository) InsertRefreshToken(userToken models.UserRefreshToken) error {
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


	return err*/

	_, err := t.conn.Exec("INSERT users_refresh_tokens(user_id, token, expired_at, expired) VALUES(?, ?, ?, ?)",
		userToken.UserID, userToken.RefreshToken, userToken.ExpiredAt, userToken.Expired)
	if err != nil {
		return err
	}
	return err
}

func (t UserRefreshTokenRepository) IsExistRefresh(userID int) (bool, error) {
	var exist bool
	err := t.conn.QueryRow("SELECT EXISTS(SELECT * FROM users_refresh_tokens WHERE user_id=? AND expired=?)", userID, "false").Scan(&exist)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return exist, err
}

func (t UserRefreshTokenRepository) GetRefreshTokenByUserID(userID int) (string, error) {
	var tokenHash string
	err := t.conn.QueryRow("SELECT token FROM users_refresh_tokens WHERE user_id=? AND expired=?", userID, "false").Scan(&tokenHash)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenHash, err
}

func (t UserRefreshTokenRepository) GetByRefreshToken(refreshToken string) (models.UserRefreshToken, error) {
	var userToken models.UserRefreshToken
	rows, err := t.conn.Query("SELECT user_id, token, expired FROM users_refresh_tokens WHERE token = ?", refreshToken)
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

func (t UserRefreshTokenRepository) UpdateOldAndInsertNewRefreshToken(userID int,
	response models.UserRefreshToken) error {

	_, err := t.conn.Exec(
		"UPDATE users_refresh_tokens SET expired = ? WHERE user_id = ?", "true", userID)
	if err != nil {
		log.Fatal(err)

		return err
	}
	_, err = t.conn.Exec("INSERT users_refresh_tokens(user_id, token, expired_at, expired) VALUES(?, ?, ?, ?)",
		response.UserID, response.RefreshToken, response.ExpiredAt, response.Expired)
	if err != nil {
		log.Fatal(err)

		return err
	}

	return err
}

func (t UserRefreshTokenRepository) ExpiredRefreshToken(userID int) error {
	//часть фичи с только 1 девайс
	_, err := t.conn.Exec(
		"UPDATE users_refresh_tokens SET expired = ? WHERE user_id = ?", "true", userID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
func (t UserRefreshTokenRepository) DeleteNaturallyExpiredRefreshToken() {
	_, err := t.conn.Exec("DELETE FROM users_refresh_tokens WHERE expired_at<=?", time.Now())
	if err != nil {
		log.Fatal(err)
		return
	}
}
