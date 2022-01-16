package main

import (
	"database/sql"
	"delivery/development/parse_with_goroutines/models"
	open "delivery/internal/repositories/database/connection"
	"log"
)

func main() {
	conn, err := open.OpenMyDB()
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	//TX, err := dbTXBegin(conn)
	if err != nil {
		return
	}
}

func InsertMenuType(conn *sql.DB, supplier models.SupplierJSON) (int, error) {
	var id int
	var category string
	err := conn.QueryRow(
		"SELECT exists ",
		id).Scan(&user.ID, &user.Email, &user.Name, &user.PasswordHash)
}
