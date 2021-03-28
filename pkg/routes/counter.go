package routes

import (
	"github.com/counterapi/counterapi/pkg/config"
	"github.com/counterapi/counterapi/pkg/controllers"
	"github.com/counterapi/counterapi/pkg/repositories"

	"github.com/gin-gonic/gin"
)

// addCounter is for counter route group.
func (r Routes) addCounter(rg *gin.RouterGroup) {
	route := rg.Group("/")

	counter := controllers.CounterController{
		Repository: repositories.CounterRepository{DB: config.DB},
	}

	route.GET("/up", counter.Up)
	route.GET("/down", counter.Down)
	route.GET("/get", counter.Get)
}
