package routes

import (
	"github.com/counterapi/counter/config"
	"github.com/counterapi/counter/controllers"
	"github.com/counterapi/counter/repositories"
	"github.com/gin-gonic/gin"
)

func addCounterRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("")

	counter := controllers.CounterController{
		Repository: repositories.CounterRepository{DB: config.DB},
	}

	ping.GET("/up", counter.Up)
	ping.GET("/down", counter.Down)
	ping.GET("/get", counter.Get)
}
