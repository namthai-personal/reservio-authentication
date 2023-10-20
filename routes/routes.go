package routes

import (
	"reservio-be/app/handlers"
	"reservio-be/app/middlewares"
	"reservio-be/app/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Initialize Gin router
	router := gin.Default()

	// Use middleware
	router.Use(middlewares.TracerIDMiddleware())
	router.Use(middlewares.ErrorHandling())
	router.Use(gin.Recovery())

	// Default Routes
	router.GET("/", handlers.DefaultHandler)
	router.GET("/healthcheck", handlers.Healthcheck)

	// User api group
	userApi := router.Group("/user")
	{
		userApi.GET("/", user.ListUser)
		userApi.GET("/:id", user.GetUserById)
		userApi.POST("/", user.CreateUser)
		userApi.POST("/login", user.Login)
	}

	return router
}
