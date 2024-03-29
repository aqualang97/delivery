package worker_pool

import (
	"database/sql"
	"delivery/internal/models"
	"delivery/internal/parser"
	"github.com/aqualang97/logger/v4"
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
	StartSendData chan models.SupplierForParse //modelsParse.SupplierJSON
	StopSend      chan bool                    // like flag for switch
	StartSendProd chan models.ProductsSuppliers
	StopSendProd  chan bool
	//New           Constructor
}

func NewPool(count int) *WorkerPool {
	return &WorkerPool{
		Count:         count,
		StartSendData: make(chan models.SupplierForParse), // modelsParse.SupplierJSON
		StopSend:      make(chan bool),
		StartSendProd: make(chan models.ProductsSuppliers),
		StopSendProd:  make(chan bool),
		//New:           new,
	}
}
func (pool *WorkerPool) Stop() {
	for i := 0; i < pool.Count; i++ {
		pool.StopSend <- false
	}
}

func (pool *WorkerPool) StopParsePrice() {
	for i := 0; i < pool.Count; i++ {
		pool.StopSend <- false
	}
}

func (pool *WorkerPool) Start(wg *sync.WaitGroup, goNum int, conn *sql.DB, TX *sql.Tx, logger *logger.Logger) {
	//var wg *sync.WaitGroup
	var supp models.SupplierForParse
	defer wg.Done()
	for {
		select {
		case supp = <-pool.StartSendData:
			parser.ParseFromAPI(supp, goNum, conn, TX, logger)
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
func (pool *WorkerPool) StartParsePrice(wg *sync.WaitGroup, goNum int, conn *sql.DB, TX *sql.Tx, logger *logger.Logger) {
	var prod models.ProductsSuppliers
	defer wg.Done()
	for {
		select {
		case prod = <-pool.StartSendProd:
			_ = parser.ParsePriceToDB(prod.Price, prod.ExternalProductID, prod.ExternalSupplierID+1, goNum, conn, TX, logger)

		case <-pool.StopSend:
			return
		}

	}
}
