package database

//
//import (
//	"database/sql"
//	"delivery/internal/models"
//	"log"
//)
//
//type SupplierDBRepository struct {
//	conn *sql.DB
//	TX   *sql.Tx
//}
//
//func NewSupplierRepo(conn *sql.DB) SupplierDBRepository {
//	return SupplierDBRepository{conn: conn}
//}
//
//func (s SupplierDBRepository) GetSupplierByID(id string) (models.Supplier, error) {
//	var supp models.Supplier
//	err := s.conn.QueryRow(
//		"SELECT name, category_of_supplier, start_of_work, end_of_work, image FROM suppliers WHERE id = ?",
//		id).Scan(&supp.Name, &supp.CategoryOfSupplier, &supp.StartOfWork, &supp.EndOfWork, &supp.Image)
//	return supp, err
//}
//
////
////func (s SupplierDBRepository) GetSupplierCategoryID(supp models.Supplier) (int, error) {
////	var exist bool
////	var id int
////	err := s.conn.QueryRow("SELECT EXISTS(SELECT * FROM suppliers_categories WHERE name=?)",
////		supp.CategoryOfSupplier).Scan(&exist)
////	if err != nil {
////		log.Println(err)
////		return id, err
////	}
////
////	if !exist {
////		_, err := s.conn.Exec("INSERT suppliers_categories(name) VALUES(?) ON DUPLICATE KEY UPDATE name=(?)",
////			supp.CategoryOfSupplier, supp.CategoryOfSupplier)
////		if err != nil {
////			log.Println(err)
////
////			return 0, err
////		}
////	}
////	// Считаем что название категории уникальное
////
////	err = s.conn.QueryRow("SELECT id FROM suppliers_categories WHERE name=?", supp.CategoryOfSupplier).Scan(&id)
////	if err != nil {
////		log.Println(err)
////
////		return id, err
////	}
////	return id, err
////}
//
//func (s SupplierDBRepository) InsertToSuppliers(ms models.Supplier) (int, error) {
//	var id int
//
//	if s.TX != nil {
//		err := s.TX.QueryRow(
//			"INSERT suppliers(id, name, category_of_supplier, start_of_work, end_of_work) VALUES(?, ?, ?, ?, ?) RETURNING id",
//			nil, ms.Name, ms.CategoryOfSupplier, ms.StartOfWork, ms.EndOfWork).Scan(&id)
//		if err != nil {
//			_ = s.TX.Rollback()
//		}
//		err = s.TX.Commit()
//		if err != nil {
//			_ = s.TX.Rollback()
//		}
//		return id, err
//	}
//	err := s.conn.QueryRow(
//		"INSERT suppliers(name, category_of_supplier, start_of_work, end_of_work) VALUES(?, ?, ?, ?) RETURNING id",
//		ms.Name, ms.CategoryOfSupplier, ms.StartOfWork, ms.EndOfWork).Scan(&id)
//
//	return id, err
//}
//
//func (s SupplierDBRepository) UpdateSuppliersById(ms *models.Supplier) {
//	_, err := s.conn.Exec("UPDATE  suppliers SET name, category_of_supplier, start_of_work, end_of_work) VALUES(?, ?, ?, ?) WHERE id = ?",
//		&ms.Name, &ms.CategoryOfSupplier, &ms.StartOfWork, &ms.EndOfWork)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return
//}
