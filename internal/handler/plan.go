package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/idilhaq/finbuddy/internal/db"
)

type MonthlyPlanRequest struct {
	UserID  uuid.UUID `json:"user_id"`
	Month   string    `json:"month"` // e.g. "2025-05"
	Needs   int       `json:"needs"`
	Wants   int       `json:"wants"`
	Savings int       `json:"savings"`
}

// @Summary      Create or update a monthly plan
// @Description  Create or update a monthly budget plan split by needs, wants, and savings
// @Tags         Plans
// @Accept       json
// @Produce      json
// @Param        request body MonthlyPlanRequest true "Monthly plan input"
// @Success      201 {object} db.MonthlyPlan
// @Success      200 {object} db.MonthlyPlan
// @Failure      400 {object} handler.ErrorResponse
// @Failure      500 {object} handler.ErrorResponse
// @Router       /api/plans [post]
func CreateOrUpdateMonthlyPlan(c *gin.Context) {
	var req MonthlyPlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var plan db.MonthlyPlan
	result := db.DB.Where("month = ? AND user_id = ?", req.Month, req.UserID).First(&plan)

	if result.RowsAffected > 0 {
		// Update existing
		plan.Needs = req.Needs
		plan.Wants = req.Wants
		plan.Savings = req.Savings
		db.DB.Save(&plan)
		c.JSON(http.StatusOK, plan)
		return
	}

	// Create new
	newPlan := db.MonthlyPlan{
		ID:      uuid.New(),
		UserID:  req.UserID,
		Month:   req.Month,
		Needs:   req.Needs,
		Wants:   req.Wants,
		Savings: req.Savings,
	}
	if err := db.DB.Create(&newPlan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPlan)
}

// @Summary      Get a monthly plan by month
// @Description  Retrieve a user's monthly budget plan by YYYY-MM
// @Tags         Plans
// @Produce      json
// @Param        month path string true "Month in format YYYY-MM"
// @Param        user_id query string true "User UUID"
// @Success      200 {object} db.MonthlyPlan
// @Failure      404 {object} handler.ErrorResponse
// @Router       /api/plans/{month} [get]
func GetMonthlyPlan(c *gin.Context) {
	month := c.Param("month")
	userID := c.Query("user_id") // for now, use query param

	var plan db.MonthlyPlan
	if err := db.DB.Where("month = ? AND user_id = ?", month, userID).First(&plan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}
	c.JSON(http.StatusOK, plan)
}
