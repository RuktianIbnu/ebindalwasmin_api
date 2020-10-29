package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	gh "ebindalwasmin_api/handler/general"
	uh "ebindalwasmin_api/handler/user"
	"ebindalwasmin_api/middleware/auth"
)

// Routes ...
func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	// for health check purpose only
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userHandler := uh.NewHandler()
	generalHandler := gh.NewHandler()

	v1 := r.Group("/v1")
	{
		v1.POST("/login", generalHandler.Login)

		resources := v1.Group("/resources").Use(auth.Middleware())
		{
			resources.GET("/user/:id", userHandler.GetOneByID)
		}
	}

	return r
}
