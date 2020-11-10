package paspor

import (
	resp "ebindalwasmin_api/helpers/response"
	pu "ebindalwasmin_api/usecase/paspor"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler interface {
	// Create(c *gin.Context)
	// UpdateOneByID(c *gin.Context)
	// GetOneByID(c *gin.Context)
	GetAllByDate(c *gin.Context)
	GetPivotPerwilayah(c *gin.Context)
	GetKelaminPer10hari(c *gin.Context)
	// DeleteOneByID(c *gin.Context)
}

type handler struct {
	pasporUsecase pu.Usecase
}

// NewHandler ...
func NewHandler() Handler {
	return &handler{
		pu.NewUsecase(),
	}
}

// func (m *handler) Create(c *gin.Context) {
// 	var (
// 		data model.DomainClient
// 	)

// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		return
// 	}

// 	err := m.dcUsecase.Create(&data)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, resp.Format(500, err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp.Format(200, nil, data))
// }

// func (m *handler) UpdateOneByID(c *gin.Context) {
// 	var (
// 		data   model.DomainClient
// 		ids, _ = strconv.ParseInt(c.Param("id"), 10, 64)
// 	)

// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		return
// 	}

// 	if ids <= 0 {
// 		c.JSON(http.StatusInternalServerError, resp.Format(500, errors.New("Provide a valid ID")))
// 		return
// 	}

// 	data.ID = ids

// 	_, err := m.dcUsecase.UpdateOneByID(&data)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, resp.Format(500, err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp.Format(200, nil, data))
// }

// func (m *handler) GetOneByID(c *gin.Context) {
// 	var (
// 		ids, _ = strconv.ParseInt(c.Param("id"), 10, 64)
// 	)

// 	if ids <= 0 {
// 		c.JSON(http.StatusInternalServerError, resp.Format(500, errors.New("Provide a valid ID")))
// 		return
// 	}

// 	data, err := m.userUsecase.GetOneByID(ids)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, resp.Format(500, err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp.Format(200, nil, data))
// }

func (m *handler) GetAllByDate(c *gin.Context) {
	// var (
	// 	limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))
	// 	page, _  = strconv.Atoi(c.DefaultQuery("page", "1"))
	// 	search   = c.Query("search")
	// )
	var (
		tm, _ = time.Parse("2006-01-02", c.Param("tanggal"))
	)

	log.Println(tm.Unix())
	list, err := m.pasporUsecase.GetAllByDate(tm.Unix())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Format(500, err))
		return
	}

	c.JSON(http.StatusOK, resp.Format(200, nil, list))
}

func (m *handler) GetPivotPerwilayah(c *gin.Context) {
	list, err := m.pasporUsecase.GetPivotPerwilayah()
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Format(500, err))
		return
	}

	c.JSON(http.StatusOK, resp.Format(200, nil, list))
}

func (m *handler) GetKelaminPer10hari(c *gin.Context) {
	type Tgl struct {
		TanggalAwal  string `json:"tanggal_awal"`
		TanggalAkhir string `json:"tanggal_akhir"`
	}
	var (
		date Tgl
	)

	c.ShouldBindJSON(&date)

	date1, _ := time.Parse("2006-01-02", date.TanggalAwal)
	date2, _ := time.Parse("2006-01-02", date.TanggalAkhir)

	log.Println(date1.Unix(), date2.Unix())
	list, err := m.pasporUsecase.GetKelaminPer10hari(date1.Unix(), date2.Unix())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Format(500, err))
		return
	}

	c.JSON(http.StatusOK, resp.Format(200, nil, list))
}

// func (m *handler) DeleteOneByID(c *gin.Context) {
// 	var (
// 		ids, _ = strconv.ParseInt(c.Param("id"), 10, 64)
// 	)

// 	if ids <= 0 {
// 		c.JSON(http.StatusInternalServerError, resp.Format(500, errors.New("Provide a valid ID")))
// 		return
// 	}

// 	_, err := m.dcUsecase.DeleteOneByID(ids)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, resp.Format(500, err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp.Format(200, nil))
// }
