package main

import (
	"database/sql"
	"delivery/development/parse_with_goroutines/models"
	"delivery/development/parse_with_goroutines/worker_pool"
	open "delivery/internal/repositories/database/connection"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	conn, err := open.OpenMyDB()
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//TX, err := dbTXBegin(conn)
	if err != nil {
		return
	}
	//var supp models.Supplier
	pool := worker_pool.NewPool(4)
	count := pool.Count
	wg := sync.WaitGroup{}
	runtime.GOMAXPROCS(1) //чтобы не  обрывалась мейн до завершения горутин

	wg.Add(count)
	for i := 0; i < count; i++ {
		go pool.Start(&wg, func(supp models.Supplier) {
			productCategoryID, _ := GetProductCategoryID(supp, conn)
			supplierCategoryID, _ := GetSupplierCategoryID(supp, conn)
			supplierID, _ := CreateSupplier(supp, conn, supplierCategoryID)
			productID, _ := CreateProduct(supp, conn, productCategoryID)
			fmt.Println("supplierID", supplierID, "productID", productID)
			time.Sleep(3 * time.Second) //для проверки. сначала 4 горутины, потом 3

		})
		wg.Add(1)
	}

	for i := 1; i <= 7; i++ {
		var rest models.Supplier
		//что-то с путем data/supplier_%d.json, не хочет находить файл никакой
		rest, err := models.ReadFromJSON(fmt.Sprintf("development/parse_with_goroutines/data/supplier_%d.json", i))
		if err != nil {
			log.Println(err)
			return
		}
		pool.StartSendData <- rest

	}
	pool.Stop()

	//pool.StartSendData = make(chan interface{})
	//runtime.GOMAXPROCS(1)
	//go pool.Start()
	//
	//for i := 1; i <= 7; i++ {
	//
	//	pool.StartSendData <- i
	//}

}

//
//type Parser struct {
//}
//
//func (p Parser) Do(data interface{}, i int) {
//	supp := models.Supplier{}
//	supp, _ = models.ReadFromJSON("data/supplier_1.json")
//	println(supp.Name)
//	fmt.Printf("Go Num %d, file %d, input %d \n", i, i, supp.Id)
//	time.Sleep(3 * time.Second)
//
//}
//
//func (p Parser) Stop() {
//
//}

//
//func InsertMenuType(conn *sql.DB, supplier models.SupplierJSON) (int, error) {
//	var id int
//	var category string
//}

func GetProductCategoryID(supp models.Supplier, conn *sql.DB) (int, error) {
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
			_, err := conn.Exec("INSERT products_categories(name) VALUES(?)",
				product.Type)
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

func GetSupplierCategoryID(supp models.Supplier, conn *sql.DB) (int, error) {
	var exist bool
	var id int
	err := conn.QueryRow("SELECT EXISTS(SELECT * FROM suppliers_categories WHERE name=?)", supp.Type).Scan(&exist)
	if err != nil {
		log.Println(err)
		return id, err
	}
	if !exist {
		_, err := conn.Exec("INSERT suppliers_categories(name) VALUES(?)",
			supp.Type)
		if err != nil {
			log.Println(err)

			return 0, err
		}
	}
	err = conn.QueryRow("SELECT id FROM suppliers_categories WHERE name=?", supp.Type).Scan(&id)
	if err != nil {
		log.Println(err)

		return id, err
	}
	return id, err
}

func CreateSupplier(supp models.Supplier, conn *sql.DB, categorySupplierID int) (int, error) {
	_, err := conn.Exec(
		"INSERT suppliers(id, name, category_of_supplier, start_of_work, end_of_work, image)VALUES(?, ?, ?, ?, ?, ?)",
		supp.Id, supp.Name, categorySupplierID, supp.WorkingHours.Opening, supp.WorkingHours.Closing, supp.Image)
	if err != nil {
		log.Println(err)

		return 0, err
	}
	supplierID := supp.Id
	return supplierID, err
}
func CreateProduct(supp models.Supplier, conn *sql.DB, categoryProductID int) (int, error) {
	//-------------------------------------
	//Переписать через транзакции
	// ВПРОЧЕМ КАК И ВСЁ :)
	//-------------------------------------

	var err error

	for _, product := range supp.Menu {

		_, err := conn.Exec(
			"INSERT products(id, name, price, image)VALUES(?, ?, ?, ?)",
			product.Id, product.Name, product.Price, product.Image)

		if err != nil {

			log.Println(err)
			return 0, err
		}
		//This code is only for the case when one product has one category.

		_, err = conn.Exec(
			"INSERT products_category(product_id, category_id)VALUES(?, ?)",
			product.Id, categoryProductID)
		if err != nil {

			log.Println(err)
			return 0, err
		}

		ingredients := product.Ingredients
		for _, ing := range ingredients {

			var exist bool
			var ingredientID int
			err := conn.QueryRow("SELECT EXISTS(SELECT * FROM ingredients WHERE name=?)", ing).Scan(&exist)
			if err != nil {
				log.Println(err)
				return ingredientID, err
			}
			if !exist {
				_, err := conn.Exec("INSERT ingredients(name) VALUE(?)",
					ing)
				if err != nil {
					log.Println(err)

					return 0, err
				}
			}
			err = conn.QueryRow("SELECT id FROM ingredients WHERE name=?", ing).Scan(&ingredientID)
			if err != nil {
				log.Println(err)

				return ingredientID, err
			}
			_, err = conn.Exec(
				"INSERT products_ingredients(product_id, ingredient_id)VALUES(?, ?)",
				product.Id, ingredientID)
			if err != nil {
				log.Println(err)
				return 0, err
			}
		}
	}
	supplierID := supp.Id

	return supplierID, err

}
