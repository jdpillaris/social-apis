package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Relationship controller
type Relationship struct {
	// Base
}

// Connect will form a friendship between 2 IDs
func (b *Relationship) Connect(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}

// Subscribe will subscribe an ID to another
func (b *Relationship) Subscribe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}

// Block will block an ID from another
func (b *Relationship) Block(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}

// GetMutualFriends returns all mutual friends between 2 IDs
func (b *Relationship) GetMutualFriends(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}
