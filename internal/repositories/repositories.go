package repositories

//
//import (
//	"delivery/internal/models"
//)
//
//type UserRepositoryInterface interface {
//	GetByEmail(email string) models.User
//	GetByID(id string) models.User
//	InsertToUsers(user *models.User) (int, error)
//	//	Insert(user *models.User) models.User
//	UpdateById(user *models.User)
//}
//type SuppliersRepositoryInterface interface {
//	GetByName(name string) models.Suppliers
//	Insert(supplier *models.Suppliers) models.Suppliers
//	DeleteSupplier(name string) models.Suppliers
//	Update(supplier *models.Suppliers) models.Suppliers
//}
//
////
////type ProductsRepositoryInterface interface {
////	GetByCategory(category string) models.Products
////	Insert(product *models.Products) models.Products
////	Delete(product *models.Products) models.Products
////	UpdatePrise(product *models.Products) models.Products
////}
//
//type Order interface {
//	InsertProduct(order *models.Order) error
//	DeleteOrderById(id string) models.Order
//	ClearAllByUser(userid string) models.Order
//	ChangeNumber(id string) models.Order
//}
//
//type OrderProducts interface {
//	InsertProduct(order *models.Order) error
//	DeleteOrderById(id string) error
//	ClearAllByUser(userid string) error
//	ChangeNumber(id string) error
//}
