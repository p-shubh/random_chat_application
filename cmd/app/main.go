// /cmd/app/main.go
package main

import (

	"rca/internal/handlers"
	"rca/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDatabase()

	router.GET("/ws", handlers.HandleConnections)

	router.Run(":8080")
}
