package worker_pool

import (
	"database/sql"
	"delivery/development/parse_with_goroutines/models"
	"delivery/development/parse_with_goroutines/models/parser"
	"sync"
)

//type Worker interface {
//	//Stop()
//	//Do(data interface{}, handlerIndex int) //data modelsParse.SupplierJSON
//}
//
//type Constructor func() Worker

type WorkerPool struct {
	Count         int
	StartSendData chan models.Supplier //modelsParse.SupplierJSON
	StopSend      chan bool            // like flag for switch
	//New           Constructor
}

func NewPool(count int) *WorkerPool {
	return &WorkerPool{
		Count:         count,
		StartSendData: make(chan models.Supplier), // modelsParse.SupplierJSON
		StopSend:      make(chan bool),
		//New:           new,
	}
}
func (pool *WorkerPool) Stop() {
	for i := 0; i < pool.Count; i++ {
		pool.StopSend <- false
	}

}

func (pool *WorkerPool) Start(wg *sync.WaitGroup, i int, conn *sql.DB) {
	//var wg *sync.WaitGroup
	var supp models.Supplier
	defer wg.Done()
	for {
		select {
		case supp = <-pool.StartSendData:
			parser.Parser(supp, conn)
		case <-pool.StopSend:
			return
		}
	}
	//for i := 0; i < pool.Count; i++ {
	//	workerPool := pool.New()
	//	go func(index int) {
	//		for {
	//			select {
	//			case data, ok := <-pool.StartSendData:
	//				if !ok {
	//					workerPool.Stop()
	//					return
	//				}
	//				if data == nil {
	//					continue
	//				}
	//				workerPool.Do(data, index)
	//			case <-pool.StopSend:
	//				workerPool.Stop()
	//				return
	//
	//			}
	//
	//		}
	//	}(i)
	//	wg.Add(1)
	//}
	//wg.Wait()
}
