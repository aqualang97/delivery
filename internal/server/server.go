package server

import (
	"database/sql"
	config "delivery/configs"
	handProv "delivery/internal/auth/handle_provide"
	rp "delivery/internal/repositories/database"
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

	handlerProvider := &handProv.HandlerProvider{
		UserRepository:             rp.NewUserRepo(conn, TX),
		UserAccessTokenRepository:  rp.NewAccessTokenRepo(conn, TX),
		UserRefreshTokenRepository: rp.NewRefreshTokenRepo(conn, TX),
		Config:                     s.cfg,
	}

	repProvider := &repositories_provider.RepositoriesProvider{
		IngredientRepository:          rp.NewIngredientRepo(conn, TX, s.cfg.Logger),
		OrderProductRepository:        rp.NewOrderProductRepo(conn, TX),
		OrderRepository:               rp.NewOrderRepo(conn, TX),
		ProductRepository:             rp.NewProductRepo(conn, TX),
		ProductsCategoriesRepository:  rp.NewProductsCategoriesRepo(conn, TX),
		ProductsIngredientsRepository: rp.NewProductsIngredientsRepo(conn, TX),
		ProductsSuppliersRepository:   rp.NewProductsSuppliersRepo(conn, TX),
		SupplierRepository:            rp.NewSupplierRepo(conn, TX),
		SuppliersCategoriesRepository: rp.NewSuppliersCategoriesRepo(conn, TX),
		UserContactRepository:         rp.NewUserContactRepo(conn, TX),
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
