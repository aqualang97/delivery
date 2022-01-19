package models

import (
	"time"
)

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    *time.Time
}

//type UserToken struct {
//	UserID           int
//	AccessToken      string `json:"access_token"`
//	RefreshToken     string `json:"refresh_token"`
//	AccessExpiredAt  time.Time
//	RefreshExpiredAt time.Time
//	Expired          string
//}
