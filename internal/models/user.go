package models

import (
	"time"
)

type User struct {
	ID           int
	Email        string
	Login        string
	PasswordHash string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegistrationRequest struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    int
	Email string
	Login string
}

type UserRequestPairTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponsePairTokens struct {
	UserID       int    `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//type UserToken struct {
//	UserID           int
//	AccessToken      string `json:"access_token"`
//	RefreshToken     string `json:"refresh_token"`
//	AccessExpiredAt  time.Time
//	RefreshExpiredAt time.Time
//	Expired          string
//}
