package routes

import (
	"app/controllers"
	"app/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GetEngine gin router
func GetEngine() *gin.Engine {
	r := gin.New()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "HEAD", "DELETE"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Link",
			"X-Total-Count",
			"X-Last-Page",
		},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))
	r.Use(middlewares.Trace())

	// Register api end-points
	v1 := r.Group("/v1")
	{
		user := new(controllers.User)
		v1.GET("/friends", user.GetFriends)
		v1.GET("/followers", user.GetFollowers)
		
		// network := new(controllers.Network)
		// v1.POST("/connect", network.Connect) 
		// v1.POST("/subscribe", network.Subscribe)
		// v1.POST("/block", network.Block)
		// v1.GET("/mutual-friends", network.GetMutualFriends)

		connection := new(controllers.Connection)
		v1.POST("/connect", connection.Connect) 
		v1.POST("/subscribe", connection.Subscribe)
		v1.POST("/block", connection.Block)
		v1.GET("/mutual-friends", connection.GetMutualFriends)
	}

	return r
}
