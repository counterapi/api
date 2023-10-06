package routes

import (
	"time"

	"github.com/counterapi/counterapi/pkg/config"
	"github.com/counterapi/counterapi/pkg/controllers"
	"github.com/counterapi/counterapi/pkg/repositories"

	cache "github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
)

// addCounter is for counter's route group.
func (r Routes) addCounter(rg *gin.RouterGroup) {
	route := rg.Group("/:namespace/:counter/")

	counter := controllers.CounterController{
		Repository: repositories.CounterRepository{DB: config.DB},
	}

	route.GET("/", cache.CacheByRequestURI(r.cacheConfig.Store, r.cacheConfig.DefaultCacheTime), counter.Get)
	route.GET("/up", counter.Up)
	route.GET("/down", counter.Down)
	route.GET("/set", counter.Set)
	route.GET("/list", cache.CacheByRequestURI(r.cacheConfig.Store, 1*time.Hour), counter.GetCounts)
}
