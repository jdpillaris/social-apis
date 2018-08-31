package models

import (
	"fmt"
	"time"

	"github.com/apex/log"
)

// Person model
type Person struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Create a person
func (p *Person) Create() (*Person, error) {
	fmt.Println(p.CreatedAt)
	stmt, err := GetDB().Prepare(`
        INSERT INTO Persons(email, created_at, updated_at)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE updated_at = ?
	`)
	if err != nil {
		log.WithError(err).Error("Failed to create person")
		return p, err
	}

	row, err := stmt.Exec(
		p.Email,
		p.CreatedAt,
		p.UpdatedAt,
		p.UpdatedAt,
	)
	if err != nil {
		log.WithError(err).Error("Failed to insert person")
		return p, err
	}

	rowID, _ := row.LastInsertId()
	p.ID = rowID

	return p, nil
}
