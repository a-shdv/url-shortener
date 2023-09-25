package helper

import (
	"github.com/gin-gonic/gin"
	"log"
)

// ErrorResponse структура.
type ErrorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse конструктор.
func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
	log.Fatalf(message)
}
