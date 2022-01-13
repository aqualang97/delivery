package models

import "time"

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

type UserToken struct {
	ID               int
	UserID           int
	AccessToken      string
	RefreshToken     string
	AccessExpiredAt  time.Time
	RefreshExpiredAt time.Time
	Expired          string
}
