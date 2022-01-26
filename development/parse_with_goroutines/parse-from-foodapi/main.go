package main

//
//import (
//	"delivery/development/parse_with_goroutines/worker_pool"
//	open "delivery/internal/repositories/database/connection"
//	"log"
//	"net/http"
//	"sync"
//)
//
//func main() {
//	conn, err := open.OpenMyDB()
//	defer conn.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer conn.Close()
//	//TX, err := dbTXBegin(conn)
//	//if err != nil {
//	//	return
//	//}
//	//parse := parser.NewConnParse(conn, TX)
//	//
//	//var supp models.Supplier
//	pool := worker_pool.NewPool(4)
//	count := pool.Count
//	wg := sync.WaitGroup{}
//
//	for i := 0; i < count; i++ {
//		go func(i int) {
//			pool.Start(&wg, i, conn)
//		}(i)
//
//		wg.Add(1)
//	}
//	resp, err := http.Get("https://www.google.com/")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	println(resp.Body)
//	pool.Stop()
//	wg.Wait()
//}
