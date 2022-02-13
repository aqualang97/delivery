package server

import (
	"database/sql"
	config "delivery/configs"
	"delivery/internal/controllers"
	"delivery/internal/route"
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
	route.Router(controller, mux)

	s.cfg.Logger.InfoConsole("Start listen...")
	log.Fatal(http.ListenAndServe(s.cfg.Port, mux)) //слушаем порт 8080 для входящих запросов

	return err

}
