package routes

import (
	"github.com/counterapi/counter/pkg/config"
	"github.com/counterapi/counter/pkg/controllers"
	"github.com/counterapi/counter/pkg/repositories"

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
