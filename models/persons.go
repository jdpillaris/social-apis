package models

import (
	"time"
)

// Person model
type Person struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
}

// Create a person
func (p *Person) Create() (*Person, error) {
	return p, nil
}
