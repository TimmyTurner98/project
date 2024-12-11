package handler

import "github.com/TimmyTurner98/project/pkg/service"

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
