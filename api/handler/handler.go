package handler

import (
	"github.com/a-shdv/url-shortener/api/service"
	"github.com/gin-gonic/gin"
)

// Handler структура.
type Handler struct {
	service *service.Service
}

// NewHandler конструктор.
func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// InitRoutes инициализация маршрутов.
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/a/", h.createShortUrl)
	router.GET("/s/:code", h.getOriginalUrl)

	return router
}
