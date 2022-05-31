package models

import "time"

type UserContactData struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
