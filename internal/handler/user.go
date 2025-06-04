package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idilhaq/finbuddy/internal/db"
)

// GetUserInfo handles GET requests to retrieve information about the authenticated user.
//
// @Summary      Get user information
// @Description  Retrieves the authenticated user's information based on the user_id stored in the context.
// @Tags         Users
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "User information"
// @Failure      401  {object}  map[string]string       "Unauthorized"
// @Failure      404  {object}  map[string]string       "User not found"
// @Router       /user/info [get]
func GetUserInfo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var user db.User
	if err := db.DB.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userInfo := map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"role":       user.Role,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}
	c.JSON(http.StatusOK, userInfo)
}

// GetCurrentUser handles the HTTP request to retrieve the current authenticated user's information.
// @Summary      Get current user
// @Description  Retrieves the current authenticated user's information from the context.
// @Tags         Users
// @Produce      json
// @Success      200  {object}  interface{}
// @Failure      401  {object}  map[string]string       "Unauthorized"
// @Router       /user/me [get]
func GetCurrentUser(c *gin.Context) {
	user, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	c.JSON(http.StatusOK, user)
}
