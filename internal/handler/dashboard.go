package handler

import (
	"github.com/gin-gonic/gin"
)

type DashboardResponse struct {
	MonthlyBudget int `json:"monthlyBudget"`
	TotalExpenses int `json:"totalExpenses"`
	TotalSavings  int `json:"totalSavings"`
}

func DashboardHandler(c *gin.Context) {
	response := DashboardResponse{
		MonthlyBudget: 5000000,
		TotalExpenses: 2850000,
		TotalSavings:  2150000,
	}

	c.JSON(200, response)
}
