package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/apex/log"
)

// Relationship model
type Relationship struct {
	ID         int64     `json:"id"`
	Person1    int64     `json:"person_1"`
	Person2    int64     `json:"person_2"`
	Email1     string    `json:"email_1"`
	Email2     string    `json:"email_2"`
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
func (r *Relationship) GetAllFriends(email string) ([]Relationship, error) {
	rs := make([]Relationship, 0)
	rows, err := GetDB().Query(
		`SELECT r.id, r.person_1, r.person_2, p1.email, p2.email,
		r.is_friends, r.follows, r.created_at, r.updated_at
		FROM Relationships AS r
		JOIN Persons AS p2
		JOIN Persons AS p1
		ON r.person_2 = p2.id
		ON r.person_1 = p1.id
		WHERE p2.email = ? AND
		r.is_friends = 1`,
		email,
	)
	if err != nil {
		log.WithError(err).Error("Failed to fetch friends")
		return rs, err
	}
	defer rows.Close()
	for rows.Next() {
		var relationship Relationship
		err := rows.Scan(
			&relationship.ID,
			&relationship.Person1,
			&relationship.Person2,
			&relationship.Email1,
			&relationship.Email2,
			&relationship.IsFriends,
			&relationship.Follows,
			&relationship.CreatedAt,
			&relationship.UpdatedAt,
		)
		if err != nil {
			log.WithError(err).Error("Failed to get friendships by email")
			return rs, err
		}
		rs = append(rs, relationship)
	}
	err = rows.Err()
	if err != nil && err != sql.ErrNoRows {
		log.WithError(err).Error("Friendships by email row error")
		return rs, err
	}

	return rs, nil
}

// GetAll followers
func (r *Relationship) GetAllFollowers(email string) ([]Relationship, error) {
	rs := make([]Relationship, 0)
	rows, err := GetDB().Query(
		`SELECT r.id, r.person_1, r.person_2, p1.email, p2.email,
		r.is_friends, r.follows, r.created_at, r.updated_at
		FROM Relationships AS r
		JOIN Persons AS p2
		JOIN Persons AS p1
		ON r.person_2 = p2.id
		ON r.person_1 = p1.id
		WHERE p2.email = ? AND
		r.follows = 1`,
		email,
	)
	if err != nil {
		log.WithError(err).Error("Failed to fetch followers")
		return rs, err
	}
	defer rows.Close()
	for rows.Next() {
		var relationship Relationship
		err := rows.Scan(
			&relationship.ID,
			&relationship.Person1,
			&relationship.Person2,
			&relationship.Email1,
			&relationship.Email2,
			&relationship.IsFriends,
			&relationship.Follows,
			&relationship.CreatedAt,
			&relationship.UpdatedAt,
		)
		if err != nil {
			log.WithError(err).Error("Failed to get followers by email")
			return rs, err
		}
		rs = append(rs, relationship)
	}
	err = rows.Err()
	if err != nil && err != sql.ErrNoRows {
		log.WithError(err).Error("Followers by email row error")
		return rs, err
	}

	return rs, nil
}

// GetMutualFriends friends
func (r *Relationship) GetMutualFriends(email1, email2 string) ([]Relationship, error) {
	rs := make([]Relationship, 0)
	rows, err := GetDB().Query(
		`
		SELECT r.id, r.person_1, r.person_2, p1.email, p2.email,
		r.is_friends, r.follows, r.created_at, r.updated_at
		FROM Relationships AS r
		JOIN Persons AS p2
		JOIN Persons AS p1
		ON r.person_2 = p2.id
		ON r.person_1 = p1.id
		WHERE p2.email = ? AND
		r.is_friends = 1
		UNION ALL
		SELECT r.id, r.person_1, r.person_2, p1.email, p2.email,
		r.is_friends, r.follows, r.created_at, r.updated_at
		FROM Relationships AS r
		JOIN Persons AS p2
		JOIN Persons AS p1
		ON r.person_2 = p2.id
		ON r.person_1 = p1.id
		WHERE p2.email = ? AND
		r.is_friends = 1
		`,
		email1,
		email2,
	)
	if err != nil {
		log.WithError(err).Error("Failed to fetch mutual friends")
		return rs, err
	}
	defer rows.Close()
	for rows.Next() {
		var relationship Relationship
		err := rows.Scan(
			&relationship.ID,
			&relationship.Person1,
			&relationship.Person2,
			&relationship.Email1,
			&relationship.Email2,
			&relationship.IsFriends,
			&relationship.Follows,
			&relationship.CreatedAt,
			&relationship.UpdatedAt,
		)
		if err != nil {
			log.WithError(err).Error("Failed to get mutual friends")
			return rs, err
		}
		rs = append(rs, relationship)
	}
	err = rows.Err()
	if err != nil && err != sql.ErrNoRows {
		log.WithError(err).Error("Mutual friends row error")
		return rs, err
	}

	return rs, nil
}

// GetAll relationships
func (r *Relationship) GetAll(emailID string) ([]Relationship, error) {
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
