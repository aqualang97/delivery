package models

import "time"

type Suppliers struct {
	Id                 int
	Name               string
	CategoryOfSupplier int
	StartOfWork        time.Time
	EndOfWork          time.Time
	CreatedAt          time.Time
}
