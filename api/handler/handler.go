package handler

import (
	"github.com/a-shdv/url-shortener/api/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/a/", h.createShortUrl)

	router.GET("/s/:code", h.getOriginalUrl)

	return router
}
