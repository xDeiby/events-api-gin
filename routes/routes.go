package routes

import (
	"example.com/gin-project/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Public routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signUp)
	server.POST("/login", login)

	// Authenticated routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	
	events := authenticated.Group("/events")
	events.POST("", createEvent)
	events.PUT("/:id", updateEvent)
	events.DELETE("/:id", deleteEvent)
}