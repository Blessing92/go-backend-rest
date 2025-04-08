package main

import (
	"github.com/gin-gonic/gin"
	"go-backend-rest/db"
	"go-backend-rest/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
