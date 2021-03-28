package main

import (
	"github.com/counterapi/counterapi/pkg/routes"
)

func main() {
	r := routes.NewRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
