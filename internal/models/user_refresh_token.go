package models

import "time"

type UserRefreshToken struct {
	UserID           int
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiredAt *time.Time
	Expired          string
}
