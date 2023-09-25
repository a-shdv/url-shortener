package handler

import (
	"github.com/a-shdv/url-shortener/api/helper"
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createShortUrl(c *gin.Context) {
	var request *model.Url

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "cannot parse json!",
		})
		return
	}

	reqUrl := helper.ParseUrlAddr(request.OriginalUrl)
	if !govalidator.IsURL(reqUrl) {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "wrong url format!",
		})
		return
	}

	shortUrl, err := h.service.UrlService.CreateShortUrl(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "this url is already in database!",
			"code":  shortUrl,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": shortUrl,
	})
}
