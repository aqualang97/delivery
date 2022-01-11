package models

import "time"

type User struct {
	Id           int
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
