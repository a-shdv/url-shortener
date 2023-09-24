package main

import (
	"github.com/a-shdv/url-shortener/api/handler"
	"github.com/a-shdv/url-shortener/api/repo"
	"github.com/a-shdv/url-shortener/api/service"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// env variables loading
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// db creation
	redisDb := repo.NewRedisDb(0)
	defer redisDb.Close()

	// server startup
	repositories := repo.NewRepository(redisDb)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	log.Println("App has been started...")
	srv := new(Server)
	go func() {
		if err = srv.Run(os.Getenv("SERVER_ADDR"), handlers.InitRoutes()); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	// server shut down
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("App is shutting down...")
}
