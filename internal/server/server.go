package server

import (
	"database/sql"
	config "delivery/configs"
	"delivery/internal/controllers"
	"delivery/internal/models"
	"delivery/internal/parser"
	"delivery/internal/parser/request"
	"delivery/internal/parser/worker_pool"
	"delivery/internal/router"
	"fmt"
	"github.com/aqualang97/logger/v4"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Server struct {
	cfg  *config.Config
	conn *sql.DB
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
	return &Server{cfg: cfg, conn: db}
}

func (s *Server) Start() error {

	conn := s.conn
	TX, err := conn.Begin()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	myLogger := s.cfg.Logger
	//argsFull := os.Args
	//if len(argsFull) == 2 {
	//	if argsFull[1] == "--create" {
	//		err = s.MakeTables(conn, myLogger)
	//		log.Println("START CREATE")
	//		if err != nil {
	//			log.Println(err)
	//			return err
	//		}
	//	}
	//}
	//log.Println(argsFull)

	controller := controllers.NewController(s.cfg, myLogger, conn, TX)
	mux := http.NewServeMux()

	//mux.HandleFunc("/profile", muxCORS)
	router.Router(controller, mux, conn, TX, myLogger)
	s.cfg.Logger.InfoConsole("Start listen...")
	log.Println(s.cfg.Port)

	go func() {
		log.Fatal(http.ListenAndServe(s.cfg.Port, mux)) //слушаем порт 8080 для входящих запросов
	}()
	err = s.StartParse(conn, TX, myLogger)
	if err != nil {
		myLogger.Error("Can't parse with err \n", err)
		myLogger.ErrorConsole("Can't parse with err \n", err)
	}
	return err
}

func (s Server) StartParse(conn *sql.DB, TX *sql.Tx, myLogger *logger.Logger) error {
	pool := worker_pool.NewPool(4)
	count := pool.Count
	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		go func(i int) {
			pool.Start(&wg, i, conn, TX, myLogger)

		}(i)
		go func(i int) {
			pool.StartParsePrice(&wg, i, conn, TX, myLogger)
		}(i)
		wg.Add(2)
	}
	allSupp, err := request.GetSuppliers()
	if err != nil {
		log.Println(err, "\n Can't parse supplier")
		return err
	}
	println(allSupp)
	time.Sleep(10 * time.Second)
	for i, s := range allSupp.Suppliers {
		println("shop", i)
		menu, err := request.GetMenuWithSuppID(i + 1)
		if err != nil {
			log.Println(err, "\n Can't parse menu")
			continue
		}
		s.Menu = menu.Menu
		pool.StartSendData <- s
	}
	for {
		time.Sleep(10 * time.Second)
		println("start upd price")
		for suppID, _ := range allSupp.Suppliers {
			listProdId := parser.ParseProdSuppByDB(suppID+1, conn, TX, myLogger)
			for _, prodID := range listProdId {
				//можно конечно и не делать запрос на GetProductFromAPI,
				// но раз он есть, можем походить конкретно по продукту

				// либо могу выташить все экстернал айди из бд и пройтись по ним
				position, err := request.GetProductFromAPI(suppID+1, prodID)
				if err != nil {
					log.Println(err, "\n Can't parse price")
					continue
				}
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
	return err
}

//func muxCORS(w http.ResponseWriter, r *http.Request) {
//	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token"
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
//	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
//	log.Println("pass")
//	log.Println("___")
//	log.Println(w.Header())
//	log.Println("___")
//
//}

func (s Server) MakeTables(conn *sql.DB, myLogger *logger.Logger) error {
	rows, err := conn.Query("SHOW TABLES")
	var table string
	var numTables = 0
	for rows.Next() {
		err := rows.Scan(&table)
		if err != nil {
			return err
		}
		numTables++
		//fmt.Println(table)
	}
	fmt.Println("NUM", numTables)

	if numTables != 13 {
		s.CreateTables(conn, myLogger)
		if err != nil {
			myLogger.Error("can't create tables: \n", err)
			return err
		}
	} else {
		log.Println("passed")
	}
	return err
}
func (s Server) CreateTables(conn *sql.DB, myLogger *logger.Logger) {

	//
	//create table ingredients
	//
	_, err := conn.Exec(
		"create table if not exists ingredients (id int auto_increment primary key, name varchar(50) not null, created_at timestamp default CURRENT_TIMESTAMP not null, constraint ingredients_name_uindex unique (name) ) auto_increment = 0")
	if err != nil {
		log.Println("Can't create table ingredients", err)
		myLogger.Error("Can't create table ingredients", err)
		myLogger.ErrorConsole("Can't create table ingredients", err)
	}
	//
	// create table products_categories
	//

	_, err = conn.Exec("create table if not exists products_categories ( id     int auto_increment primary key, name varchar(50) not null, created_at timestamp default CURRENT_TIMESTAMP null, updated_at timestamp null on update CURRENT_TIMESTAMP, constraint products_categories_id_uindex unique (id), constraint products_categories_name_uindex unique (name) ) auto_increment = 0")
	if err != nil {
		myLogger.Error("Can't create table products_categories", err)
	}
	myLogger.ErrorConsole("Can't create table products_categories", err)

	//
	// create table products
	//

	_, err = conn.Exec("create table if not exists products ( id int auto_increment primary key, name    varchar(100)     not null, category    int  not null, external_id int  not null, created_at  timestamp default CURRENT_TIMESTAMP null, updated_at  timestamp   null on update CURRENT_TIMESTAMP, constraint products_products_categories_id_fk foreign key (category) references products_categories (id) ) auto_increment = 0")
	if err != nil {
		myLogger.Error("Can't create table products", err)
		myLogger.ErrorConsole("Can't create table products", err)
	}
	//
	// create table products_ingredients
	//
	_, err = conn.Exec("create table if not exists products_ingredients ( product_id    int  not null, ingredient_id int  not null, created_at    timestamp default CURRENT_TIMESTAMP null, constraint products_ingredients_ingredients_id_fk foreign key (ingredient_id) references ingredients (id), constraint products_ingredients_products_id_fk foreign key (product_id) references products (id) )")
	if err != nil {
		myLogger.Error("Can't create table products_ingredients", err)
		myLogger.ErrorConsole("Can't create table products_ingredients", err)
	}
	//
	// create table suppliers_categories
	//
	_, err = conn.Exec("create table if not exists suppliers_categories ( id     int auto_increment primary key, name   varchar(100)     not null, created_at timestamp default CURRENT_TIMESTAMP null, updated_at timestamp   null on update CURRENT_TIMESTAMP, constraint suppliers_categories_name_uindex unique (name) ) auto_increment = 0")
	if err != nil {
		myLogger.Error("Can't create table suppliers_categories", err)
		myLogger.ErrorConsole("Can't create table suppliers_categories", err)
	}
	//
	// create table suppliers
	//
	_, err = conn.Exec("create table if not exists suppliers ( id     int auto_increment primary key, name   varchar(100)     not null, category_of_supplier int  not null, start_of_work    varchar(5)  not null, end_of_work varchar(5)  not null, image  text not null, external_id int  null, created_at  timestamp default CURRENT_TIMESTAMP null, updated_at  timestamp   null on update CURRENT_TIMESTAMP, constraint suppliers_id_uindex unique (id), constraint suppliers_suppliers_categories_id_fk foreign key (category_of_supplier) references suppliers_categories (id) ) auto_increment = 0")
	if err != nil {
		myLogger.Error("Can't create table suppliers", err)
		myLogger.ErrorConsole("Can't create table suppliers", err)
	}
	//
	// create table products_suppliers
	//
	_, err = conn.Exec("create table if not exists products_suppliers ( product_id  int  not null, supplier_id int  not null, external_product_id  int  null, external_supplier_id int  not null, price  float    not null, image  text not null, created_at  timestamp default CURRENT_TIMESTAMP null, updated_at  timestamp   null on update CURRENT_TIMESTAMP, constraint products_suppliers_products_id_fk foreign key (product_id) references products (id), constraint products_suppliers_suppliers_id_fk foreign key (supplier_id) references suppliers (id) )")
	if err != nil {
		myLogger.Error("Can't create table products_suppliers", err)
		myLogger.ErrorConsole("Can't create table products_suppliers", err)
	}
	//
	// create table users
	//
	_, err = conn.Exec("create table if not exists users ( id     int auto_increment primary key, login  varchar(100)     not null, email  varchar(100)     not null, password   text not null, created_at timestamp default CURRENT_TIMESTAMP null, updated_at timestamp   null on update CURRENT_TIMESTAMP ) auto_increment = 0")
	if err != nil {
		myLogger.Error("Can't create table users", err)
		myLogger.ErrorConsole("Can't create table users", err)
	}
	//
	// create table orders
	//
	_, err = conn.Exec("create table if not exists orders ( id    int auto_increment primary key, user_id    int not null, price float   not null, payment_method varchar(15)  null, status     enum ('empty', 'formatting', 'not paid', 'cash', 'in processing', 'paid', 'completed') not null, created_at timestamp default CURRENT_TIMESTAMP null, updated_at timestamp    null on update CURRENT_TIMESTAMP, constraint orders_users_id_fk foreign key (user_id) references users (id) ) auto_increment = 0")
	if err != nil {
		myLogger.Error("Can't create table orders", err)
		myLogger.ErrorConsole("Can't create table orders", err)
	}
	//
	// create table orders_products
	//
	_, err = conn.Exec("create table if not exists orders_products ( id   int auto_increment primary key, product_id     int  not null, order_id  int  not null, numbers_of_product int  not null, purchase_price float    not null, created_at     timestamp default CURRENT_TIMESTAMP null, updated_at     timestamp   null on update CURRENT_TIMESTAMP, constraint orders_products_orders_id_fk foreign key (order_id) references orders (id), constraint orders_products_products_id_fk foreign key (product_id) references products (id) ) auto_increment = 0")
	if err != nil {
		myLogger.Error("Can't create table orders_products", err)
		myLogger.ErrorConsole("Can't create table orders_products", err)
	}
	//
	// create table users_access_tokens
	//
	_, err = conn.Exec("create table if not exists users_access_tokens ( id     int auto_increment primary key, user_id    int  not null, token  text not null, created_at timestamp default CURRENT_TIMESTAMP not null, expired_at timestamp   not null, expired    enum ('true', 'false')     not null, constraint users_access_tokens_users_id_fk foreign key (user_id) references users (id) ) auto_increment = 0")
	if err != nil {
		myLogger.Error("Can't create table users_access_tokens", err)
		myLogger.ErrorConsole("Can't create table users_access_tokens", err)
	}
	//
	// create table users_contact_data
	//
	_, err = conn.Exec("create table if not exists users_contact_data ( id  int auto_increment primary key, user_id  int  not null, first_name   varchar(50) null, last_name    varchar(50) null, phone_number varchar(13) not null, address  text not null, created_at   timestamp default CURRENT_TIMESTAMP null, updated_at   timestamp   null on update CURRENT_TIMESTAMP, constraint users_contact_data_users_id_fk foreign key (user_id) references users (id) )")
	if err != nil {
		myLogger.Error("Can't create table users_contact_data", err)

		myLogger.ErrorConsole("Can't create table users_contact_data", err)
	}
	//
	// create table users_refresh_tokens
	//
	_, err = conn.Exec("create table if not exists users_refresh_tokens ( id     int auto_increment primary key, user_id    int  not null, token  text not null, created_at timestamp default CURRENT_TIMESTAMP not null, expired_at timestamp   not null, expired    enum ('true', 'false')     not null, constraint users_refresh_tokens_users_id_fk foreign key (user_id) references users (id)) auto_increment = 0")
	if err != nil {
		myLogger.Error("Can't create table users_refresh_tokens", err)
		myLogger.ErrorConsole("Can't create table users_refresh_tokens", err)
	}

}

func (s Server) TruncateTables(conn *sql.DB) {
	_, err := conn.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE ingredients")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE products")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE products_ingredients")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE products_categories")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE products_suppliers")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE suppliers")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE orders")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE orders_products")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE users")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE users_access_tokens")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE users_contact_data")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE users_refresh_tokens")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("TRUNCATE TABLE suppliers_categories;")
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		log.Println(err)
	}
	log.Println("Pass truncate")
}
