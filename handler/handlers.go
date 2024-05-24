package handler

import (
	"net/http"

	"github.com/dn46/go-url-shortener/shortener"
	"github.com/dn46/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	// LongUrl string `json:"long_url" binding:"required"`
	// UserID  string `json:"user_id" binding:"required"`
	LongUrl string `form:"url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest

	if err := c.ShouldBind(&creationRequest); err != nil { // we remove ShouldBindJSON to display the correct format of the data
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortUrl, salt := shortener.GenerateShortLink(creationRequest.LongUrl)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, salt)

	host := "http://localhost:9808/"
	// JSON response below
	// c.JSON(200, gin.H{
	// 	"message":   "short url created successfully",
	// 	"short_url": host + shortUrl,
	// })

	// HTML response
	c.HTML(http.StatusOK, "result.html", gin.H{
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl) // so we redirect the short url generated to the initial long one provided
}
