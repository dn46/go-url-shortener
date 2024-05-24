package main

import (
	"fmt"

	"github.com/dn46/go-url-shortener/routes"
	"github.com/dn46/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.SetupRoutes(r)

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - ERROR: %v", err))
	}
}
