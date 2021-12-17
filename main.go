package main

import (
	"delivery/repositories/filesystem"
	"fmt"
)

func main() {
	repo := filesystem.UserFileRepository{}
	user := repo.GetByEmail("")

	fmt.Println(user)
}
