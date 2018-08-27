package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Connection controller
type Connection struct {
	// Base
}

// Connect will form a connection between 2 IDs
func (b *Connection) Connect(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}

// Subscribe will subscribe an ID to another
func (b *Connection) Subscribe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}

// Block will block an ID from another
func (b *Connection) Block(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}

// GetMutualFriends returns all mutual connections between 2 IDs
func (b *Connection) GetMutualFriends(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}
