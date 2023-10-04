package routes

import (
	"github.com/counterapi/counterapi/pkg/config"
	"github.com/counterapi/counterapi/pkg/controllers"
	"github.com/counterapi/counterapi/pkg/repositories"

	"github.com/gin-gonic/gin"
)

// addCounter is for counter's route group.
func (r Routes) addCounter(rg *gin.RouterGroup) {
	route := rg.Group("/:namespace/:counter/")

	counter := controllers.CounterController{
		Repository: repositories.CounterRepository{DB: config.DB},
	}

	route.GET("/", counter.Get)
	route.GET("/up", counter.Up)
	route.GET("/down", counter.Down)
	route.GET("/set", counter.Set)
}
