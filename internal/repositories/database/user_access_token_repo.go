package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

type UserAccessTokenRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func NewAccessTokenRepo(conn *sql.DB, TX *sql.Tx) *UserAccessTokenRepository {
	return &UserAccessTokenRepository{conn: conn, TX: TX}
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
	println(userToken.RefreshToken, userId)

	return err*/
	//println(userToken.UserID)
	_, err := t.conn.Exec("INSERT users_access_tokens(user_id, token, expired_at, expired) VALUES(?, ?, ?, ?)",
		userToken.UserID, userToken.AccessToken, userToken.ExpiredAt, userToken.Expired)
	if err != nil {
		return err
	}
	return err
}

func (t UserAccessTokenRepository) GetByAccessToken(accessToken string) (models.UserAccessToken, error) {
	var userToken models.UserAccessToken
	println(accessToken)
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

	println(userToken.UserID, userToken.AccessToken)
	if err != nil {
		return userToken, err
	}
	return userToken, nil
}

func (t UserAccessTokenRepository) UpdateOldAndInsertNewAccessToken(oldAccess string,
	response models.UserAccessToken) error {

	_, err := t.conn.Exec(
		"UPDATE users_access_tokens SET expired = ? WHERE token = ?", "true", oldAccess)
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
