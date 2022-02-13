package repositories_provider

import (
	r "delivery/internal/repositories"
	"encoding/json"
	"fmt"
	"github.com/aqualang97/logger/v4"
	"log"
	"net/http"
	"sort"
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

// specific suppliers
// /suppliers/?/products/?
// specific product of specific supplier

func (rp RepositoriesProvider) SupplierAndProdWithID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//supp := rp.SupplierRepository.GetSupplierByID()
		path := r.URL.Path
		parts := strings.Split(path, "/")
		if len(parts) >= 3 && len(parts) <= 6 {
			strSuppID := parts[2]
			if suppID, err := strconv.Atoi(strSuppID); err == nil {
				switch len(parts) {
				case 3:
					supp, err := rp.SupplierRepository.GetSupplierByID(suppID)
					if err != nil {
						log.Println(err)
						rp.ErrorHandler(w, r, 404)
					}
					fmt.Fprint(w, supp.Name, supp.Image)
				case 4:
					// list of products of specific supplier
					// /suppliers/?/products
					if parts[3] != "products" {
						return
					}
					listOfProd := rp.ProductRepository.GetListOfProdBySupplier(suppID)

					for _, p := range listOfProd {
						fmt.Fprint(w, p.ID, "	", p.Name, p.Category, "\n")
					}
				case 5:
					if parts[3] != "products" {
						rp.ErrorHandler(w, r, 404)
						return
					}
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
							return
						}
						fmt.Fprint(w, prod.Name)
					}
				case 6:
					if parts[3] != "products" || parts[5] != "ingredients" {
						rp.ErrorHandler(w, r, 404)
						return
					}
					strProdID := parts[4]
					if prodID, err := strconv.Atoi(strProdID); err == nil {
						exist := rp.ProductsSuppliersRepository.IsExist(prodID, suppID)
						if !exist {
							rp.ErrorHandler(w, r, 404)
							return
						}
						listPrIn, _ := rp.ProductsIngredientsRepository.GetIngredientsByProductID(prodID)
						for _, ing := range listPrIn {
							ingredient, _ := rp.IngredientRepository.GetIngredientByID(ing.IngredientID)
							fmt.Fprint(w, ingredient)
						}
					}
				}
			} else {
				rp.ErrorHandler(w, r, 404)
				return
			}
		} else {
			rp.ErrorHandler(w, r, 404)
			return
		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}

//list of categories

func (rp RepositoriesProvider) Categories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		categories, err := rp.ProductsCategoriesRepository.GetAllCategories()
		sort.Slice(categories, func(i, j int) bool {

			return categories[i].ID < categories[j].ID
		})
		if err != nil {
			rp.Logger.Error("Handler Categories\n", err)
			rp.ErrorHandler(w, r, 404)
		}
		for _, c := range categories {

			fmt.Fprint(w, c.ID, "	", c.Name, "\n")

		}
	default:
		{
			http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		}
	}
}

//ListOfProductsInSpecificCategory
// /categories/id
func (rp RepositoriesProvider) ListOfProductsInSpecificCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		path := r.URL.Path
		parts := strings.Split(path, "/")
		if len(parts) == 3 {
			strCatID := parts[2]
			if categoryID, err := strconv.Atoi(strCatID); err == nil {
				listOfProd := rp.ProductRepository.GetListOfProdInCategory(categoryID)
				println(len(listOfProd))
				for _, prod := range listOfProd {
					fmt.Fprint(w, prod.ID, "	", prod.Name, "\n")
				}
			}
		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}

}

func (rp RepositoriesProvider) ListOfAllProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		listAllProducts := rp.ProductRepository.GetAllProducts()
		data, _ := json.Marshal(listAllProducts)
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write(data)

	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}

//
//func (rp RepositoriesProvider) IndividualProduct(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "GET":
//
//		////supp := rp.SupplierRepository.GetSupplierByID()
//		//path := r.URL.Path
//		//parts := strings.Split(path, "/suppliers/")
//		//if len(parts) != 2 {
//		//	return
//		//}
//		//if id, err := strconv.Atoi(parts[1]); err == nil {
//		//	supp, err := rp.SupplierRepository.GetSupplierByID(id)
//		//	if err != nil {
//		//		log.Println(err)
//		//	}
//		//	fmt.Fprint(w, supp.Name, supp.Image)
//		//
//		//}
//
//		fmt.Fprintf(w, r.URL.Path)
//	default:
//		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
//
//	}
//}

//Ingredients of product
//
//func (rp RepositoriesProvider) IngredientsOfProduct(w http.ResponseWriter, r *http.Request) {
//
//}

func (rp RepositoriesProvider) ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Error\n Page not found")
	}
}
