package general

import (
	resp "ebindalwasmin_api/helpers/response"
	ur "ebindalwasmin_api/repository/user"
	gu "ebindalwasmin_api/usecase/general"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler interface {
	Login(c *gin.Context)
	// UpdateOneByID(c *gin.Context)
	// GetOneByID(c *gin.Context)
	// GetAll(c *gin.Context)
	// DeleteOneByID(c *gin.Context)
}

type handler struct {
	generalUsecase gu.Usecase
}

type usecase struct {
	userRepo ur.Repository
}

// NewHandler ...
func NewHandler() Handler {
	return &handler{
		gu.NewUsecase(),
	}
}

func (m *handler) Login(c *gin.Context) {
	type login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var (
		loginData login
	)

	if err := c.ShouldBindJSON(&loginData); err != nil {
		return
	}

	token, err := m.generalUsecase.Login(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Format(500, err))
		return
	}

	id, _, err := ur.NewRepository().GetUserPasswordByEmail(loginData.Email)
	if err != nil {
	}

	user, err := ur.NewRepository().GetOneByID(id)
	if err != nil {
	}

	c.JSON(http.StatusOK, resp.Format(200, nil, gin.H{"token": token, "user": user}))
}
