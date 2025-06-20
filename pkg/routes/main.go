package routes

import (
	"github.com/counterapi/api/pkg/config"
	"github.com/counterapi/api/pkg/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Routes is main route struct.
type Routes struct {
	router      *gin.Engine
	cacheConfig *config.RedisCache
}

// NewRoutes generates Routes for the application.
func NewRoutes() Routes {
	r := Routes{
		router:      gin.Default(),
		cacheConfig: config.SetupRedisCache(),
	}

	setDB()
	setMiddlewares(r.router)

	main := r.router.Group("")
	v1 := r.router.Group("/v1")

	r.addHome(main)
	r.addHealth(main)
	r.addMetric(main)
	r.addCounter(v1)
	r.addErrors()

	return r
}

// Run runs application with routes.
func (r Routes) Run() error {
	return r.router.Run(":80")
}

// setDB will create Database instance.
func setDB() {
	db, err := config.SetupDatabase()
	if err != nil {
		panic(err)
	}

	config.DB = db
}

// setMiddlewares will set middlewares.
func setMiddlewares(router *gin.Engine) {
	lm := middlewares.NewRedisRateLimiter(func(ctx *gin.Context) (string, error) {
		return ctx.Request.URL.Path, nil
	})

	router.Use(lm.Middleware())
	router.Use(cors.Default())
}
