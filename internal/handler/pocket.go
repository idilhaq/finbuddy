package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idilhaq/finbuddy/internal/db"
)

func CreatePocket(c *gin.Context) {
	var pocket db.Pocket
	if err := c.ShouldBindJSON(&pocket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&pocket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pocket)
}

func GetAllPockets(c *gin.Context) {
	var pockets []db.Pocket
	if err := db.DB.Find(&pockets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pockets)
}

func GetPocketByID(c *gin.Context) {
	id := c.Param("id")
	var pocket db.Pocket
	if err := db.DB.First(&pocket, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pocket not found"})
		return
	}
	c.JSON(http.StatusOK, pocket)
}

func UpdatePocket(c *gin.Context) {
	id := c.Param("id")
	var pocket db.Pocket
	if err := db.DB.First(&pocket, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pocket not found"})
		return
	}
	if err := c.ShouldBindJSON(&pocket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Save(&pocket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pocket)
}

func DeletePocket(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&db.Pocket{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
