package main

import (
	"github.com/counterapi/api/pkg/routes"
)

func main() {
	r := routes.NewRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
