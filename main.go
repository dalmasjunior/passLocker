package main

import (
	"github.com/dalmasjunior/passLocker/router"
)

func main() {
	cr := router.NewCustomRouter()

	cr.SetupRoutes()

	cr.Listen(":8000")
}
