package handler

import (
	"net/http"

	"github.com/dn46/go-url-shortener/shortener"
	"github.com/dn46/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `form:"url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest

	if err := c.ShouldBind(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortUrl, salt := shortener.GenerateShortLink(creationRequest.LongUrl)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, salt)

	host := "http://localhost:9808/"

	// HTML response
	c.HTML(http.StatusOK, "result.html", gin.H{
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
