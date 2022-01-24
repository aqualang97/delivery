package models

import "time"

type UserAccessToken struct {
	ID          int
	UserID      int
	AccessToken string `json:"access_token"`
	CreatedAt   *time.Time
	ExpiredAt   *time.Time
	Expired     string
}

//type RequestAccessToken struct {
//	AccessToken string `json:"access_token"`
//}
