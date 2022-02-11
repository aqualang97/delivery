package parser

import (
	"database/sql"
	tempModels "delivery/development/parse_with_goroutines/models"
	"delivery/internal/models"
	db "delivery/internal/repositories/database"
	"fmt"
	"log"
)

//type ConnParse struct {
//	conn *sql.DB
//	TX   *sql.Tx
//}
//
//func NewConnParse(conn *sql.DB, TX *sql.Tx) *ConnParse {
//	return &ConnParse{conn: conn, TX: TX}
//}

type ConnDBParse struct {
	IngredientRepo          *db.IngredientRepository
	ProductRepo             *db.ProductDBRepository
	ProductsCategoriesRepo  *db.ProductsCategoriesRepo
	ProductsIngredientsRepo *db.ProductsIngredientsRepository
	ProductsSuppliersRepo   *db.ProductsSuppliersRepository
	SupplierRepo            *db.SupplierDBRepository
	SuppliersCategoriesRepo *db.SuppliersCategoriesRepository
}

func Parser(supp tempModels.Supplier, conn *sql.DB) {

	productCategoryID, _ := GetProductCategoryID(supp, conn)
	supplierCategoryID, _ := GetSupplierCategoryID(supp, conn)
	supplierID, _ := CreateSupplier(supp, conn, supplierCategoryID)
	productExternalID, _ := CreateProduct(supp, conn, productCategoryID, supplierID)
	fmt.Println("supplierID", supplierID, "productID", productExternalID)

}
func ParseFromAPI(supp models.SupplierForParse, goNum int, conn *sql.DB, TX *sql.Tx) {
	connection := ConnDBParse{
		IngredientRepo:          db.NewIngredientRepo(conn, TX),
		ProductRepo:             db.NewProductRepo(conn, TX),
		ProductsCategoriesRepo:  db.NewProductsCategoriesRepo(conn, TX),
		ProductsIngredientsRepo: db.NewProductsIngredientsRepo(conn, TX),
		ProductsSuppliersRepo:   db.NewProductsSuppliersRepo(conn, TX),
		SupplierRepo:            db.NewSupplierRepo(conn, TX),
		SuppliersCategoriesRepo: db.NewSuppliersCategoriesRepo(conn, TX),
	}
	suppCat := models.SuppliersCategories{
		Name: supp.CategoryOfSupplier,
	}

	_ = connection.SuppliersCategoriesRepo.CreateCategory(suppCat)
	suppCatID, _ := connection.SuppliersCategoriesRepo.GetSupplierCategoryID(supp.CategoryOfSupplier)
	suppId, _ := connection.SupplierRepo.CreateSupplier(supp, suppCatID)
	menu := supp.Menu
	for _, product := range menu {
		prodCat := models.ProductsCategories{
			Name: product.Type,
		}

		productCategoryID, _ := connection.ProductsCategoriesRepo.CreateCategory(prodCat)

		productID, _ := connection.ProductRepo.InsertToProducts(product, productCategoryID)
		prodSupModel := models.ProductsSuppliers{
			ProductID:          productID,
			SupplierID:         suppId,
			ExternalProductID:  product.ExternalID,
			ExternalSupplierID: supp.ExternalID,
			Price:              product.Price,
			Image:              product.Image,
		}
		exist := connection.ProductsSuppliersRepo.IsExistProductSupplier(prodSupModel)
		if !exist {
			//
			_ = connection.ProductsSuppliersRepo.InsertProductSupplier(prodSupModel)
			//
			ingredients := product.Ingredients
			for _, ing := range ingredients {
				exist := connection.IngredientRepo.IsExistIngredient(ing)
				if !exist {
					_ = connection.IngredientRepo.InsertIngredient(ing)
				}
				ingId, _ := connection.IngredientRepo.GetIngredientIDByName(ing)
				//
				_ = connection.ProductsIngredientsRepo.InsertProductIngredient(productID, ingId)
				//
			}
		}

	}

	println("goNum", goNum)

}

func ParseProdSuppByDB(extSuppID int, conn *sql.DB, TX *sql.Tx) []int {
	connection := ConnDBParse{
		ProductsSuppliersRepo: db.NewProductsSuppliersRepo(conn, TX),
	}
	listProdId, _ := connection.ProductsSuppliersRepo.GetAllExternalProductIDByExternalSupplierID(extSuppID)
	return listProdId
}
func ParsePriceToDB(price float64, extProdID, extSuppID, goNum int, conn *sql.DB, TX *sql.Tx) error {
	connection := ConnDBParse{
		ProductsSuppliersRepo: db.NewProductsSuppliersRepo(conn, TX),
	}
	err := connection.ProductsSuppliersRepo.UpdatePriceByExternalData(price, extProdID, extSuppID)
	println("goNum", goNum, "Prod", extProdID)

	return err
}

