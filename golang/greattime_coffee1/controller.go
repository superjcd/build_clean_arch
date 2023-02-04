package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCoffees(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"data":  AllCoffees,
		"total": len(AllCoffees),
	})
}

func GetPrice(c *gin.Context) {
	coffeeQuery := c.Param("name")

	var Target Coffee

	for _, coffee := range AllCoffees {
		// Find the first coffee
		if coffee.Name == coffeeQuery {
			Target = coffee
			break
		}
	}

	if Target.Price > 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": Target.Price,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "coffee not found",
			"data": -1,
		})
	}

}
