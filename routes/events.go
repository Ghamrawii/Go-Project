package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func createEvents(context *gin.Context) {
	
	var event models.Events

	err := context.ShouldBindJSON(&event)
	if err != nil { 
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data request"})
		return
	}

	userId:= context.GetInt64("userId")
	event.UserID = userId
	err = event.Save()

	if err != nil {
		fmt.Println("Error saving event:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event. try again later "})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event is created ", "event": event})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id "})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event by id. try again later "})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {

	events, err := models.GetallEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event. try again later "})
		return
	}

	context.JSON(http.StatusOK, events)
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id "})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event by id. try again later "})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to updated the event"})
		return
	}

	var updateEvent models.Events
	err = context.ShouldBindJSON(&updateEvent)

	updateEvent.ID = eventId

	err = updateEvent.UpdateEvent()
	if err != nil {
		fmt.Println("err is", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event "})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event update successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id "})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event by id. try again later "})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete the event"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event . try again later "})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event delte successfully"})
}
