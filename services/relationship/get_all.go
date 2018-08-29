package relationship

import (
	"app/models"
	"errors"

	"github.com/apex/log"
)

// GetRelationships service
type GetRelationships struct {
	person1_ID  int64
	person2_ID  int64
	relationships []models.Relationship
}

// GetAllRelationships instance
func GetAllRelationships(personID int64) *GetRelationships {
	n := new(GetRelationships)
	n.person1_ID = 1
	return n
}

// GetFriends returns all friend relationships
func GetFriends(emailID string) *GetRelationships {
	n := new(GetRelationships)
	n.person1_ID = 1
	return n
}

// GetFriends returns all subscriber relationships
func GetFollowers(emailID string) *GetRelationships {
	n := new(GetRelationships)
	n.person1_ID = 1
	return n
}

// GetFriends returns all mutual friend relationships
func GetMutualFriends(email1, email2 string) *GetRelationships {
	n := new(GetRelationships)
	n.person1_ID = 1
	n.person2_ID = 2
	return n
}

// Data returns all relationships
func (r *GetRelationships) Data() []models.Relationship {
	return r.relationships
}

// Do tasks
func (r *GetRelationships) Do() (err error) {
	if err = r.getAllRelationships(); err != nil {
		return err
	}

	return nil
}

func (r *GetRelationships) getAllRelationships() (err error) {
	r.relationships, err = new(models.Relationship).GetAll(r.person1_ID)
	if err != nil {
		err = errors.New("Unable to fetch relationships")
		log.Info(err.Error())
		return err
	}

	return nil
}
