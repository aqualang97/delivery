package repositories_provider

import (
	r "delivery/internal/repositories/database"
	"net/http"
)

type RepositoriesProvider struct {
	IngredientRepository          *r.IngredientRepository
	OrderProductRepository        *r.OrderProductDBRepository
	OrderRepository               *r.OrderDBRepository
	ProductRepository             *r.ProductDBRepository
	ProductsCategoriesRepository  *r.ProductsCategoriesRepo
	ProductsIngredientsRepository *r.ProductsIngredientsRepository
	ProductsSuppliersRepository   *r.ProductsSuppliersRepository
	SupplierRepository            *r.SupplierDBRepository
	SuppliersCategoriesRepository *r.SuppliersCategoriesRepository
	UserContactRepository         *r.UserContactRepository
}

func (rp RepositoriesProvider) Suppliers(w http.ResponseWriter, r *http.Request) {
	//list of all suppliers
}

func (rp RepositoriesProvider) IndividualSupplier(w http.ResponseWriter, r *http.Request) {
	//list of all suppliers
}
