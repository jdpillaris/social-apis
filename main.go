package main

import (
	"app/routes"

	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
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
