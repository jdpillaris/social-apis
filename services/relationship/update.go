package relationship

import (
	"app/models"
	"errors"

	"github.com/apex/log"
)

// UpdateRelationship service
type UpdateRelationship struct {
	id 			 int64
	person1_ID   int64
	person2_ID   int64
	is_friends	 bool
	follows		 bool
	relationship *models.Relationship
}

// Update instance
func Update(person1_ID, person2_ID int64, is_friends, follows bool) *UpdateRelationship {
	n := new(UpdateRelationship)
	n.person1_ID = person1_ID
	n.person2_ID = person2_ID
	n.is_friends = is_friends
	n.follows = follows
	return n
}

// Data returns the updated project
func (r *UpdateRelationship) Data() *models.Relationship {
	return r.relationship
}

// Do tasks
func (r *UpdateRelationship) Do() (err error) {
	if err = r.validate(); err != nil {
		return err
	}

	if err = r.update(); err != nil {
		return err
	}

	return nil
}

func (p *UpdateRelationship) validate() (err error) {
	return nil
}

func (r *UpdateRelationship) update() (err error) {
	savedRelationship, errGet := r.get()
	if errGet != nil {
		return errors.New("unable to update relationship")
	}

	relationship := new(models.Relationship)
	// if r.name != "" {
	// 	relationship.Name = r.name
	// } else {
	// 	relationship.Name = savedRelationship.Name
	// }

	relationship.Person1 = savedRelationship.Person1
	relationship.Person2 = savedRelationship.Person2
	relationship.CreatedAt = savedRelationship.CreatedAt

	r.relationship, err = relationship.Update(r.id)
	if err != nil {
		return errors.New("Unable to update relationship")
	}

	return nil
}

func (r *UpdateRelationship) get() (models.Relationship, error) {
	relationship, err := new(models.Relationship).Get(r.id)

	if err != nil {
		err = errors.New("Unable to fetch relationship")
		log.Info(err.Error())
		return models.Relationship{}, err
	}

	// Check if there are any relationships found.
	if relationship.ID <= 0 {
		err = errors.New("Relationship not found")
		log.Info(err.Error())
		return models.Relationship{}, err
	}

	return relationship, nil
}
