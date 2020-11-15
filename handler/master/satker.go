package master

import (
	resp "ebindalwasmin_api/helpers/response"
	su "ebindalwasmin_api/usecase/master"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler interface {
	// Create(c *gin.Context)
	// UpdateOneByID(c *gin.Context)
	// GetOneByID(c *gin.Context)
	GetAllSatker(c *gin.Context)
	GetReportMonthYear(c *gin.Context)
	// DeleteOneByID(c *gin.Context)
}

type handler struct {
	satkerUsecase su.Usecase
}

// NewHandler ...
func NewHandler() Handler {
	return &handler{
		su.NewUsecase(),
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

func (m *handler) GetAllSatker(c *gin.Context) {
	list, err := m.satkerUsecase.GetAllSatker()
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Format(500, err))
		return
	}

	c.JSON(http.StatusOK, resp.Format(200, nil, list))
}

func (m *handler) GetReportMonthYear(c *gin.Context) {
	type Body struct {
		Cekbox  string `json:"cekbox"`
		IDJenis string `json:"id_jenis"`
	}
	var (
		body Body
	)
	c.ShouldBindJSON(&body)
	// tm_awal, _   = time.Parse("2006-01-02", c.Param("tanggal_awal"))
	// tm_akhir, _  = time.Parse("2006-01-02", c.Param("tanggal_akhir"))
	cekbox, _ := strconv.ParseBool(body.Cekbox)
	id_jenis := body.IDJenis

	//id_satker    = c.Param("id_satker")

	log.Println(cekbox, id_jenis)
	if cekbox == true {
		switch id_jenis {
		case "all":
			list, err := m.satkerUsecase.GetDataReportMonthYearAll("all")
			if err != nil {
				c.JSON(http.StatusInternalServerError, resp.Format(500, err))
				return
			}
			c.JSON(http.StatusOK, resp.Format(200, nil, list))
		case "1":
			list, err := m.satkerUsecase.GetDataReportMonthYear("visa")
			if err != nil {
				c.JSON(http.StatusInternalServerError, resp.Format(500, err))
				return
			}
			c.JSON(http.StatusOK, resp.Format(200, nil, list))
		case "2":
			list, err := m.satkerUsecase.GetDataReportMonthYear("paspor")
			if err != nil {
				c.JSON(http.StatusInternalServerError, resp.Format(500, err))
				return
			}
			c.JSON(http.StatusOK, resp.Format(200, nil, list))
		case "3":
			list, err := m.satkerUsecase.GetDataReportMonthYear("izintinggal")
			if err != nil {
				c.JSON(http.StatusInternalServerError, resp.Format(500, err))
				return
			}
			c.JSON(http.StatusOK, resp.Format(200, nil, list))
		case "4":
			list, err := m.satkerUsecase.GetDataReportMonthYear("pnbplainnya")
			if err != nil {
				c.JSON(http.StatusInternalServerError, resp.Format(500, err))
				return
			}
			c.JSON(http.StatusOK, resp.Format(200, nil, list))
		}
	} else {

	}
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
