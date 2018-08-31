package main

import (
	"app/models"
	"app/routes"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/gin-gonic/gin"
)

func init() {
	log.SetHandler(text.New(os.Stderr))

	// Connect to DB
	models.Connect()

	// Configure gin run mode
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	// Start HTTP server
	err := routes.GetEngine().Run(":" + "8080")

	if err != nil {
		log.WithError(err).Error("gin error")
	}
}
