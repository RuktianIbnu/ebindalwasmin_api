package intal

import (
	resp "ebindalwasmin_api/helpers/response"
	iu "ebindalwasmin_api/usecase/intal"
	"net/http"
	"strconv"
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
	intalUsecase iu.Usecase
}

// NewHandler ...
func NewHandler() Handler {
	return &handler{
		iu.NewUsecase(),
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
	type Body struct {
		Tanggal  string `json:"tanggal"`
		IDSatker int64  `json:"id_satker"`
	}
	var (
		body Body
	)
	c.ShouldBindJSON(&body)

	var (
		tm, _  = time.Parse("2006-01-02", body.Tanggal)
		satker = body.IDSatker
	)

	list, err := m.intalUsecase.GetAllByDate(tm.Unix(), satker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Format(500, err))
		return
	}

	c.JSON(http.StatusOK, resp.Format(200, nil, list))
}

func (m *handler) GetPivotPerwilayah(c *gin.Context) {
	list, err := m.intalUsecase.GetPivotPerwilayah()
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Format(500, err))
		return
	}

	c.JSON(http.StatusOK, resp.Format(200, nil, list))
}

func (m *handler) GetKelaminPer10hari(c *gin.Context) {
	id_kantor, _ := strconv.ParseInt(c.Param("id_kantor"), 10, 64)
	list, err := m.intalUsecase.GetKelaminPer10hari(id_kantor)
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
