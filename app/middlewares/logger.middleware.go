package middlewares

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func LogRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		color.Yellow("User-Agent: %s", c.GetHeader("User-Agent"))

		// Process the request
		c.Next()
	}
}
