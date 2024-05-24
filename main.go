package main

import (
	"fmt"
	"net/http"

	"github.com/dn46/go-url-shortener/handler"
	"github.com/dn46/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Welcome to the url shortener API!",
		// })
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome to the url shortener!",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - ERROR: %v", err))
	}
}
