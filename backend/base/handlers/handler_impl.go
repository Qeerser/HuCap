package handlers

import "HuCap/base/services"

type handler struct {
	userHandler *userHandler
}

func NewHandler(service services.Service) Handler {
	return &handler{
		userHandler: NewUserHandler(service.User()),
	}
}

func (h *handler) User() *userHandler{
	return h.userHandler
}