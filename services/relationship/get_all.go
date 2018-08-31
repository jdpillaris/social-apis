package relationship

import (
	"app/models"
	"errors"

	"github.com/apex/log"
)

// GetRelationships service
type GetRelationships struct {
	// person1_ID   int64
	// person2_ID   int64
	email1   	 string
	email2   	 string
	relationships []models.Relationship
}

// GetAllRelationships instance
func GetAllRelationships(email2 string) *GetRelationships {
	n := new(GetRelationships)
	n.email2 = email2
	return n
}

// GetFriends returns all friend relationships
func (r *GetRelationships) GetFriends(emailID string) (err error) {
	r.relationships, err = new(models.Relationship).GetAllFriends(emailID)
	if err != nil {
		err = errors.New("Unable to fetch friends")
		log.Info(err.Error())
		return err
	}

	return nil
}

// GetFriends returns all subscriber relationships
func (r *GetRelationships) GetFollowers(emailID string) (err error) {
	r.relationships, err = new(models.Relationship).GetAllFollowers(emailID)
	if err != nil {
		err = errors.New("Unable to fetch followers")
		log.Info(err.Error())
		return err
	}

	return nil
}

// GetFriends returns all mutual friend relationships
func GetMutualFriends(email1, email2 string) *GetRelationships {
	n := new(GetRelationships)
	n.email1 = email1
	n.email2 = email2
	return n
}

// Data returns all relationships
func (r *GetRelationships) Data() []models.Relationship {
	return r.relationships
}

// Data returns connected emails from all returned relationships
func (r *GetRelationships) ConnectedList() []string {
	list := make([]string, 0)
	for _, elem := range r.relationships {
		list = append(list, elem.Email1)
	}

	return list
}

// Do tasks
func (r *GetRelationships) Do() (err error) {
	if err = r.getAllRelationships(); err != nil {
		return err
	}

	return nil
}

func (r *GetRelationships) getAllRelationships() (err error) {
	r.relationships, err = new(models.Relationship).GetAll(r.email1)
	if err != nil {
		err = errors.New("Unable to fetch relationships")
		log.Info(err.Error())
		return err
	}

	return nil
}
