package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	uh "ebindalwasmin_api/handler/user"
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

	v1 := r.Group("/v1")
	{
		v1.POST("/login", nil)

		// v1.POST("/yayasan", yy.Create)
		// v1.PUT("/yayasan/:id", yy.UpdateOneByID)
		v1.GET("/user/:id", userHandler.GetOneByID)
		// v1.GET("/yayasan", yy.GetAll)
		// v1.DELETE("/yayasan/:id", yy.DeleteOneByID)
	}

	return r
}
