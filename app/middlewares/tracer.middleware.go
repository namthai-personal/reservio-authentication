package middlewares

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TracerIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId := uuid.New()
		d := color.New(color.BgCyan, color.Bold)
		d.Print("Session:")
		color.Cyan(" %s", sessionId)
		if c.Request.Header.Get("x-tracer-id") == "" {
			c.Header("x-tracer-id", sessionId.String())
		}

		c.Next()
	}
}
