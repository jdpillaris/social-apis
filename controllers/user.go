package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User controller
type User struct {
	// Base
}

// getFriends returns all connections of user
func (b *User) GetFriends(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}

// getFollowers returns all subscribers of user
func (b *User) GetFollowers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": ""})
}
