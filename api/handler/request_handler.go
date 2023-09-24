package handler

import (
	"github.com/a-shdv/url-shortener/api/model"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) сreateShortUrl(c *gin.Context) {
	var request model.Request
	err := c.BindJSON(&request)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf("%s", err)

	h.service.Request.CreateShortUrl(request)
}
