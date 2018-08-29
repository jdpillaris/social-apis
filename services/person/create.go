package person

import (
	"app/models"
	"time"

	"github.com/apex/log"
)

// CreatePerson service
type CreatePerson struct {
	email  string
	person *models.Person
}

// CreatePerson instance
func NewPerson(email string) *CreatePerson {
	n := new(CreatePerson)
	n.email = email
	n.person = new(models.Person)
	return n
}

// Data returns the new person
func (p *CreatePerson) Data() *models.Person {
	return p.person
}

// Do will validate data and create a person
func (p *CreatePerson) Do() (err error) {
	if err = p.validate(); err != nil {
		return err
	}

	if err = p.create(); err != nil {
		return err
	}

	return nil
}

//validate person email
func (p *CreatePerson) validate() (err error) {
	return nil
}

// Create a record
func (p *CreatePerson) create() (err error) {
	person := new(models.Person)
	person.Email = p.email
	person.CreatedAt = time.Now().UTC()
	p.person, err = person.Create()
	if err != nil {
		log.WithError(err).Error("Unable to add person")
		return err
	}

	return nil
}
