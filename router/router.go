package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	gh "ebindalwasmin_api/handler/general"
	ih "ebindalwasmin_api/handler/intal"
	kh "ebindalwasmin_api/handler/kategoripnbp"
	ph "ebindalwasmin_api/handler/paspor"
	pnbph "ebindalwasmin_api/handler/pnbp"
	uh "ebindalwasmin_api/handler/user"
	vh "ebindalwasmin_api/handler/visa"
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
	kategoriPNBPHandler := kh.NewHandler()
	pasporHandler := ph.NewHandler()
	visaHandler := vh.NewHandler()
	intalHandler := ih.NewHandler()
	pnbpHandler := pnbph.NewHandler()

	v1 := r.Group("/v1")
	{
		v1.POST("/login", generalHandler.Login)

		resources := v1.Group("/resources")
		{
			resources.GET("/user/:id", userHandler.GetOneByID)

			resources.GET("/kategori-pnbp/:parent", kategoriPNBPHandler.GetAllByParent)

			resources.GET("/paspor-by/:tanggal", pasporHandler.GetAllByDate)
			resources.GET("/visa-by/:tanggal", visaHandler.GetAllByDate)
			resources.GET("/intal-by/:tanggal", intalHandler.GetAllByDate)
			resources.GET("/pnbp-by/:tanggal", pnbpHandler.GetAllByDate)
		}
	}

	return r
}
