package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/list", GetCoffees)
		v1.GET("/price/:name", GetPrice)
	}

	r.Run(":5000")

}
