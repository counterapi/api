package routes

import (
	"github.com/counterapi/counter/config"
	"github.com/counterapi/counter/controllers"
	"github.com/counterapi/counter/repositories"
	"github.com/gin-gonic/gin"
)

func addCountRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("counts")

	counter := controllers.CountController{
		Repository: repositories.CountRepository{DB: config.DB},
	}

	ping.GET("/", counter.GetCounts)
}
