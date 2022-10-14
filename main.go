package main

import (
	"assignment2/controller"
	"assignment2/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	router := gin.Default()

	router.POST("/orders", controller.CreateOrder)
	router.GET("/orders", controller.GetAllOrder)
	router.GET("/orders/:orderId", controller.GetOrderByID)
	router.PUT("/orders/:orderId", controller.UpdateOrder)
	router.DELETE("/orders/:orderId", controller.DeleteOrder)

	log.Println("server is running at port ", database.APP_PORT)
	router.Run(database.APP_PORT)
}
