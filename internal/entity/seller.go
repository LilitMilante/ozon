package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Seller struct {
	ID                uuid.UUID `json:"id"`
	FullName          string    `json:"full_name"`
	Login             string    `json:"login"`
	Password          string    `json:"password,omitempty"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (s *Seller) Sanitize() {
	s.Password = ""
}

func (s *Seller) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(s.EncryptedPassword), []byte(password)) == nil
}
