package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Person controller
type Person struct {
	// Base
}

// getFriends returns all friends of a person
func (b *Person) GetFriends(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}

// getFollowers returns all subscribers of a person
func (b *Person) GetFollowers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}
