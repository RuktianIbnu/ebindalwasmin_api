package auth

import (
	"errors"
	"log"
	"net/http"

	"ebindalwasmin_api/helpers/jwt"
	resp "ebindalwasmin_api/helpers/response"

	"github.com/gin-gonic/gin"
)

// Middleware ...
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Header struct {
			Authorization string `header:"Authorization"`
		}

		header := Header{}
		c.ShouldBindHeader(&header)

		if header.Authorization == "" {
			c.JSON(http.StatusUnauthorized, resp.Format(500, errors.New("please provide authorization")))
			c.Abort()
			log.Println("kosong", header.Authorization)
			return
		}

		if err := jwt.TokenValid(header.Authorization); err != nil {
			c.JSON(http.StatusUnauthorized, resp.Format(500, errors.New("unauthorize")))
			c.Abort()
			log.Println("err != nil", header.Authorization)
			return
		}

		if metadata, err := jwt.ExtractTokenMetadata(header.Authorization); err == nil {
			log.Println(metadata)
			c.Set("user_id", metadata.UserID)
			c.Set("email", metadata.Email)
			log.Println("err == nil", header.Authorization)
		} else {
			c.JSON(http.StatusUnauthorized, resp.Format(500, errors.New("unable to extract token metadata")))
			c.Abort()
			return
		}

		c.Set("token", header.Authorization)

		c.Next()
	}
}
