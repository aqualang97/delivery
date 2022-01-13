package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type dbInfo struct {
	driverName     string
	dataSourceName string
}

func (db dbInfo) dbOpen(driverName, dataSourceName string) *sql.DB {
	dbOpen, err := sql.Open(
		"mysql",
		"oboznyi:123123@tcp(127.0.0.1:3306)/oboznyi_db",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer dbOpen.Close()

	err = dbOpen.Ping()
	if err != nil {
		//
	}
	return dbOpen
}
