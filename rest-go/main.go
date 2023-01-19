package main

import (
	"github.com/gin-gonic/gin"

	"rest-go/controllers"

	"rest-go/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/products", controllers.FindAllProduct)
	r.GET("/products/:id", controllers.FindOneProduct)
	r.POST("/products", controllers.AddProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products", controllers.DeleteProduct)

	r.Run()
}
