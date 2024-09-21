package models

type Pricing struct {
	ID            string  `json:"id" db:"id"`
	ProductID     string  `json:"product_id" db:"product_id"` // Foreign key
	Price         float64 `json:"price" db:"price"`
	EffectiveDate string  `json:"effective_date" db:"effective_date"` // Date format e.g., YYYY-MM-DD
}
