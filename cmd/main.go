package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// env variables loading
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// server launching
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	r.POST("/test/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "testId",
		})
	})

	go func() {
		err = r.Run(":" + os.Getenv("PORT"))
		if err != nil {
			log.Fatalf(err.Error())
		}
	}()

}
