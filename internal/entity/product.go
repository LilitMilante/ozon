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
