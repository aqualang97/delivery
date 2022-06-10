package request

import (
	"context"
	"delivery/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type FoodAPIReq struct {
}

const url = "http://foodapi.golang.nixdev.co"
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
	if err != nil {
		log.Println(err)
	}
	return resp
}

func GetSuppliers() (*models.AllSuppliersForParse, error) {
	var supp *models.AllSuppliersForParse

	//resp, err := http.Get(url)
	//инициализайия контекста с таймаутом
	resp := ReqWithCont(url + endpointSupp)
	if resp.StatusCode != http.StatusOK {
		return supp, errors.New(fmt.Sprintf("server not response\n status: %d", resp.StatusCode))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &supp)
	if err != nil {
		log.Println(err)
	}
	println("supp yes")
	return supp, nil
	//for _, s := range supp.Suppliers {
	//	fmt.Println(s)
	//}
}

func GetMenuWithSuppID(shopID int) (*models.AllMenu, error) {
	var menu *models.AllMenu

	resp := ReqWithCont(fmt.Sprintf("%s%s/%d%s", url, endpointSupp, shopID, endpointMenu))

	if resp.StatusCode != http.StatusOK {
		return menu, errors.New(fmt.Sprintf("server not response\n status: %d", resp.StatusCode))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &menu)
	if err != nil {
		log.Println(err)
	}
	println("menu yes")
	return menu, nil
}

func GetProductFromAPI(suppID, productID int) (*models.Position, error) {
	//it's externalID
	resp := ReqWithCont(fmt.Sprintf("%s%s/%d%s/%d", url, endpointSupp, suppID, endpointMenu, productID))
	var product *models.Position
	//fmt.Println(resp)
	if resp == nil {
		return product, errors.New(fmt.Sprintf("Response is nul. Unknown err. Server not response"))

	}
	if resp.StatusCode != http.StatusOK {
		return product, errors.New(fmt.Sprintf("server not response\n status: %d", resp.StatusCode))
	}
	defer resp.Body.Close()
	//fmt.Printf("%s%s/%d%s/%d", url, endpointSupp, suppID, endpointMenu, productID)
	//time.Sleep(1 * time.Millisecond)
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Println(err)
	}
	return product, nil
}
