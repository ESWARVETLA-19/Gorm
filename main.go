package main

import (
	"log"
	"myproject/db"
	"myproject/models"
	"myproject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Run migrations with injected DB
	models.Migrate(conn)
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
