package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/idilhaq/finbuddy/internal/db"
	"github.com/idilhaq/finbuddy/internal/handlers"
)

func main() {
	log.Println("Initializing DB...")
	err := db.Init()
	if err != nil {
		log.Fatalf("❌ DB initialization error: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL successfully.")

	// Auto-migrate models
	err = db.DB.AutoMigrate(&db.Expense{})
	if err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello FinBuddy!"})
	})

	r.GET("/healthz", handlers.HealthzHandler)

	log.Println("🚀 Starting API on :8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
