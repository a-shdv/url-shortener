package handler

import (
	"github.com/a-shdv/url-shortener/api/helper"
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) createShortUrl(c *gin.Context) {
	var request *model.Url

	if err := c.BindJSON(&request); err != nil {
		helper.NewErrorResponse(c, http.StatusBadRequest, "cannot parse json!")
		return
	}

	if reqUrl := helper.ParseUrlAddr(request.OriginalUrl); !govalidator.IsURL(reqUrl) {
		helper.NewErrorResponse(c, http.StatusBadRequest, "wrong url format!")
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

func (h *Handler) getOriginalUrl(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		log.Printf("code provided in url is empty!")
		return
	}

	url := h.service.UrlService.GetOriginalUrl(code)

	c.Redirect(http.StatusFound, "https://www."+url)
}
