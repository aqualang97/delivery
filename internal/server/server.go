package server

import (
	"database/sql"
	config "delivery/configs"
	"delivery/internal/controllers"
	"delivery/internal/router"
	"fmt"
	"log"
	"net/http"
	"os"
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

	controller := controllers.NewController(s.cfg, myLogger, conn, TX)
	mux := http.NewServeMux()

	//mux.HandleFunc("/profile", muxCORS)
	router.Router(controller, mux)
	s.cfg.Logger.InfoConsole("Start listen...")
	log.Fatal(http.ListenAndServe(s.cfg.Port, mux)) //слушаем порт 8080 для входящих запросов

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
