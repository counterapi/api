package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController controls Health operations.
type HomeController struct{}

// Redirect returns constant response.
func (h HomeController) Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://counterapi.dev/")
}
