package database

import (
	"database/sql"
	"delivery/internal/models"
	"log"
)

type SupplierDBRepository struct {
	conn *sql.DB
	TX   *sql.Tx
}

func (s SupplierDBRepository) GetSupplierByID(id string) (models.Suppliers, error) {
	var supp models.Suppliers
	err := s.conn.QueryRow(
		"SELECT name, category_of_supplier, start_of_work, end_of_work FROM suppliers WHERE id = ?",
		id).Scan(&supp.Name, &supp.CategoryOfSupplier, supp.StartOfWork, supp.EndOfWork)
	return supp, err
}

func (s SupplierDBRepository) InsertToSuppliers(ms models.Suppliers) (int, error) {
	var id int

	if s.TX != nil {
		err := s.TX.QueryRow(
			"INSERT suppliers(id, name, category_of_supplier, start_of_work, end_of_work) VALUES(?, ?, ?, ?, ?) RETURNING id",
			nil, ms.Name, ms.CategoryOfSupplier, ms.StartOfWork, ms.EndOfWork).Scan(&id)
		if err != nil {
			_ = s.TX.Rollback()
		}
		err = s.TX.Commit()
		if err != nil {
			_ = s.TX.Rollback()
		}
		return id, err
	}
	err := s.conn.QueryRow(
		"INSERT suppliers(name, category_of_supplier, start_of_work, end_of_work) VALUES(?, ?, ?, ?) RETURNING id",
		ms.Name, ms.CategoryOfSupplier, ms.StartOfWork, ms.EndOfWork).Scan(&id)

	return id, err
}

func (s SupplierDBRepository) UpdateSuppliersById(ms *models.Suppliers) int64 {
	rows, err := s.conn.Prepare(
		"UPDATE  suppliers SET name, category_of_supplier, start_of_work, end_of_work) VALUES(?, ?, ?, ?) WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := rows.Exec(ms.Name, ms.CategoryOfSupplier, ms.StartOfWork, ms.EndOfWork)
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

func NewSupplierRepo(conn *sql.DB) SupplierDBRepository {
	return SupplierDBRepository{conn: conn}
}
