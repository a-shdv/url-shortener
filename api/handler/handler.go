package handler

import (
	"github.com/a-shdv/url-shortener/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/a/:url", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "testId",
		})
	})

	router.GET("/s/:code", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	return router
}
