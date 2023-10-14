package routes

import (
	"os"

	"github.com/counterapi/api/pkg/config"
	"github.com/counterapi/api/pkg/controllers"
	"github.com/counterapi/api/pkg/repositories"

	"github.com/gin-gonic/gin"
)

// addMetric is for metrics route group.
func (r Routes) addMetric(rg *gin.RouterGroup) {
	metricsAuthUser := func(c *gin.Context) {}

	if enableBasicAuth := os.Getenv("METRICS_BASIC_AUTH"); enableBasicAuth == "true" {
		metricsAuthUser = gin.BasicAuth(gin.Accounts{"admin": "admin"})

		if username := os.Getenv("METRICS_BASIC_AUTH_USERNAME"); username != "" {
			if password := os.Getenv("METRICS_BASIC_AUTH_PASSWORD"); password != "" {
				metricsAuthUser = gin.BasicAuth(gin.Accounts{username: password})
			}
		}
	}

	route := rg.Group("/metrics", metricsAuthUser)

	ctrl := controllers.NewMetricController(
		repositories.CounterRepository{DB: config.DB},
	)

	route.GET("/",
		ctrl.Serve,
	)
}
