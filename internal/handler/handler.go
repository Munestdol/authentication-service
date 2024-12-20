package handler

import (
	"authentication-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use()

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.GET("/protected", h.Authenticate, h.ProtectedEndpoint)
	}

	return router
}
