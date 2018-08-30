package relationship

import (
	"app/services/person"
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
	var person1_id, person2_id int64
	if person1_id, person2_id, err = r.createPersons(); err != nil {
		return err
	}

	if err = r.validate(); err != nil {
		return err
	}

	if err = r.create(person1_id, person2_id); err != nil {
		return err
	}
	//For friendship, add extra record - since it is a symmetrical relationship
	if r.is_friends {
		if err = r.create(person2_id, person1_id); err != nil {
			return err
		}
	}

	return nil
}

//validate function validates the relationship attributes 
func (r *CreateRelationship) validate() (err error) {
	return nil
}

func (r *CreateRelationship) createPersons() (person1_id, person2_id int64, err error) {	
	personService := person.NewPerson(r.email1)
	if err = personService.Do(); err != nil {
		return person1_id, person2_id, err
	}
	person1_id = personService.Data().ID

	personService = person.NewPerson(r.email2)
	if err = personService.Do(); err != nil {
		return person1_id, person2_id, err
	}
	person2_id = personService.Data().ID

	return person1_id, person2_id, nil
}

// Create a record
func (r *CreateRelationship) create(person1_id, person2_id int64) (err error) {
	relationship := new(models.Relationship)
	relationship.Person1 = person1_id
	relationship.Person2 = person2_id
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
