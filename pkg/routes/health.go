package routes

import (
	"github.com/counterapi/api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

// addHealth is for health route group.
func (r Routes) addHealth(rg *gin.RouterGroup) {
	route := rg.Group("/health")
	apiRoute := rg.Group("/v1")

	health := new(controllers.HealthController)

	route.GET("/",
		health.Status,
	)

	apiRoute.GET("/",
		health.Status,
	)
}
