package routes

import (
	"github.com/counterapi/api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

// addHome is for home route group.
func (r Routes) addHome(rg *gin.RouterGroup) {
	route := rg.Group("/")

	home := new(controllers.HomeController)

	route.GET("/", home.Redirect)
}
