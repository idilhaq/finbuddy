package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/idilhaq/finbuddy/internal/db"
)

func CreateInvestmentGoal(c *gin.Context) {
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

	var goal db.InvestmentGoal
	if err := c.ShouldBindJSON(&goal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	goal.UserID = userID // Apply userID
	if err := db.DB.Create(&goal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, goal)
}

func GetAllInvestmentGoals(c *gin.Context) {
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

	var goals []db.InvestmentGoal
	if err := db.DB.Where("user_id = ?", userID).Find(&goals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, goals)
}

func GetInvestmentGoalByID(c *gin.Context) {
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

	id := c.Param("id")
	var goal db.InvestmentGoal
	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&goal).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Goal not found"})
		return
	}
	c.JSON(http.StatusOK, goal)
}

func UpdateInvestmentGoal(c *gin.Context) {
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

	id := c.Param("id")
	var goal db.InvestmentGoal
	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&goal).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Goal not found"})
		return
	}
	if err := c.ShouldBindJSON(&goal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	goal.UserID = userID // Ensure userID is not changed
	if err := db.DB.Save(&goal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, goal)
}

func DeleteInvestmentGoal(c *gin.Context) {
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

	id := c.Param("id")
	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&db.InvestmentGoal{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
