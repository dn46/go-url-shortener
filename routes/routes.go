package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dn46/go-url-shortener/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome to the url shortener!",
		})
	})

	r.POST("/create-short-url", handler.CreateShortUrl)

	r.GET("/:shortUrl", handler.HandleShortUrlRedirect)
}
