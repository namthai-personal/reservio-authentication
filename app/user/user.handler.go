package user

import (
	"reservio-be/infrastructure/database"

	"github.com/gin-gonic/gin"
)

type UserPagination struct {
	Items []User `json:"items"`
	Page  int    `json:"page"`
}

const (
	QUERY_LIMIT = 1000
)

func GetUser(c *gin.Context) {
	db := database.GetDb()

	users := []User{}

	db.Limit(QUERY_LIMIT).Find(&users)

	c.JSON(200, UserPagination{users, 1})
}
