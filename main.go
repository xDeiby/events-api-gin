package main

import (
	"net/http"
	"strconv"

	"example.com/gin-project/db"
	"example.com/gin-project/models"
	"github.com/gin-gonic/gin"
)

func main() {
		db.InitDb()
		server := gin.Default()

		server.GET("/events", getEvents)
		server.GET("/events/:id", getEvent)
		server.POST("/events", createEvent)

		server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get events",
		})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to convertd id to number",
		})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to find event id",
		})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})

		return
	}

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save event in database",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event": event,
	})
}