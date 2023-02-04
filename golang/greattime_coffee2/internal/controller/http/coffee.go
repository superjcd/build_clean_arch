package http

import (
	"greattime_coffee/internal/entity"
	"greattime_coffee/internal/usecase"
	"greattime_coffee/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type coffeeServiceRoutes struct {
	cs usecase.CoffeeService
	l  logger.Interface
}

type productsResponse struct {
	Products []entity.Coffee `json:"products"`
}

type coffeePriceResponse struct {
	Price float64 `json:"price"`
}

func newCoffeeServiceRoutes(handler *gin.RouterGroup, cs usecase.CoffeeService, l logger.Interface) {
	r := &coffeeServiceRoutes{cs, l}

	h := handler.Group("/coffee")

	{
		h.GET("/list", r.list)
		h.GET("/price/:name", r.getPrice)
	}
}

func (r *coffeeServiceRoutes) list(c *gin.Context) {
	products, err := r.cs.List(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productsResponse{Products: products})

}

func (r *coffeeServiceRoutes) getPrice(c *gin.Context) {
	name := c.Param("name")

	price, err := r.cs.GetPrice(c.Request.Context(), name)
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coffeePriceResponse{Price: price})
}
