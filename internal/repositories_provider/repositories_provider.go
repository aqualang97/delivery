package repositories_provider

import (
	r "delivery/internal/repositories"
	"fmt"
	"github.com/aqualang97/logger/v4"
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
	Logger                        *logger.Logger
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

// /suppliers/?/products/?
func (rp RepositoriesProvider) SupplierAndProdWithID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//supp := rp.SupplierRepository.GetSupplierByID()
		path := r.URL.Path
		parts := strings.Split(path, "/")
		if len(parts) == 3 || len(parts) == 5 {
			strSuppID := parts[2]
			if suppID, err := strconv.Atoi(strSuppID); err == nil {

				if len(parts) == 3 {
					supp, err := rp.SupplierRepository.GetSupplierByID(suppID)
					if err != nil {
						log.Println(err)
						rp.ErrorHandler(w, r, 404)
					}
					fmt.Fprint(w, supp.Name, supp.Image)
				} else {
					strProdID := parts[4]
					if prodID, err := strconv.Atoi(strProdID); err == nil {
						exist := rp.ProductsSuppliersRepository.IsExist(prodID, suppID)
						if !exist {
							rp.ErrorHandler(w, r, 404)

							return
						}
						prod, err := rp.ProductRepository.GetProductByID(prodID)
						if err != nil {
							log.Println(err)
							rp.ErrorHandler(w, r, 404)
						}
						fmt.Fprint(w, prod.Name)

					} else {
						rp.ErrorHandler(w, r, 404)
					}
				}

			} else {
				rp.ErrorHandler(w, r, 404)
			}
		} else {
			rp.ErrorHandler(w, r, 404)
		}
		//} else {
		//	partsProd := strings.Split(path, "/products/")
		//	if len(partsProd) != 2 {
		//		return
		//	}
		//
		//	partsSupAndOther := strings.Split(partsProd[0], "/suppliers/")
		//	if len(partsSupAndOther) != 2 {
		//		return
		//	}
		//	suppID, err := strconv.Atoi(partsSupAndOther[1])
		//
		//	if err != nil {
		//		log.Println(err)
		//	}
		//	if prodID, err := strconv.Atoi(partsProd[1]); err == nil {
		//		exist := rp.ProductsSuppliersRepository.IsExist(prodID, suppID)
		//		if !exist {
		//			return
		//		}
		//
		//		prod, err := rp.ProductRepository.GetProductByID(prodID)
		//		if err != nil {
		//			log.Println(err)
		//		}
		//		fmt.Fprint(w, prod.Id, "	", prod.Name)
		//	}

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
	switch r.Method {
	case "GET":

		////supp := rp.SupplierRepository.GetSupplierByID()
		//path := r.URL.Path
		//parts := strings.Split(path, "/suppliers/")
		//if len(parts) != 2 {
		//	return
		//}
		//if id, err := strconv.Atoi(parts[1]); err == nil {
		//	supp, err := rp.SupplierRepository.GetSupplierByID(id)
		//	if err != nil {
		//		log.Println(err)
		//	}
		//	fmt.Fprint(w, supp.Name, supp.Image)
		//
		//}

		fmt.Fprintf(w, r.URL.Path)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)

	}
}

//Ingredients of product

func (rp RepositoriesProvider) IngredientsOfProduct(w http.ResponseWriter, r *http.Request) {

}

func (rp RepositoriesProvider) ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Error\n Page not found")
	}
}
