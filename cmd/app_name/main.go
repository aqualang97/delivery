package cmd

import (
	"database/sql"
	"delivery/internal/models"
	"delivery/internal/repositories/filesystem"
	"log"
	"time"
)

func main() {
	//db.NewUserRepo()
	repo := filesystem.UserFileRepository{}
	user1 := models.User{Email: "a@a", PasswordHash: "21321", CreatedAt: "123"}
	err := repo.Insert(&user1)
	if err != nil {
		log.Fatal(err)
	}
}

var workTime time.Time
var nulTime = sql.NullTime{}
