package controllers

import (
	"app/services/relationship"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Relationship controller
type Relationship struct {
	// Base
}

// Connect will form a friendship between 2 email-IDs
func (b *Relationship) Connect(c *gin.Context) {
	var err error
	email_1 := "person1@domain.com"
	email_2 := "person2@domain.com"

	// personService := person.NewPerson(email_1)
	// if err = personService.Do(); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"success": false})
	// 	return
	// }
	// personService = person.NewPerson(email_2)
	// if err = personService.Do(); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"success": false})
	// 	return
	// }

	relationshipService := relationship.NewRelationship(email_1, email_2, true, true)
	if err = relationshipService.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// Subscribe will subscribe an email-ID to another
func (b *Relationship) Subscribe(c *gin.Context) {
	var err error
	requestor := "person1@domain.com"
	target := "person2@domain.com"

	s := relationship.NewRelationship(requestor, target, false, true)
	if err = s.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// Block will block an ID from another
func (b *Relationship) Block(c *gin.Context) {
	var err error
	requestor := "person1@domain.com"
	target := "person2@domain.com"

	s := relationship.NewRelationship(requestor, target, false, false)
	if err = s.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// getFriends returns all friends of a person
func (b *Relationship) GetFriends(c *gin.Context) {
	email := "someone@domain.com"
	s := relationship.GetFriends(email)

	if err := s.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// // getFollowers returns all subscribers of a person
func (b *Relationship) GetFollowers(c *gin.Context) {
	email := "someone@domain.com"
	s := relationship.GetFollowers(email)

	if err := s.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// GetMutualFriends returns all mutual friends between 2 IDs
func (b *Relationship) GetMutualFriends(c *gin.Context) {
	email_1 := "person1@domain.com"
	email_2 := "person2@domain.com"
	s := relationship.GetMutualFriends(email_1, email_2)

	if err := s.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
