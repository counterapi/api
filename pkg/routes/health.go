package routes

import (
	"github.com/counterapi/api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

// addHealth is for health route group.
func (r Routes) addHealth(rg *gin.RouterGroup) {
	route := rg.Group("/health")

	health := new(controllers.HealthController)

	route.GET("/",
		health.Status,
	)
}
