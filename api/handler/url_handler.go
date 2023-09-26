package handler

import (
	"github.com/a-shdv/url-shortener/api/helper"
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

// createShortUrl метод, отвечающий за создание короткого url.
func (h *Handler) createShortUrl(c *gin.Context) {
	var request *model.Url

	// проверка json входящего запроса.
	if err := c.BindJSON(&request); err != nil {
		helper.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// парсинг url-адреса входящего запроса, проверяет на лишние слова, вроде "https://" или "www.".
	reqUrl := helper.ParseUrlAddr(request.OriginalUrl)
	if !govalidator.IsURL(reqUrl) {
		helper.NewErrorResponse(c, http.StatusBadRequest, "wrong url format!")
		return
	}
	request.OriginalUrl = reqUrl

	// создание короткого url-адреса.
	shortUrl, err := h.service.UrlService.CreateShortUrl(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
			"code":  shortUrl,
		})
		return
	}

	// ответ в виде json, выводящий уже укороченный url-адрес.
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": shortUrl,
	})
}

// getOriginalUrl метод, отвечающий за редирект на другую страницу с помощью укороченного url-адреса.
func (h *Handler) getOriginalUrl(c *gin.Context) {
	// извлечение параметра, введённого пользователем.
	code := c.Param("code")
	if code == "" {
		helper.NewErrorResponse(c, http.StatusBadRequest, "code provided in url is empty!")
		return
	}

	// поиск исходного url-адреса c помощью укороченного url-адреса.
	url := h.service.UrlService.GetOriginalUrlByCode(code)

	// редирект на страницу, указанную при создании короткой ссылки.
	c.Redirect(http.StatusFound, "https://www."+url)
}
