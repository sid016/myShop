package models

type Order struct {
	ID         string `json:"id" db:"id"`
	ProductID  string `json:"product_id" db:"product_id"` // Foreign key
	BuyerName  string `json:"buyer_name" db:"buyer_name"`
	BuyerEmail string `json:"buyer_email" db:"buyer_email"`
	Quantity   int    `json:"quantity" db:"quantity"`
	CreatedAt  string `json:"created_at" db:"created_at"` // Date format e.g., YYYY-MM-DD HH:MM:SS
}
