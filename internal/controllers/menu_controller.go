package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// Different action with suppliers, products, etc.
// ex: /all-products, /suppliers/id-supp/products/id-prod
func (rp MenuController) Home(w http.ResponseWriter, r *http.Request) {
	//http.Redirect(w, r, fmt.Sprintf("localhost:8080/home.html"), 301)

}
func (rp MenuController) Suppliers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		listOfSupp, _ := rp.SupplierRepository.GetAllSuppliers()
		//for _, supp := range listOfSupp {
		//	fmt.Fprintf(w, "%d.\t Supplier: %s; Openning: %s, Closing: %s, Logo: %s \n",
		//		supp.ID, supp.Name, supp.WorkingHours.Opening, supp.WorkingHours.Closing, supp.Image)
		//}
		data, _ := json.Marshal(listOfSupp)
		w.Write(data)

	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)

	}
}

// specific suppliers

func (rp MenuController) IndividualSupplier(w http.ResponseWriter, r *http.Request) {
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

func (rp MenuController) SupplierAndProdWithID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		//supp := rp.SupplierRepository.GetSupplierByID()
		path := r.URL.Path
		parts := strings.Split(path, "/")[1:]
		log.Println(len(parts))
		log.Println(parts)
		log.Println(parts)
		log.Println(parts)
		log.Println(parts)
		log.Println(parts)
		if len(parts) >= 3 && len(parts) <= 6 {
			strSuppID := parts[2]
			if suppID, err := strconv.Atoi(strSuppID); err == nil {
				switch len(parts) {
				case 3:
					supp, err := rp.SupplierRepository.GetSupplierByID(suppID)
					if err != nil {
						log.Println(err)
						rp.ConfigController.ErrorHandler(w, r, 404)
					}
					fmt.Println(supp)
					data, _ := json.Marshal(supp)
					w.Write(data)

				case 4:
					// list of products of specific supplier
					// /suppliers/?/products
					if parts[3] != "products" {
						return
					}
					listOfProd := rp.ProductRepository.GetListOfProdBySupplier(suppID)

					//for _, p := range listOfProd {
					//	fmt.Fprint(w, p.ID, "	", p.Name, p.Category, "\n")
					//}
					//fmt.Println(listOfProd)
					data, _ := json.Marshal(listOfProd)
					w.Write(data)

				case 5:
					if parts[3] != "products" {
						rp.ConfigController.ErrorHandler(w, r, 404)
						return
					}
					strProdID := parts[4]
					if prodID, err := strconv.Atoi(strProdID); err == nil {
						exist := rp.ProductsSuppliersRepository.IsExist(prodID, suppID)
						if !exist {
							rp.ConfigController.ErrorHandler(w, r, 404)
							return
						}
						prod, err := rp.ProductRepository.GetProductByID(prodID)
						if err != nil {
							log.Println(err)
							rp.ConfigController.ErrorHandler(w, r, 404)
							return
						}
						data, _ := json.Marshal(prod)
						w.Write(data)
						fmt.Println(prod)
						//fmt.Fprint(w, prod.Name)
					}
				case 6:
					if parts[3] != "products" || parts[5] != "ingredients" {
						rp.ConfigController.ErrorHandler(w, r, 404)
						return
					}
					strProdID := parts[4]
					if prodID, err := strconv.Atoi(strProdID); err == nil {
						exist := rp.ProductsSuppliersRepository.IsExist(prodID, suppID)
						if !exist {
							rp.ConfigController.ErrorHandler(w, r, 404)
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
				rp.ConfigController.ErrorHandler(w, r, 404)
				return
			}
		} else {
			rp.ConfigController.ErrorHandler(w, r, 404)
			return
		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}

//list of categories

func (rp MenuController) Categories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		logger := rp.ConfigController.Logger
		categories, err := rp.ProductsCategoriesRepository.GetAllCategories()
		sort.Slice(categories, func(i, j int) bool {

			return categories[i].ID < categories[j].ID
		})
		if err != nil {
			logger.Error("Handler Categories\n", err)

			rp.ConfigController.ErrorHandler(w, r, 404)

		}
		//for _, c := range categories {
		//
		//	fmt.Fprint(w, c.ID, "	", c.Name, "\n")
		//
		//}

		data, _ := json.Marshal(categories)
		fmt.Println(categories)
		w.Write(data)
	default:
		{
			http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		}
	}
}

//ListOfProductsInSpecificCategory
// /categories/id
func (rp MenuController) ListOfProductsInSpecificCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		path := r.URL.Path
		parts := strings.Split(path, "/")[1:]
		//log.Println(parts)
		if len(parts) == 3 {
			strCatID := parts[2]
			if categoryID, err := strconv.Atoi(strCatID); err == nil {
				listOfProd := rp.ProductRepository.GetListOfProdInCategory(categoryID)
				//fmt.Println(listOfProd)
				data, _ := json.Marshal(listOfProd)
				w.Write(data)
			}
		}
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}

}

func (rp MenuController) ListOfAllProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		listAllProducts := rp.ProductRepository.GetAllProducts()
		//fmt.Println(listAllProducts)
		data, _ := json.Marshal(listAllProducts)
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
