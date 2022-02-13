package models

import "time"

type Ingredients struct {
	ID        int
	Name      string
	CreatedAt *time.Time
}
