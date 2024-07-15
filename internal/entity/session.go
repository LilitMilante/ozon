package entity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID
	SellerID  uuid.UUID
	CreatedAt time.Time
	ExpiredAt time.Time
}
