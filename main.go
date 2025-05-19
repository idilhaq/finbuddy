package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/idilhaq/finbuddy/internal/db"
)

func main() {
	db.Init()
	fmt.Println("✅ Connected to PostgreSQL successfully.")

	// Auto-migrate models
	db.DB.AutoMigrate(&db.Expense{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello FinBuddy!"})
	})

	r.Run()
}
