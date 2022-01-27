package main

import (
	"database/sql"
	"delivery/development/parse_with_goroutines/parse-from-foodapi/request"
	"delivery/development/parse_with_goroutines/parse-from-foodapi/worker_pool"
	open "delivery/internal/repositories/database/connection"
	"log"
	"sync"
)

const url = "http://foodapi.true-tech.php.nixdev.co"
const endpointSupp = "/suppliers"
const endpointMenu = "/menu"

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

		wg.Add(1)
	}
	allSupp := request.GetSuppliers(url + endpointSupp)
	for i, s := range allSupp.Suppliers {
		println("shop", i)
		menu := request.GetMenuWithSuppID(url, endpointMenu, endpointSupp, i+1)
		s.Menu = menu.Menu
		pool.StartSendData <- s
	}

	pool.Stop()
	wg.Wait()
}
