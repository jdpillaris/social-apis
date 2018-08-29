package relationship

import (
	"app/models"
	"time"

	"github.com/apex/log"
)

// CreateRelationship service
type CreateRelationship struct {
	// person1_ID   int64
	// person2_ID   int64
	email1   	 string
	email2   	 string
	is_friends   bool
	follows		 bool
	relationship *models.Relationship
}

// CreateRelationship instance
func NewRelationship(email1, email2 string, is_friends, follows bool) *CreateRelationship {
	n := new(CreateRelationship)
	n.email1 = email1
	n.email2 = email2
	n.is_friends = is_friends
	n.follows = follows
	n.relationship = new(models.Relationship)
	return n
}

// Do will validate data and create a relationship
func (r *CreateRelationship) Do() (err error) {
	if err = r.validate(); err != nil {
		return err
	}

	if err = r.create(); err != nil {
		return err
	}

	return nil
}

func (r *CreateRelationship) validate() (err error) {
	return nil
}

// Create a record
func (r *CreateRelationship) create() (err error) {
	relationship := new(models.Relationship)
	relationship.Person1 = 1
	relationship.Person2 = 2
	relationship.IsFriends = r.is_friends
	relationship.Follows = r.follows
	relationship.CreatedAt = time.Now().UTC()
	r.relationship, err = relationship.Create()
	if err != nil {
		log.WithError(err).Error("Unable to add relationship")
		return err
	}

	return nil
}
