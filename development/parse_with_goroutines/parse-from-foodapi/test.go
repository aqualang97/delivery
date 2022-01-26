package main

import (
	"delivery/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://foodapi.true-tech.php.nixdev.co/suppliers")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	var supp *models.AllSuppliers
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &supp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(supp.Suppliers)

}
