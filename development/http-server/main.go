package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/order", Order)
	http.HandleFunc("/menu", Menu)
	http.HandleFunc("/users", Users)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Order(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Order Page")
	if err != nil {
		return
	}
}
func Menu(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Manu Page")
	if err != nil {
		return
	}
}
func Users(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Users Page")
	if err != nil {
		return
	}
}
