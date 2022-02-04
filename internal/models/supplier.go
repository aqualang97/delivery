package models

import "time"

type Supplier struct {
	ID                 int
	Name               string `json:"name"`
	CategoryOfSupplier int    `json:"type"`

	//StartOfWork        *time.Time
	//EndOfWork          *time.Time
	WorkingHours WorkingHours `json:"workingHours"`
	Image        string       `json:"image"`
	ExternalID   int          `json:"id"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
type SupplierForParse struct {
	Name               string `json:"name"`
	CategoryOfSupplier string `json:"type"`

	//StartOfWork        *time.Time
	//EndOfWork          *time.Time
	WorkingHours WorkingHours `json:"workingHours"`
	Image        string       `json:"image"`
	ExternalID   int          `json:"id"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	Menu         []Position `json:"menu"`
}
type WorkingHours struct {
	Opening string `json:"opening"`
	Closing string `json:"closing"`
}

//type AllSuppliers struct {
//	Suppliers []Supplier
//}

type AllSuppliersForParse struct {
	Suppliers []SupplierForParse `json:"suppliers"`
}
