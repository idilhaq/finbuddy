package main

import (
	"log"

	"github.com/idilhaq/finbuddy/internal/db"
	"github.com/idilhaq/finbuddy/internal/router"

	_ "github.com/idilhaq/finbuddy/docs"
)

func main() {
	log.Println("Initializing DB...")
	if err := db.Init(); err != nil {
		log.Fatalf("❌ DB initialization error: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL successfully.")

	// Auto-migrate your models
	if err := db.DB.AutoMigrate(&db.Expense{}, &db.MonthlyPlan{}, &db.Saving{}); err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}

	// Set up router from the internal/router package
	r := router.SetupRouter()

	log.Println("🚀 Starting API on :3000")
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
