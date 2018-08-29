package models

import (
	"fmt"
	"time"

	"github.com/apex/log"
)

// Relationship model
type Relationship struct {
	ID         int64     `json:"id"`
	Person1    int64     `json:"person_1"`
	Person2    int64     `json:"person_2"`
	IsFriends  bool    	 `json:"is_friends"`
	Follows    bool    	 `json:"follows"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Create a relationship
func (r *Relationship) Create() (*Relationship, error) {
	fmt.Println(r.CreatedAt)
	stmt, err := GetDB().Prepare(`
        INSERT INTO Relationships(person_1, person_2, is_friends, follows, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.WithError(err).Error("Failed to insert/update person")
		return r, err
	}

	row, err := stmt.Exec(
		r.Person1,
		r.Person2,
		r.IsFriends,
		r.Follows,
		r.CreatedAt,
		r.CreatedAt,
	)
	if err != nil {
		log.WithError(err).Error("Failed to insert/update person")
		return r, err
	}

	rowID, _ := row.LastInsertId()
	r.ID = rowID

	return r, nil
}

// Update relationship
func (r *Relationship) Update(id int64) (*Relationship, error) {
	return r, nil
}

// GetAll friends
func (r *Relationship) GetAllFriends(UID int64) ([]Relationship, error) {
	rs := make([]Relationship, 0)

	return rs, nil
}

// GetAll followers
func (r *Relationship) GetAllFollowers(UID int64) ([]Relationship, error) {
	rs := make([]Relationship, 0)

	return rs, nil
}

// GetAll followers
func (r *Relationship) GetAll(UID int64) ([]Relationship, error) {
	rs := make([]Relationship, 0)

	return rs, nil
}

// Get information on a specific relationship based on ID
func (r *Relationship) Get(id int64) (Relationship, error) {
	var relation Relationship

	return relation, nil
}

// Delete a specific relationship based on person_1, person_2 IDs
func (r *Relationship) Delete(person_1, person_2 int64) error {
	return nil
}
