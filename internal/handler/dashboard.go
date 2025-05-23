package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/idilhaq/finbuddy/internal/db"
)

type DashboardResponse struct {
	TotalExpenses int            `json:"total_expenses"`
	Needs         int            `json:"needs"`
	Wants         int            `json:"wants"`
	Savings       int            `json:"savings"`
	TotalSavings  int            `json:"total_savings"`
	BudgetPlan    *BudgetSummary `json:"budget_plan,omitempty"`
}

type BudgetSummary struct {
	Needs   int `json:"needs"`
	Wants   int `json:"wants"`
	Savings int `json:"savings"`
}

// @Summary      Get dashboard summary
// @Description  Return top-level insights for the user’s current month including expenses, breakdowns, and savings
// @Tags         Dashboard
// @Produce      json
// @Param        user_id query string true "User UUID"
// @Param        month query string true "Month in format YYYY-MM"
// @Success      200 {object} handler.DashboardResponse
// @Failure      400 {object} handler.ErrorResponse
// @Router       /dashboard [get]
func GetDashboardSummary(c *gin.Context) {
	userIDStr := c.Query("user_id")
	month := c.Query("month")

	userID, err := uuid.Parse(userIDStr)
	if err != nil || month == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id and month are required"})
		return
	}

	parsedMonth, err := time.Parse("2006-01", month)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid month format (expected YYYY-MM)"})
		return
	}
	start := parsedMonth
	end := start.AddDate(0, 1, 0)

	var total, needs, wants, savings, totalSavings int

	db.DB.Model(&db.Expense{}).
		Where("user_id = ? AND date >= ? AND date < ?", userID, start, end).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total)

	db.DB.Model(&db.Expense{}).
		Where("user_id = ? AND category = ? AND date >= ? AND date < ?", userID, "Needs", start, end).
		Select("COALESCE(SUM(amount), 0)").Scan(&needs)

	db.DB.Model(&db.Expense{}).
		Where("user_id = ? AND category = ? AND date >= ? AND date < ?", userID, "Wants", start, end).
		Select("COALESCE(SUM(amount), 0)").Scan(&wants)

	db.DB.Model(&db.Expense{}).
		Where("user_id = ? AND category = ? AND date >= ? AND date < ?", userID, "Savings", start, end).
		Select("COALESCE(SUM(amount), 0)").Scan(&savings)

	db.DB.Model(&db.Saving{}).
		Where("user_id = ? AND date >= ? AND date < ? AND amount IS NOT NULL", userID, start, end).
		Select("COALESCE(SUM(amount), 0)").Scan(&totalSavings)

	var plan db.MonthlyPlan
	err = db.DB.Where("user_id = ? AND month = ?", userID, month).First(&plan).Error

	var planSummary *BudgetSummary
	if err == nil {
		planSummary = &BudgetSummary{
			Needs:   plan.Needs,
			Wants:   plan.Wants,
			Savings: plan.Savings,
		}
	}

	resp := DashboardResponse{
		TotalExpenses: total,
		Needs:         needs,
		Wants:         wants,
		Savings:       savings,
		TotalSavings:  totalSavings,
		BudgetPlan:    planSummary,
	}

	c.JSON(http.StatusOK, resp)
}
