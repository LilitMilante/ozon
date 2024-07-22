package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	SellerID    uuid.UUID `json:"seller_id"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       int64     `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateProduct struct {
	SellerID  uuid.UUID `json:"seller_id"`
	ProductID int64     `json:"product_id"`
	UpdateProductFields
}

type UpdateProductFields struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Category    *string `json:"category"`
	Price       *int64  `json:"price"`
}
