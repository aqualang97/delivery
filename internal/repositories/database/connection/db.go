package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func OpenMyDB() (*sql.DB, error) {
	conn, err := sql.Open(
		"mysql",
		"oboznyi:123123@tcp(127.0.0.1:3306)/oboznyi_db",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		//
	}
	return conn, err
}
