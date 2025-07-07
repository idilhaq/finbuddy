package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/idilhaq/finbuddy/internal/db"
)

func CreatePocket(c *gin.Context) {
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

	var pocket db.Pocket
	if err := c.ShouldBindJSON(&pocket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pocket.UserID = userID
	if err := db.DB.Create(&pocket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pocket)
}

func GetAllPockets(c *gin.Context) {
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

	var pockets []db.Pocket
	if err := db.DB.Where("user_id = ?", userID).Find(&pockets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pockets)
}

func GetPocketByID(c *gin.Context) {
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
	var pocket db.Pocket
	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&pocket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pocket not found"})
		return
	}
	c.JSON(http.StatusOK, pocket)
}

func UpdatePocket(c *gin.Context) {
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
	var pocket db.Pocket
	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&pocket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pocket not found"})
		return
	}
	if err := c.ShouldBindJSON(&pocket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pocket.UserID = userID
	if err := db.DB.Save(&pocket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pocket)
}

func DeletePocket(c *gin.Context) {
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
	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&db.Pocket{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
