package router

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/idilhaq/finbuddy/internal/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	env := os.Getenv("APP_ENV") // e.g., "development" or "production"

	if env == "development" {
		// 🚀 Development → Allow all origins (unsafe, but fast for local dev)
		r.Use(cors.Default())
	} else {
		// 🔒 Production → Strict, only allow the real frontend
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"https://your-production-frontend.com"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	// Health check route
	r.GET("/healthz", handler.HealthzHandler)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello FinBuddy!"})
	})

	// Dashboard route
	r.GET("/api/dashboard", handler.GetDashboardSummary)

	// Group expense routes
	expenseGroup := r.Group("/api/expenses")
	{
		expenseGroup.GET("", handler.GetAllExpenses)
		expenseGroup.POST("", handler.CreateExpense)
		expenseGroup.GET("/:id", handler.GetExpenseByID)
		expenseGroup.PUT("/:id", handler.UpdateExpense)
		expenseGroup.DELETE("/:id", handler.DeleteExpense)
	}

	// Group monthly plan routes
	planGroup := r.Group("/api/plans")
	{
		planGroup.POST("", handler.CreateOrUpdateMonthlyPlan)
		planGroup.GET("/:month", handler.GetMonthlyPlan)
	}

	return r
}
