package models

import "time"

type Supplier struct {
	Id                 int
	Name               string
	CategoryOfSupplier int

	//StartOfWork        *time.Time
	//EndOfWork          *time.Time
	WorkingHours WorkingHours
	Image        string
	ExternalID   int
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
type WorkingHours struct {
	Opening string `json:"opening"`
	Closing string `json:"closing"`
}
