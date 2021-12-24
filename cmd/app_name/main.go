package main

import (
	"delivery/repositories/filesystem"
	"delivery/repositories/models"
	"log"
)

func main() {
	repo := filesystem.UserFileRepository{}
	user1 := models.User{Email: "a@a", PasswordHash: "21321", CreatedAt: "123"}
	err := repo.Insert(&user1)
	if err != nil {
		log.Fatal(err)
	}
}
