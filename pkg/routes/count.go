package routes

import (
	"github.com/counterapi/counter/pkg/config"
	"github.com/counterapi/counter/pkg/controllers"
	"github.com/counterapi/counter/pkg/repositories"
	"github.com/gin-gonic/gin"
)

func addCountRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("counts")

	counter := controllers.CountController{
		Repository: repositories.CountRepository{DB: config.DB},
	}

	ping.GET("/", counter.GetCounts)
}
