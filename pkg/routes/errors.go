package routes

import (
	"fmt"
	"net/http"

	"github.com/counterapi/api/pkg"

	"github.com/gin-gonic/gin"
)

// addErrors is for error route group.
func (r Routes) addErrors() {
	r.router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "404",
			"message": fmt.Sprintf(pkg.ErrorMessageFormat, "page not found"),
		})
	})

	r.router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    "405",
			"message": fmt.Sprintf(pkg.ErrorMessageFormat, "method not allowed"),
		})
	})
}
