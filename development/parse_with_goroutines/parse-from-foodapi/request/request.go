package request

import (
	"delivery/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FoodAPIReq struct {
}

func GetSuppliers(url string) *models.AllSuppliers {

	//resp, err := http.Get(url)

	resp, err := http.NewRequest("Get", url, nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	var supp *models.AllSuppliers
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &supp)
	if err != nil {
		log.Println(err)
	}
	println("supp yes")
	return supp
	//for _, s := range supp.Suppliers {
	//	fmt.Println(s)
	//}
}

func GetMenuWithSuppID(url, endpointMenu, endpointSupp string, i int) *models.AllMenu {
	resp, err := http.Get(fmt.Sprintf("%s%s/%d%s", url, endpointSupp, i, endpointMenu))
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	var menu *models.AllMenu
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &menu)
	if err != nil {
		log.Println(err)
	}
	println("menu yes")
	return menu
}
