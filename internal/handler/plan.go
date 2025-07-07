package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/idilhaq/finbuddy/internal/db"
)

// POST /api/plans
func CreateOrUpdateMonthlyPlan(c *gin.Context) {
	var userID uuid.UUID
	uidParam, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	uid, ok := uidParam.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}
	parsedUUID, err := uuid.Parse(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	userID = parsedUUID

	var input db.MonthlyPlan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = userID

	// Upsert MonthlyPlan
	var existing db.MonthlyPlan
	if err := db.DB.Where("user_id = ? AND month = ?", input.UserID, input.Month).First(&existing).Error; err == nil {
		input.ID = existing.ID
		db.DB.Model(&existing).Association("Items").Clear()
	}

	if err := db.DB.Save(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, input)
}

// GET /api/plans/:month
func GetMonthlyPlan(c *gin.Context) {
	var userID uuid.UUID
	uidParam, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	uid, ok := uidParam.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}
	parsedUUID, err := uuid.Parse(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	userID = parsedUUID

	month := c.Param("month")
	parsedMonth, err := time.Parse("2006-01", month)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month format. Use YYYY-MM"})
		return
	}

	var plan db.MonthlyPlan
	if err := db.DB.Preload("Items.Pocket").Preload("Items.InvestmentGoal").
		Where("user_id = ? AND month = ?", userID, parsedMonth).
		First(&plan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}

	c.JSON(http.StatusOK, plan)
}

// DELETE /api/plans/:month
func DeleteMonthlyPlan(c *gin.Context) {
	var userID uuid.UUID
	uidParam, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	uid, ok := uidParam.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}
	parsedUUID, err := uuid.Parse(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	userID = parsedUUID

	month := c.Param("month")
	parsedMonth, err := time.Parse("2006-01", month)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month format. Use YYYY-MM"})
		return
	}

	if err := db.DB.Where("user_id = ? AND month = ?", userID, parsedMonth).Delete(&db.MonthlyPlan{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plan deleted"})
}
