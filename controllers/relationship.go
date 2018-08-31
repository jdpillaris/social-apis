package controllers

import (
	"app/controllers/bindings"
	"app/services/relationship"
	"net/http"

	"github.com/apex/log"
	"github.com/gin-gonic/gin"
)

// Relationship controller
type Relationship struct {
	// Base
}

// Connect will form a friendship between 2 email-IDs
func (b *Relationship) Connect(c *gin.Context) {
	var err error
	
	json := new(bindings.PersonPair)
	if err = c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Info("Bind parse error")
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	email_1 := json.PersonPair[0]
	email_2 := json.PersonPair[1]

	// personService := person.NewPerson(email_1)

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
	
	json := new(bindings.Request)
	if err = c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Info("Bind parse error")
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	requestor := json.Requestor
	target := json.Target

	relationshipService := relationship.NewRelationship(requestor, target, false, true)
	if err = relationshipService.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// Block will block an ID from another
func (b *Relationship) Block(c *gin.Context) {
	var err error

	json := new(bindings.Request)
	if err = c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Info("Bind parse error")
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	requestor := json.Requestor
	target := json.Target

	s := relationship.NewRelationship(requestor, target, false, false)
	if err = s.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// getFriends returns all friends of a person
func (b *Relationship) GetFriends(c *gin.Context) {
	json := new(bindings.GetFriends)

	if err := c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Info("Bind parse error")
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	email := json.Email

	s := relationship.GetAllRelationships(email)

	if err := s.GetFriends(email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	friends := s.ConnectedList()
	c.JSON(http.StatusOK, gin.H{"success": true, "friends": friends, "count": len(friends)})
}

// getFollowers returns all subscribers of a person
func (b *Relationship) GetFollowers(c *gin.Context) {
	json := new(bindings.GetFollowers)

	if err := c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Info("Bind parse error")
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	email := json.Email
	// post := json.Post
	// email := "someone@domain.com"

	s := relationship.GetAllRelationships(email)

	if err := s.GetFollowers(email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	// postMentions := getMentionsFromPost(post)
	followers := s.ConnectedList()
	c.JSON(http.StatusOK, gin.H{"success": true, "recipients": followers})
}

// GetMutualFriends returns all mutual friends between 2 IDs
func (b *Relationship) GetMutualFriends(c *gin.Context) {
	json := new(bindings.PersonPair)

	if err := c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Info("Bind parse error")
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	email_1 := json.PersonPair[0]
	email_2 := json.PersonPair[1]

	s := relationship.GetMutualFriends(email_1, email_2)

	if err := s.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	mutualFriends := s.ConnectedList()
	c.JSON(http.StatusOK, gin.H{"friends": mutualFriends})
}

func getMentionsFromPost(text string) (emails []string) {
	emails = make([]string, 0)

	return emails
}

