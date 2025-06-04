package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Health Check
// @Description  Returns OK
// @Tags         Health
// @Success      200  {object}  map[string]string
// @Router       /api/healthz [get]
func HealthzHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
