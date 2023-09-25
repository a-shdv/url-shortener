package main

import (
	"context"
	"net/http"
	"os"
	"time"
)

// Server структура.
type Server struct {
	httpServer *http.Server
}

// Run метод, отвечающий за запуск сервера.
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           os.Getenv("SERVER_ADDR"),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown метод, отвечающий за выключение сервера.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
