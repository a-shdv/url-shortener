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
	// загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// создание клиента redis db
	redisDb := repo.NewRedisDb(0)
	defer redisDb.Close()

	// внедрение зависимостей
	repositories := repo.NewRepository(redisDb)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	// запуск сервера
	log.Println("url-shortener has been started...")
	srv := new(Server)
	go func() {
		if err = srv.Run(os.Getenv("SERVER_ADDR"), handlers.InitRoutes()); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	// выключение сервера
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("url-shortener is shutting down...")
}
