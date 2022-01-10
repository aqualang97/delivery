package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func db() {
	db, err := sql.Open(
		"mysql",
		"oboznyi:123123@tcp(127.0.0.1:3306)/oboznyi_db",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		//
	}
}
