package main

import (
	"fmt"
	"gin-myqsl-example/database"
	"gin-myqsl-example/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Initialize database connection
	db, err := database.ConnectDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Initialize router
	r := gin.Default()

	// Setup routes with database connection
	routes.UserRoutes(r, db)

	fmt.Println("🚀 Server running on http://localhost:8081")
	r.Run(":8081")
}
