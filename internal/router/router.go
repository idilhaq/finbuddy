package router

import (
	"github.com/idilhaq/finbuddy/internal/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Health check route
	r.GET("/healthz", handler.HealthzHandler)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello FinBuddy!"})
	})

	// Group expense routes
	r.GET("/api/dashboard", handler.DashboardHandler)
	expenseGroup := r.Group("/expenses")
	{
		expenseGroup.GET("", handler.GetAllExpenses)
		expenseGroup.POST("", handler.CreateExpense)
		expenseGroup.GET("/:id", handler.GetExpenseByID)
		expenseGroup.PUT("/:id", handler.UpdateExpense)
		expenseGroup.DELETE("/:id", handler.DeleteExpense)
	}

	return r
}
