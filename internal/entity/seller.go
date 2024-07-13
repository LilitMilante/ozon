package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Seller struct {
	ID                int64     `json:"id"`
	FullName          string    `json:"full_name"`
	Login             string    `json:"login"`
	Password          string    `json:"password,omitempty"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (p *Seller) Sanitize() {
	p.Password = ""
}

func (p *Seller) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p.EncryptedPassword), []byte(password)) == nil
}
