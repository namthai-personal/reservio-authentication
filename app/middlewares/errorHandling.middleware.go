package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type GeneralError struct {
	Message string `json:"message"`
	Detail  string `json:"detaiil"`
}

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			detail := GeneralError{}
			err := c.Errors[0]
			var postgresErr *pgconn.PgError
			errors.As(err, &postgresErr)

			if postgresErr.Code == "23505" {
				detail = GeneralError{
					Message: "DUPLICATE_KEY",
					Detail:  c.Errors[0].Error(),
				}
				c.IndentedJSON(400, detail)
			} else {
				detail = GeneralError{
					Message: "Service Error",
					Detail:  c.Errors[0].Error(),
				}
				c.IndentedJSON(500, detail)
			}

		}
	}
}
