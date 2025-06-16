package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idilhaq/finbuddy/internal/db"
)

func CreateInvestmentGoal(c *gin.Context) {
	var goal db.InvestmentGoal
	if err := c.ShouldBindJSON(&goal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&goal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, goal)
}

func GetAllInvestmentGoals(c *gin.Context) {
	var goals []db.InvestmentGoal
	if err := db.DB.Find(&goals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, goals)
}

func GetInvestmentGoalByID(c *gin.Context) {
	id := c.Param("id")
	var goal db.InvestmentGoal
	if err := db.DB.First(&goal, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Goal not found"})
		return
	}
	c.JSON(http.StatusOK, goal)
}

func UpdateInvestmentGoal(c *gin.Context) {
	id := c.Param("id")
	var goal db.InvestmentGoal
	if err := db.DB.First(&goal, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Goal not found"})
		return
	}
	if err := c.ShouldBindJSON(&goal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Save(&goal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, goal)
}

func DeleteInvestmentGoal(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&db.InvestmentGoal{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
