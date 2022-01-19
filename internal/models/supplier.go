package models

import "time"

type Suppliers struct {
	Id                 int
	Name               string
	CategoryOfSupplier int

	//StartOfWork        *time.Time
	//EndOfWork          *time.Time
	StartOfWork string
	EndOfWork   string
	Image       string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
