package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Printf(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
