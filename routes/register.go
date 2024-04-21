package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"),10 ,64)
	if err != nil { 
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil { 
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	err = event.Register(userId)
	if err != nil { 
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
	
}

func cancelRegisteration(context *gin.Context){
    userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"),10 ,64)
	if err != nil { 
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	var event models.Events
	event.ID = eventId

	err = event.CancelRegisteration(userId)
	if err != nil { 
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "cancel Registered!"})

}