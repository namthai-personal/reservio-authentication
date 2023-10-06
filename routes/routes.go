package routes

import (
	"reservio-be/app/handlers"
	"reservio-be/app/middlewares"
	"reservio-be/app/user"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the Gin router and defines the application routes.
func SetupRouter() *gin.Engine {
	// Initialize Gin router
	router := gin.Default()

	// Use middleware
	router.Use(middlewares.LogRequestMiddleware())

	// Routes
	router.GET("/", handlers.DefaultHandler)

	// Define API routes
	api := router.Group("/user")
	{
		api.GET("/", user.GetUser)
	}

	return router
}