func GetProductCategoryID(supp tempModels.Supplier, conn *sql.DB) (int, error) {
	var exist bool
	var id int
	var err error
	for _, product := range supp.Menu {

		err := conn.QueryRow("SELECT EXISTS(SELECT * FROM products_categories WHERE name=?)", product.Type).Scan(&exist)
		if err != nil {
			log.Println(err)
			return id, err
		}

		if !exist {
			_, err := conn.Exec("INSERT products_categories(name) VALUES(?) ON DUPLICATE KEY UPDATE name=(?)",
				product.Type, product.Type)
			if err != nil {
				log.Println(err)

				return 0, err
			}
		}
		err = conn.QueryRow("SELECT id FROM products_categories WHERE name=?", product.Type).Scan(&id)
		if err != nil {
			log.Println(err)

			return id, err
		}

	}
	log.Println(err)
	return id, err
}

func GetSupplierCategoryID(supp tempModels.Supplier, conn *sql.DB) (int, error) {
	var exist bool
	var id int
	err := conn.QueryRow("SELECT EXISTS(SELECT * FROM suppliers_categories WHERE name=?)", supp.Type).Scan(&exist)
	if err != nil {
		log.Println(err)
		return id, err
	}

	if !exist {
		_, err := conn.Exec("INSERT suppliers_categories(name) VALUES(?) ON DUPLICATE KEY UPDATE name=(?)",
			supp.Type, supp.Type)
		if err != nil {
			log.Println(err)

			return 0, err
		}
	}
	// Считаем что название категории уникальное

	err = conn.QueryRow("SELECT id FROM suppliers_categories WHERE name=?", supp.Type).Scan(&id)
	if err != nil {
		log.Println(err)

		return id, err
	}
	return id, err
}

func CreateSupplier(supp tempModels.Supplier, conn *sql.DB, categorySupplierID int) (int, error) {
	res, err := conn.Exec(
		"INSERT suppliers(name, category_of_supplier, start_of_work, end_of_work, image, external_id)VALUES(?, ?, ?, ?, ?, ?)",
		supp.Name, categorySupplierID, supp.WorkingHours.Opening, supp.WorkingHours.Closing, supp.Image, supp.ExternalId)
	if err != nil {
		log.Println(err)

		return 0, err
	}

	supplierID, err := res.LastInsertId()
	return int(supplierID), err
}
func CreateProduct(supp tempModels.Supplier, conn *sql.DB, categoryProductID, supplierID int) (int, error) {
	//-------------------------------------
	//Переписать через транзакции
	// ВПРОЧЕМ КАК И ВСЁ :)
	//-------------------------------------

	var err error

	for _, product := range supp.Menu {

		res, err := conn.Exec(
			"INSERT products(name, category, external_id)VALUES(?, ?, ?)",
			product.Name, categoryProductID, product.ExternalId)

		if err != nil {

			log.Println(err)
			return 0, err
		}

		productID, err := res.LastInsertId()
		if err != nil {
			return int(productID), err
		}
		//отсюда переписать под новую бд

		_, err = conn.Exec(
			"INSERT products_suppliers(product_id, supplier_id, external_product_id, external_supplier_id, price, image)VALUES(?, ?, ?, ?, ?, ?)",
			productID, supplierID, product.ExternalId, supp.ExternalId, product.Price, product.Image)

		if err != nil {

			log.Println(err)
			return 0, err
		}

		//This code is only for the case when one product has one category.

		//_, err = conn.Exec(
		//	"INSERT products_category(product_id, category_id)VALUES(?, ?)",
		//	product.Id, categoryProductID)
		//if err != nil {
		//
		//	log.Println(err)
		//	return 0, err
		//}

		ingredients := product.Ingredients
		for _, ing := range ingredients {

			var exist bool
			err := conn.QueryRow("SELECT EXISTS(SELECT * FROM ingredients WHERE name=?)", ing).Scan(&exist)
			if err != nil {
				log.Println(err)
				return 0, err
			}
			var ingredientID int64
			if !exist {
				res, err := conn.Exec("INSERT ingredients(name) VALUE(?) ON DUPLICATE KEY UPDATE name=(?)",
					ing, ing)
				if err != nil {
					log.Println(err)

					return 0, err
				}
				ingredientID, err = res.LastInsertId()
				if err != nil {
					log.Println(err)

					return 0, err
				}
			}
			err = conn.QueryRow("SELECT id FROM ingredients WHERE name=?", ing).Scan(&ingredientID)
			if err != nil {
				log.Println(err)

				return int(ingredientID), err
			}
			_, err = conn.Exec(
				"INSERT products_ingredients(product_id, ingredient_id)VALUES(?, ?)",
				productID, ingredientID)
			if err != nil {
				log.Println(err)
				return 0, err
			}
		}
	}
	productID := supp.ExternalId

	return productID, err

}
