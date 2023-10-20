package middlewares

import (
	"github.com/gin-gonic/gin"
)

type GeneralError struct {
	Message string `json:"message"`
	Detail  string `json:"detaiil"`
}

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			detail := GeneralError{
				Message: "Service Error",
				Detail:  c.Errors[0].Error(),
			}
			c.IndentedJSON(500, detail)
		}
	}
}
