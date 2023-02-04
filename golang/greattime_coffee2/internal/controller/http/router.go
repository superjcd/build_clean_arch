package http

import (
	"greattime_coffee/internal/usecase"
	"greattime_coffee/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.Interface, s usecase.CoffeeService) {
	// middleware
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")
	{
		newCoffeeServiceRoutes(h, s, l)
	}
}
