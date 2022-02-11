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

func NewSupplierRepo(conn *sql.DB, TX *sql.Tx) *SupplierDBRepository {
	return &SupplierDBRepository{conn: conn, TX: TX}
}

//models.Supplier изменил на SupplierForParse Для парсинга

func (s SupplierDBRepository) CreateSupplier(supp models.SupplierForParse, categorySupplierID int) (int, error) {
	res, err := s.conn.Exec(
		"INSERT suppliers(name, category_of_supplier, start_of_work, end_of_work, image, external_id)VALUES(?, ?, ?, ?, ?, ?)",
		supp.Name, categorySupplierID, supp.WorkingHours.Opening, supp.WorkingHours.Closing, supp.Image, supp.ExternalID)
	if err != nil {
		log.Println(err)

		return 0, err
	}

	supplierID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return int(supplierID), err
}

func (s SupplierDBRepository) GetSupplierByID(id int) (models.Supplier, error) {
	var supp models.Supplier
	err := s.conn.QueryRow(
		"SELECT id, name, category_of_supplier, start_of_work, end_of_work, image, external_id FROM suppliers WHERE id = ?",
		id).Scan(&supp.ID, &supp.Name, &supp.CategoryOfSupplier, &supp.WorkingHours.Opening, &supp.WorkingHours.Closing, &supp.Image, &supp.ExternalID)
	return supp, err
}

func (s SupplierDBRepository) GetSupplierByName(name string) ([]models.Supplier, error) {
	var supp models.Supplier
	var listSupp []models.Supplier
	// может быть > 1 одинакового названия поставщика
	rows, err := s.conn.Query(
		"SELECT id, name, category_of_supplier, start_of_work, end_of_work, image, external_id FROM suppliers WHERE name = ?",
		name)
	if err != nil {
		log.Println(err)
		return listSupp, err
	}
	for rows.Next() {
		err = rows.Scan(&supp.ID, &supp.Name, &supp.CategoryOfSupplier, &supp.WorkingHours.Opening, &supp.WorkingHours.Closing, &supp.Image, &supp.ExternalID)
		if err != nil {
			log.Println(err)
			return listSupp, err
		}
		listSupp = append(listSupp, supp)
	}
	return listSupp, err
}

func (s SupplierDBRepository) GetAllSuppliers() ([]models.Supplier, error) {
	var supp models.Supplier
	var listSupp []models.Supplier
	rows, err := s.conn.Query(
		"SELECT id, name, category_of_supplier, start_of_work, end_of_work, image, external_id FROM suppliers")
	if err != nil {
		log.Println(err)
		return listSupp, err
	}

	for rows.Next() {
		err = rows.Scan(&supp.ID, &supp.Name, &supp.CategoryOfSupplier, &supp.WorkingHours.Opening, &supp.WorkingHours.Closing, &supp.Image, &supp.ExternalID)
		if err != nil {
			log.Println(err)
			return listSupp, err
		}
		listSupp = append(listSupp, supp)
	}
	return listSupp, err
}

//
//func (s SupplierDBRepository) GetSupplierCategoryID(supp models.Supplier) (int, error) {
//	var exist bool
//	var id int
//	err := s.conn.QueryRow("SELECT EXISTS(SELECT * FROM suppliers_categories WHERE name=?)",
//		supp.CategoryOfSupplier).Scan(&exist)
//	if err != nil {
//		log.Println(err)
//		return id, err
//	}
//
//	if !exist {
//		_, err := s.conn.Exec("INSERT suppliers_categories(name) VALUES(?) ON DUPLICATE KEY UPDATE name=(?)",
//			supp.CategoryOfSupplier, supp.CategoryOfSupplier)
//		if err != nil {
//			log.Println(err)
//
//			return 0, err
//		}
//	}
//	// Считаем что название категории уникальное
//
//	err = s.conn.QueryRow("SELECT id FROM suppliers_categories WHERE name=?", supp.CategoryOfSupplier).Scan(&id)
//	if err != nil {
//		log.Println(err)
//
//		return id, err
//	}
//	return id, err
//}

func (s SupplierDBRepository) UpdateWorkingHoursByID(ms models.Supplier) error {
	_, err := s.conn.Exec("UPDATE  suppliers SET start_of_work, end_of_work) VALUES(?, ?) WHERE id = ?",
		ms.WorkingHours.Opening, ms.WorkingHours.Closing, ms.ID)
	if err != nil {
		log.Println(err)
	}
	return err
}
