package main

import (
	"database/sql"
	"delivery/development/parse_with_goroutines/models"
	"delivery/development/parse_with_goroutines/worker_pool"
	open "delivery/internal/repositories/database/connection"
	"fmt"
	"log"
	"sync"
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
	//TX, err := dbTXBegin(conn)
	//if err != nil {
	//	return
	//}
	//parse := parser.NewConnParse(conn, TX)
	//
	//var supp models.Supplier
	pool := worker_pool.NewPool(4)
	count := pool.Count
	wg := sync.WaitGroup{}

	for i := 0; i < count; i++ {
		go func(i int) {
			pool.Start(&wg, i, conn)
		}(i)

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
	wg.Wait()

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
