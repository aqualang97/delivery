package main

import (
	"database/sql"
	"delivery/internal/middlware"
	db "delivery/internal/repositories/database"
	//connection "delivery/internal/repositories/database/connection"
	handProv "delivery/internal/auth/handle_provide"
	open "delivery/internal/repositories/database/connection"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

const AccessSecret = "access_secret"
const RefreshSecret = "refresh_secret"

const AccessTokenLifetimeMinutes = 20
const RefreshTokenLifetimeMinutes = 60

func dbTXBegin(conn *sql.DB) (*sql.Tx, error) {
	TX, err := conn.Begin()
	return TX, err

}

func dbOpen() (*sql.DB, error) {

	dbConn, err := sql.Open(
		"mysql",
		"oboznyi:123123@tcp(127.0.0.1:3306)/oboznyi_db",
	)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	//var id int
	//var name string
	//rows, err := db.Query("SELECT id, login FROM users WHERE id = ?", 1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for rows.Next() {
	//	err := rows.Scan(&id, &name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(id, name)
	//
	//}
	return dbConn, err
}

func main() {
	//реализуем флоу логина юзера
	// юзер дает логин пароль
	// получаем ответ верный илинет юзе
	conn, err := open.OpenMyDB()
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	TX, err := dbTXBegin(conn)
	if err != nil {
		return
	}
	handlerProvider := &handProv.HandlerProvider{
		UserRepository:             db.NewUserRepo(conn, TX),
		UserAccessTokenRepository:  db.NewAccessTokenRepo(conn, TX),
		UserRefreshTokenRepository: db.NewRefreshTokenRepo(conn, TX),
	}
	m := middlware.NewMiddleware(handlerProvider)
	handlerProvider.UserAccessTokenRepository.DeleteNaturallyExpiredAccessToken()
	handlerProvider.UserRefreshTokenRepository.DeleteNaturallyExpiredRefreshToken()
	http.HandleFunc("/login", handlerProvider.Login) //умеем обрабатывать логин с помощью ф-ции логин
	http.Handle("/profile", m.RequireAuthentication(http.HandlerFunc(handlerProvider.Profile)))
	http.HandleFunc("/refresh", handlerProvider.Refresh)
	http.HandleFunc("/registration", handlerProvider.Registration)
	http.Handle("/logout", m.RequireAuthentication(http.HandlerFunc(handlerProvider.Logout)))

	log.Fatal(http.ListenAndServe(":8080", nil)) //слушаем порт 8080 для входящих запросов

}
