package repositories

import (
	"database/sql"
	"delivery/internal/models"
	"fmt"
	"github.com/aqualang97/logger/v4"
	"log"
	"time"
)

type UserAccessTokenRepository struct {
	conn   *sql.DB
	TX     *sql.Tx
	logger *logger.Logger
}

func NewAccessTokenRepo(conn *sql.DB, TX *sql.Tx, logger *logger.Logger) *UserAccessTokenRepository {
	return &UserAccessTokenRepository{conn: conn, TX: TX, logger: logger}
}
func (t UserAccessTokenRepository) InsertAccessToken(userToken models.UserAccessToken) error {
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

	_, err := t.conn.Exec("INSERT users_access_tokens(user_id, token, expired_at, expired) VALUES(?, ?, ?, ?)",
		userToken.UserID, userToken.AccessToken, userToken.ExpiredAt, userToken.Expired)
	fmt.Println(userToken.ExpiredAt)

	if err != nil {
		return err
	}
	return err
}

func (t UserAccessTokenRepository) IsExistAccess(userID int) (bool, error) {
	var exist bool
	err := t.conn.QueryRow("SELECT EXISTS(SELECT * FROM users_access_tokens WHERE user_id=? AND expired=?)", userID, "false").Scan(&exist)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return exist, err
}

func (t UserAccessTokenRepository) GetAccessTokenByUserID(userID int) (string, error) {
	var tokenHash string
	err := t.conn.QueryRow("SELECT token FROM users_access_tokens WHERE user_id=? AND expired=?", userID, "false").Scan(&tokenHash)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenHash, err
}

func (t UserAccessTokenRepository) GetByAccessToken(accessToken string) (models.UserAccessToken, error) {
	var userToken models.UserAccessToken
	rows, err := t.conn.Query("SELECT user_id, token, expired FROM users_access_tokens WHERE token = ?", accessToken)
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

	if err != nil {
		return userToken, err
	}
	return userToken, nil
}

func (t UserAccessTokenRepository) UpdateOldAndInsertNewAccessToken(userID int,
	response models.UserAccessToken) error {

	_, err := t.conn.Exec(
		"UPDATE users_access_tokens SET expired = ? WHERE user_id = ?", "true", userID)
	if err != nil {
		return err
	}

	_, err = t.conn.Exec("INSERT users_access_tokens(user_id, token, expired_at, expired) VALUES(?, ?, ?, ?)",
		response.UserID, response.AccessToken, response.ExpiredAt, response.Expired)
	if err != nil {
		return err
	}

	return err
}

//func (t UserAccessTokenRepository)ExpiredAccessWithOneActiveDevice(userID int) error {
//	_, err := t.conn.Exec(
//		"UPDATE users_access_tokens SET expired = ? WHERE userID = ?", "true", userID)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return err
//}

func (t UserAccessTokenRepository) ExpiredAccessToken(userID int) error {
	//_, err := t.conn.Exec(
	//	"UPDATE users_access_tokens SET expired = ? WHERE token = ?", "true", oldToken)
	//if err != nil {
	//	log.Fatal(err)
	//	return err
	//}
	_, err := t.conn.Exec(
		"UPDATE users_access_tokens SET expired = ? WHERE user_id = ?", "true", userID)
	if err != nil {
		log.Fatal(err)
	}
	return err

}

func (t UserAccessTokenRepository) DeleteNaturallyExpiredAccessToken() {
	_, err := t.conn.Exec("DELETE FROM users_access_tokens WHERE expired_at<=?", time.Now())
	if err != nil {
		log.Fatal(err)
		return
	}
}
