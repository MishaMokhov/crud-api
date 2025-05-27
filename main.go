package main

import (
	"crud_api/database"
	"crud_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
