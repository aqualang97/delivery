package models

import "time"

type UserRefreshToken struct {
	ID           int
	UserID       int
	RefreshToken string `json:"refresh_token"`
	CreatedAt    *time.Time
	ExpiredAt    *time.Time
	Expired      string
}

//type RequestRefreshToken struct {
//	RefreshToken string `json:"refresh_token"`
//}
