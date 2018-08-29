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
		// person := new(controllers.Person)
		// v1.GET("/friends", person.GetFriends)
		// v1.GET("/followers", person.GetFollowers)

		relationship := new(controllers.Relationship)
		v1.GET("/friends", relationship.GetFriends)
		v1.GET("/followers", relationship.GetFollowers)
		v1.POST("/connect", relationship.Connect) 
		v1.POST("/subscribe", relationship.Subscribe)
		v1.POST("/block", relationship.Block)
		v1.GET("/mutual-friends", relationship.GetMutualFriends)
	}

	return r
}
