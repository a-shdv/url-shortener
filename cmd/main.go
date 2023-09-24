package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	app := fiber.New()

	err = app.Listen(os.Getenv("URL_ADDR"))
	if err != nil {
		log.Fatalf(err.Error())
	}
}
