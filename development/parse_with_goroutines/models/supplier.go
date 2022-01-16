package models

import (
	"encoding/json"
	"os"
)

type SupplierJSON struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	Image        string       `json:"image"`
	WorkingHours WorkingHours `json:"workingHours"`
	Menu         Menu         `json:"menu"`
}

type WorkingHours struct {
	Opening string `json:"opening"`
	Closing string `json:"closing"`
}

type Menu struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
}

func ReadFromJSON(filepath string) (Supplier, error) {
	var supp Supplier
	openFile, err := os.Open(filepath)
	if err != nil {
		return supp, err
	}
	err = json.NewDecoder(openFile).Decode(&supp)
	return supp, err

}
