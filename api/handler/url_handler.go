package handler

import (
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) createShortUrl(c *gin.Context) {
	var request model.Url
	err := c.BindJSON(&request)
	if err != nil {
		log.Fatalf(err.Error())
	}

	shortUrl := h.service.UrlService.CreateShortUrl(&request)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": shortUrl,
	})
}

func (h *Handler) getOriginalUrl(c *gin.Context) {

}
