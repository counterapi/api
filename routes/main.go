package routes

import (
	"github.com/counterapi/counter/config"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run() {
	setDB()
	getRoutes()

	router.Run(":8080")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	main := router.Group("")
	v1 := router.Group("/v1")

	addHealthRoutes(main)
	addCounterRoutes(v1)
	addCountRoutes(v1)
}

// setDB will create Database instance
func setDB() {
	db, err := config.SetupDatabase()
	if err != nil {
		panic(err)
	}

	config.DB = db
}
