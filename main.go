package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"korie/api-product/app"
	"korie/api-product/controller"
	"korie/api-product/repository"
	"korie/api-product/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	ProductRepository := repository.NewProductRepository()
	productService := service.NewProductService(ProductRepository, db, validate)
	productController := controller.NewProductController(productService)

	r := gin.Default()

	r.GET("api/products", productController.FindAll)       // GET ALL
	r.POST("api/products", productController.Create)       // CREATE
	r.PUT("api/products/:id", productController.Update)    // UPDATE
	r.DELETE("api/products/:id", productController.Delete) // DELETE
	r.GET("api/products/:id", productController.FindById)  // GET BY ID

	r.Run("localhost:3000")
}
