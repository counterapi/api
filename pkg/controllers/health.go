package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController controls Health operations.
type HealthController struct{}

// Status returns constant response.
func (h HealthController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Operational"})
}
