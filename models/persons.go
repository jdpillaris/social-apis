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
}

// Create a person
func (p *Person) Create() (*Person, error) {
	fmt.Println(p.CreatedAt)
	stmt, err := GetDB().Prepare(`
        INSERT IGNORE INTO Persons(email, created_at)
        VALUES (?, ?)
	`)
	if err != nil {
		log.WithError(err).Error("Failed to create person")
		return p, err
	}

	row, err := stmt.Exec(
		p.Email,
		p.CreatedAt,
	)
	if err != nil {
		log.WithError(err).Error("Failed to insert person")
		return p, err
	}

	rowID, _ := row.LastInsertId()
	p.ID = rowID

	return p, nil
}
