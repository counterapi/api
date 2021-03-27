package routes

import (
	"github.com/counterapi/counter/pkg/config"
	"github.com/counterapi/counter/pkg/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Run will start the server.
func Run() {
	router := gin.Default()

	setDB()
	setMiddlewares(router)
	getRoutes(router)

	err := router.Run(":80")
	if err != nil {
		panic(err)
	}
}

// getRoutes will create our routes of our entire application.
func getRoutes(router *gin.Engine) {
	main := router.Group("")
	v1 := router.Group("/v1")

	addHealthRoutes(main)
	addCounterRoutes(v1)
	addCountRoutes(v1)
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
	lm := middlewares.NewRateLimiter(func(ctx *gin.Context) (string, error) {
		return ctx.ClientIP(), nil
	})

	router.Use(lm.Middleware())
	router.Use(cors.Default())
}
