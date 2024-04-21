package main

import (
	"example.com/events/db"
	"example.com/events/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.AppRoutes(server)

	server.Run(":8080")
}
