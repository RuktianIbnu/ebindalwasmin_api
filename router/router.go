package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	// userHandler := uh.NewHandler()
	// generalHandler := gh.NewHandler()
	// kategoriPNBPHandler := kh.NewHandler()
	// pasporHandler := ph.NewHandler()
	// visaHandler := vh.NewHandler()
	// intalHandler := ih.NewHandler()
	// pnbpHandler := pnbph.NewHandler()
	// satkerHandler := sh.NewHandler()

	// v1 := r.Group("/v1")
	// {
	// 	v1.POST("/login", generalHandler.Login)

	// 	resources := v1.Group("/resources")
	// 	{
	// 		resources.GET("/user/:id", userHandler.GetOneByID)
	// 		resources.GET("/satker", satkerHandler.GetAllSatker)
	// 		resources.POST("/filter-monthyear/", satkerHandler.GetReportMonthYear)
	// 		resources.POST("/get-pnbp/", pasporHandler.GetPnbpPaspor)
	// 		resources.POST("/get-PivotPerwilayah/", pnbpHandler.GetPivotPerwilayah)

	// 		resources.POST("/paspor-by/", pasporHandler.GetAllByDate)
	// 		resources.GET("/paspor-pivot-perwilayah", pasporHandler.GetPivotPerwilayah)
	// 		resources.GET("/paspor-byKelaminPer10hari/:id_kantor", pasporHandler.GetKelaminPer10hari)

	// 		resources.POST("/visa-by/tanggal", visaHandler.GetAllByDate)
	// 		//resources.GET("/visa-by/:variable/:value", visaHandler.GetAllByDate)

	// 		resources.POST("/intal-by/", intalHandler.GetAllByDate)
	// 		resources.GET("/intal-pivot-perwilayah", intalHandler.GetPivotPerwilayah)
	// 		resources.GET("/intal-byKelaminPer10hari/:id_kantor", intalHandler.GetKelaminPer10hari)

	// 		resources.GET("/pnbp-kategori", kategoriPNBPHandler.GetAllByParent)
	// 		resources.POST("/pnbp-by/", pnbpHandler.GetAllByDate)
	// 		resources.GET("/pnbp-perbulantahun", pnbpHandler.GetAllkategoriPNBPPerbulanTahun)
	// 		resources.GET("/pnbp-total", pnbpHandler.GetTotalPnbp)
	// 		resources.GET("/pnbp-byKelaminPer10hari/:id_kantor", pnbpHandler.GetKelaminPer10hari)
	// 	}
	// }

	return r
}
