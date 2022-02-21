package request

import (
	"context"
	"delivery/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type FoodAPIReq struct {
}

const url = "http://foodapi.true-tech.php.nixdev.co"
const endpointSupp = "/suppliers"
const endpointMenu = "/menu"

func ReqWithCont(url string) *http.Response {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	return resp
}

func GetSuppliers() *models.AllSuppliersForParse {

	//resp, err := http.Get(url)
	//инициализайия контекста с таймаутом
	resp := ReqWithCont(url + endpointSupp)
	defer resp.Body.Close()
	var supp *models.AllSuppliersForParse
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

func GetMenuWithSuppID(shopID int) *models.AllMenu {

	resp := ReqWithCont(fmt.Sprintf("%s%s/%d%s", url, endpointSupp, shopID, endpointMenu))
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

func GetProductFromAPI(suppID, productID int) *models.Position {
	//it's externalID
	resp := ReqWithCont(fmt.Sprintf("%s%s/%d%s/%d", url, endpointSupp, suppID, endpointMenu, productID))
	var product *models.Position
	//fmt.Printf("%s%s/%d%s/%d", url, endpointSupp, suppID, endpointMenu, productID)
	//time.Sleep(1 * time.Millisecond)
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Println(err)
	}
	return product
}
