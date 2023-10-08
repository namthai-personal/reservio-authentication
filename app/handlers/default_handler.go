package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"reservio-be/common/dto"
)

func DefaultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Go Backend, From Nam Thai with Love 💗",
	})
}

// Healthcheck godoc
//
//	@Summary	Healthcheck
//	@Tags		Default API
//	@Router		/healthcheck [get]
func Healthcheck(c *gin.Context) {
	dto.DefaultSuccessResponse(c)
}
