package controllers

import (
	"fmt"
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
	fmt.Println("##### WE ARE HERE: Connect #####")
	fmt.Println("JSON: %v", json)
	if err = c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Info("Bind parse error")
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	email_1 := json.PersonPair[0]
	email_2 := json.PersonPair[1]
	// email_1 := "person1@domain.com"
	// email_2 := "person2@domain.com"

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
	fmt.Println("##### WE ARE HERE: Subscribe #####")
	fmt.Println("JSON: %v", json)
	if err = c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Info("Bind parse error")
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	requestor := json.Requestor
	target := json.Target
	// requestor := "person1@domain.com"
	// target := "person2@domain.com"

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

	json := new(bindings.Request)
	fmt.Println("##### WE ARE HERE: Block #####")
	fmt.Println("JSON: %v", json)
	if err = c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Info("Bind parse error")
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	requestor := json.Requestor
	target := json.Target
	// requestor := "person1@domain.com"
	// target := "person2@domain.com"

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
	fmt.Println("##### WE ARE HERE: GetFriends #####")
	fmt.Println("JSON: %v", json)
	if err := c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Info("Bind parse error")
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	email := json.Email
	// email := "someone@domain.com"

	s := relationship.GetAllRelationships(email)
	// s := relationship.GetFriends(email)

	if err := s.GetFriends(email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// getFollowers returns all subscribers of a person
func (b *Relationship) GetFollowers(c *gin.Context) {
	json := new(bindings.GetFollowers)
	fmt.Println("##### WE ARE HERE: GetMutualFriends #####")
	fmt.Println("JSON: %v", json)
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
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// GetMutualFriends returns all mutual friends between 2 IDs
func (b *Relationship) GetMutualFriends(c *gin.Context) {
	json := new(bindings.PersonPair)
	fmt.Println("##### WE ARE HERE: GetMutualFriends #####")
	fmt.Println("JSON: %v", json)
	if err := c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Info("Bind parse error")
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	email_1 := json.PersonPair[0]
	email_2 := json.PersonPair[1]
	// email_1 := "person1@domain.com"
	// email_2 := "person2@domain.com"

	s := relationship.GetMutualFriends(email_1, email_2)

	if err := s.Do(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func getMentionsFromPost(text string) (emails []string) {
	emails = make([]string, 0)

	return emails
}

