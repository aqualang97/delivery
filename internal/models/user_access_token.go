package models

import "time"

type UserAccessToken struct {
	UserID          int
	AccessToken     string `json:"access_token"`
	AccessExpiredAt *time.Time
	Expired         string
}
