package controllers

import (
	"database/sql"
	config "delivery/configs"
	rp "delivery/internal/repositories"
	i "delivery/internal/repository_interfaces"
	"github.com/aqualang97/logger/v4"
)

type ConfigController struct {
	Config *config.Config
	Logger *logger.Logger
	ErrorsController
}
type Controllers struct {
	ConfigController ConfigController
	Auth             AuthController
	User             UserController
	Menu             MenuController
}

type AuthController struct {
	ConfigController           ConfigController
	UserRepository             i.UserRepositoryInterface
	UserAccessTokenRepository  i.UserAccessTokenRepositoryInterface
	UserRefreshTokenRepository i.UserRefreshTokenRepositoryInterface
}

type UserController struct {
	ConfigController           ConfigController
	UserRepository             i.UserRepositoryInterface
	UserAccessTokenRepository  i.UserAccessTokenRepositoryInterface
	UserRefreshTokenRepository i.UserRefreshTokenRepositoryInterface
	UserContactRepository      i.UserContactRepositoryInterface
	OrderRepository            i.OrderRepositoryInterface
	OrderProductRepository     i.OrderProductRepositoryInterface
}

type MenuController struct {
	ConfigController              ConfigController
	IngredientRepository          i.IngredientRepositoryInterface
	ProductRepository             i.ProductRepositoryInterface
	ProductsCategoriesRepository  i.ProductsCategoriesRepositoryInterface
	ProductsIngredientsRepository i.ProductsIngredientsRepositoryInterface
	ProductsSuppliersRepository   i.ProductsSuppliersRepositoryInterface
	SupplierRepository            i.SupplierRepositoryInterface
	SuppliersCategoriesRepository i.SuppliersCategoriesRepositoryInterface
}

func NewController(config *config.Config, myLogger *logger.Logger, conn *sql.DB, TX *sql.Tx) *Controllers {
	return &Controllers{
		ConfigController: ConfigController{Config: config, Logger: myLogger},
		Auth: AuthController{
			ConfigController:           ConfigController{Config: config, Logger: myLogger},
			UserRepository:             rp.NewUserRepo(conn, TX, myLogger),
			UserAccessTokenRepository:  rp.NewAccessTokenRepo(conn, TX, myLogger),
			UserRefreshTokenRepository: rp.NewRefreshTokenRepo(conn, TX, myLogger),
		},
		User: UserController{
			ConfigController:           ConfigController{Config: config, Logger: myLogger},
			UserRepository:             rp.NewUserRepo(conn, TX, myLogger),
			UserAccessTokenRepository:  rp.NewAccessTokenRepo(conn, TX, myLogger),
			UserRefreshTokenRepository: rp.NewRefreshTokenRepo(conn, TX, myLogger),
			UserContactRepository:      rp.NewUserContactRepo(conn, TX, myLogger),
			OrderRepository:            rp.NewOrderRepo(conn, TX, myLogger),
			OrderProductRepository:     rp.NewOrderProductRepo(conn, TX, myLogger),
		},
		Menu: MenuController{
			ConfigController:              ConfigController{Config: config, Logger: myLogger},
			IngredientRepository:          rp.NewIngredientRepo(conn, TX, myLogger),
			ProductRepository:             rp.NewProductRepo(conn, TX, myLogger),
			ProductsCategoriesRepository:  rp.NewProductsCategoriesRepo(conn, TX, myLogger),
			ProductsIngredientsRepository: rp.NewProductsIngredientsRepo(conn, TX, myLogger),
			ProductsSuppliersRepository:   rp.NewProductsSuppliersRepo(conn, TX, myLogger),
			SupplierRepository:            rp.NewSupplierRepo(conn, TX, myLogger),
			SuppliersCategoriesRepository: rp.NewSuppliersCategoriesRepo(conn, TX, myLogger),
		},
	}
}
