package main

import (
	"database/sql"
	"delivery/development/parse_with_goroutines/models/parser"
	"delivery/development/parse_with_goroutines/parse-from-foodapi/request"
	"delivery/development/parse_with_goroutines/parse-from-foodapi/worker_pool"
	"delivery/internal/models"
	open "delivery/internal/repositories/database/connection"
	"log"
	"sync"
	"time"
)

func dbTXBegin(conn *sql.DB) (*sql.Tx, error) {
	TX, err := conn.Begin()
	return TX, err

}
func main() {
	conn, err := open.OpenMyDB()
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	TX, err := dbTXBegin(conn)
	if err != nil {
		return
	}

	//var supp models.Supplier
	pool := worker_pool.NewPool(4)
	count := pool.Count
	wg := sync.WaitGroup{}

	for i := 0; i < count; i++ {
		go func(i int) {
			pool.Start(&wg, i, conn, TX)

		}(i)
		go func(i int) {
			pool.StartParsePrice(&wg, i, conn, TX)
		}(i)

		wg.Add(1)
	}
	allSupp := request.GetSuppliers()
	for i, s := range allSupp.Suppliers {
		println("shop", i)
		menu := request.GetMenuWithSuppID(i + 1)
		s.Menu = menu.Menu
		pool.StartSendData <- s
	}
	for {
		time.Sleep(60 * time.Second)
		for suppID, _ := range allSupp.Suppliers {
			listProdId := parser.ParseProdSuppByDB(suppID+1, conn, TX)
			for _, prodID := range listProdId {

				//можно конечно и не делать запрос на GetProductFromAPI,
				// но раз он есть, можем походить конкретно по продукту

				// либо могу выташить все экстернал айди из бд и пройтись по ним
				position := request.GetProductFromAPI(suppID+1, prodID)
				//_ = parser.ParsePriceToDB(position.Price, prodID, suppID+1, conn, TX)
				var prodSupp models.ProductsSuppliers
				prodSupp.Price = position.Price
				prodSupp.ExternalProductID = position.ExternalID
				prodSupp.ExternalSupplierID = suppID

				pool.StartSendProd <- prodSupp
			}
		}

	}
	pool.Stop()
	pool.StopParsePrice()
	wg.Wait()
}
