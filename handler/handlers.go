package handler

import (
  "fmt"
  "net/http"
  "github.com/RJHarvest/url-shortener/shortener"
  "github.com/RJHarvest/url-shortener/store"
  "github.com/gin-gonic/gin"
)

// Request model definition
type UrlCreationRequest struct {
  LongUrl string `json:"long_url" binding:"required"`
  UserId string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
  var creationRequest UrlCreationRequest
  if err := c.ShouldBindJSON(&creationRequest); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
  store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

  host := fmt.Sprintf("http://%s/", c.Request.Host)
  c.JSON(200, gin.H{
    "message": "short url created successfully",
    "short_url": host + shortUrl,
  })
}

func HandleShortUrlRedirect(c *gin.Context) {
  shortUrl := c.Param("shortUrl")
  initialUrl := store.RetrieveInitialUrl(shortUrl)
  c.Redirect(302, initialUrl)
}
