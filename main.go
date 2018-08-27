package main

import (
	"app/routes"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	log.SetHandler(text.New(os.Stderr))

	// Configure gin run mode
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	// Start HTTP server
	err := routes.GetEngine().Run(":" + viper.GetString("app.port"))
	if err != nil {
		log.WithError(err).Error("gin error")
	}
}
