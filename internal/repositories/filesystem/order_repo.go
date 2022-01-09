package filesystem

import (
	"delivery/internal/models"
	"os"
)

type OrderRepository struct {
}

func (or OrderRepository) Insert(order *models.Order) error {
	file, err := os.Open("./datastore/files/order_1.json")
	if err != nil {
		panic(err)
		return err
	}
	defer file.Close()

	return err

}
