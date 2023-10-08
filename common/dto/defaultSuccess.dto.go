package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultSuccessResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
	})
}
