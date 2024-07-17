package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Seller struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Login     string    `json:"login"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Seller) Sanitize() {
	s.Password = ""
}

func (s *Seller) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(password)) == nil
}

func (s *Seller) Validate() error {
	const minLen = 3
	const maxLen = 50

	if s.FullName == "" || len([]rune(s.FullName)) > maxLen {
		return fmt.Errorf("the full name must not be empty and must be between 1 and 50 characters long")
	}
	if s.Login == "" || len([]rune(s.Login)) < minLen || len([]rune(s.Login)) > maxLen {
		return fmt.Errorf("the login must not be empty and must be between 3 and 50 characters long")
	}
	if s.Password == "" {
		return fmt.Errorf("empty password")
	}

	return nil
}
