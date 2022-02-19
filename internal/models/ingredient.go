package models

import "time"

type Ingredients struct {
	ID        int    `json:"ID"`
	Name      string `json:"name"`
	CreatedAt *time.Time
}
