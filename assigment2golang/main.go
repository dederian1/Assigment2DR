package main

import (
	"assigment2golang/config"
	"assigment2golang/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.POST("/order", controllers.CreateOrder)
	r.GET("/order", controllers.GetAllOrders)
	r.PUT("/order/:id", controllers.UpdateOrder)
	r.DELETE("/order/:id", controllers.DeleteOrder)

	r.Run(":8080")
}
