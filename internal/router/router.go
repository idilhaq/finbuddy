package router

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/idilhaq/finbuddy/internal/handler"
	"github.com/idilhaq/finbuddy/internal/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	env := os.Getenv("APP_ENV") // e.g., "development" or "production"

	if env == "development" {
		// 🚀 Development → Allow all origins (unsafe, but fast for local dev)
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://127.0.0.1:5173"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	} else {
		// 🔒 Production → Strict, only allow the real frontend
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"https://finbuddy-ui.onrender.com", "http://127.0.0.1:8080", "http://127.0.0.1:5173", "http://localhost:3000", "http://127.0.0.1:3000"},
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

	api := r.Group("/api")
	{
		// Auth routes
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", handler.Register)
			authGroup.POST("/login", handler.Login)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.JWTAuthMiddleware())
		{
			// Expense routes
			expenseGroup := protected.Group("/expenses")
			{
				expenseGroup.GET("", handler.GetAllExpenses)
				expenseGroup.POST("", handler.CreateExpense)
				expenseGroup.GET("/:id", handler.GetExpenseByID)
				expenseGroup.PUT("/:id", handler.UpdateExpense)
				expenseGroup.DELETE("/:id", handler.DeleteExpense)
			}

			// Dashboard route
			protected.GET("/dashboard", handler.GetDashboardSummary)

			// User routes
			userGroup := protected.Group("/users")
			{
				userGroup.GET("/me", handler.GetUserInfo)
			}

			// Monthly plan routes
			planGroup := protected.Group("/plans")
			{
				planGroup.POST("", handler.CreateOrUpdateMonthlyPlan)
				planGroup.GET("/:month", handler.GetMonthlyPlan)
				planGroup.DELETE("/:month", handler.DeleteMonthlyPlan)
			}

			// Pocket routes
			pocketGroup := protected.Group("/pockets")
			{
				pocketGroup.GET("", handler.GetAllPockets)
				pocketGroup.POST("", handler.CreatePocket)
				pocketGroup.GET("/:id", handler.GetPocketByID)
				pocketGroup.PUT("/:id", handler.UpdatePocket)
				pocketGroup.DELETE("/:id", handler.DeletePocket)
			}

			// Investment goal routes
			goalGroup := protected.Group("/goals")
			{
				goalGroup.GET("", handler.GetAllInvestmentGoals)
				goalGroup.POST("", handler.CreateInvestmentGoal)
				goalGroup.GET("/:id", handler.GetInvestmentGoalByID)
				goalGroup.PUT("/:id", handler.UpdateInvestmentGoal)
				goalGroup.DELETE("/:id", handler.DeleteInvestmentGoal)
			}
		}
	}

	return r
}
