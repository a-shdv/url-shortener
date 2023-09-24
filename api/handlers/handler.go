package handlers

import "github.com/a-shdv/url-shortener/api/services"

type Handler struct {
	service *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{}
}
