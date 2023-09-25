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

	if err := c.BindJSON(&request); err != nil {
		helper.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	reqUrl := helper.ParseUrlAddr(request.OriginalUrl)
	if !govalidator.IsURL(reqUrl) {
		helper.NewErrorResponse(c, http.StatusBadRequest, "wrong url format!")
		return
	}
	request.OriginalUrl = reqUrl

	shortUrl, err := h.service.UrlService.CreateShortUrl(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
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
		helper.NewErrorResponse(c, http.StatusBadRequest, "code provided in url is empty!")
		return
	}

	url := h.service.UrlService.GetOriginalUrlByCode(code)

	c.Redirect(http.StatusFound, "https://www."+url)
}
