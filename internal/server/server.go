package server

import (
	"database/sql"
	config "delivery/configs"
	handProv "delivery/internal/auth/handle_provide"
	rp "delivery/internal/repositories"
	"delivery/internal/repositories_provider"
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
	handlerProvider := &handProv.HandlerProvider{
		UserRepository:             rp.NewUserRepo(conn, TX, myLogger),
		UserAccessTokenRepository:  rp.NewAccessTokenRepo(conn, TX, myLogger),
		UserRefreshTokenRepository: rp.NewRefreshTokenRepo(conn, TX, myLogger),
		Config:                     s.cfg,
		Logger:                     myLogger,
	}

	repProvider := &repositories_provider.RepositoriesProvider{
		IngredientRepository:          rp.NewIngredientRepo(conn, TX, myLogger),
		OrderProductRepository:        rp.NewOrderProductRepo(conn, TX, myLogger),
		OrderRepository:               rp.NewOrderRepo(conn, TX, myLogger),
		ProductRepository:             rp.NewProductRepo(conn, TX, myLogger),
		ProductsCategoriesRepository:  rp.NewProductsCategoriesRepo(conn, TX, myLogger),
		ProductsIngredientsRepository: rp.NewProductsIngredientsRepo(conn, TX, myLogger),
		ProductsSuppliersRepository:   rp.NewProductsSuppliersRepo(conn, TX, myLogger),
		SupplierRepository:            rp.NewSupplierRepo(conn, TX, myLogger),
		SuppliersCategoriesRepository: rp.NewSuppliersCategoriesRepo(conn, TX, myLogger),
		UserContactRepository:         rp.NewUserContactRepo(conn, TX, myLogger),
		Logger:                        myLogger,
	}

	//m := middlware.NewMiddleware(handlerProvider)
	mux := http.NewServeMux()

	handlerProvider.UserAccessTokenRepository.DeleteNaturallyExpiredAccessToken()
	handlerProvider.UserRefreshTokenRepository.DeleteNaturallyExpiredRefreshToken()

	route.Router(*handlerProvider, *repProvider, mux)

	s.cfg.Logger.InfoConsole("Start listen...")
	log.Fatal(http.ListenAndServe(s.cfg.Port, mux)) //слушаем порт 8080 для входящих запросов

	return err

}
