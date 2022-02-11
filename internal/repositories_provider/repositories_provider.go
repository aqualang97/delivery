package repositories_provider

import (
	r "delivery/internal/repositories/database"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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

//list of all suppliers

func (rp RepositoriesProvider) Suppliers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		listOfSupp, _ := rp.SupplierRepository.GetAllSuppliers()
		for _, supp := range listOfSupp {
			fmt.Fprintf(w, "%d.\t Supplier: %s; Openning: %s, Closing: %s, Logo: %s \n",
				supp.ID, supp.Name, supp.WorkingHours.Opening, supp.WorkingHours.Closing, supp.Image)
		}
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)

	}
}

// specific suppliers

func (rp RepositoriesProvider) IndividualSupplier(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		//supp := rp.SupplierRepository.GetSupplierByID()
		path := r.URL.Path
		parts := strings.Split(path, "/suppliers/")
		if len(parts) != 2 {
			return
		}
		if id, err := strconv.Atoi(parts[1]); err == nil {
			supp, err := rp.SupplierRepository.GetSupplierByID(id)
			if err != nil {
				log.Println(err)
			}
			fmt.Fprint(w, supp.Name, supp.Image)

		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)

	}
}

//list of categories

func (rp RepositoriesProvider) Categories(w http.ResponseWriter, r *http.Request) {

}

//ListOfProductsInSpecificCategory

func (rp RepositoriesProvider) ListOfProductsInSpecificCategory(w http.ResponseWriter, r *http.Request) {

}

// list of products of specific supplier

func (rp RepositoriesProvider) AllProductsOfSupplier(w http.ResponseWriter, r *http.Request) {

}

// specific product of specific supplier

func (rp RepositoriesProvider) IndividualProduct(w http.ResponseWriter, r *http.Request) {

}

//Ingredients of product

func (rp RepositoriesProvider) IngredientsOfProduct(w http.ResponseWriter, r *http.Request) {

}
