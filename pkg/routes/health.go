package routes

import (
	"github.com/counterapi/counter/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func addHealthRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("health")

	health := new(controllers.HealthController)

	ping.GET("/", health.Status)
}
