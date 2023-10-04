package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	errorMessageFormat = "%s. Visit https://docs.counterapi.dev for more information."
)

// addErrors is for error route group.
func (r Routes) addErrors() {
	r.router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": fmt.Sprintf(errorMessageFormat, "page not found"),
		})
	})

	r.router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    "METHOD_NOT_ALLOWED",
			"message": fmt.Sprintf(errorMessageFormat, "method not allowed"),
		})
	})
}
