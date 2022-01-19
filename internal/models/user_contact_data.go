package models

import "time"

type UserContactData struct {
	ID          int
	UserID      int
	PhoneNumber string
	Address     string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
