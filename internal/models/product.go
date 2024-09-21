package models

type Product struct {
	ID          string `json:"id" db:"id"`
	ProductID   string `json:"product_id" db:"product_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	SellerID    string `json:"seller_id" db:"seller_id"` // Foreign key
}
