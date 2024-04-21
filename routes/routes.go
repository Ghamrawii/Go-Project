package routes

import (
	"example.com/events/middlewares"
	"github.com/gin-gonic/gin"
)

func AppRoutes(server *gin.Engine) {
	// events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/events",middlewares.Authenticated, createEvents)
	server.PUT("/events/:id", middlewares.Authenticated, updateEvent)
	server.DELETE("/events/:id",middlewares.Authenticated, deleteEvent)

	// users
	server.POST("/user/signup", signup)
	server.POST("/user/login", login)

	// register
	server.POST("/events/:id/register",middlewares.Authenticated, registerForEvent)
	server.DELETE("/events/:id/register",middlewares.Authenticated, cancelRegisteration)

}
