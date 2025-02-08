package main

import (
	"log"
	"manajemen-inventaris/config"
	"manajemen-inventaris/controllers"
	"manajemen-inventaris/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading ENV")
	}

	db := config.ConnectDatabase()
	db.AutoMigrate(&models.Product{}, &models.Inventory{}, &models.Order{})

	product := controllers.NewProductController(db)
	inventaris := controllers.NewInventarisController(db)
	order := controllers.NewOrderHandler(db)
	r := gin.Default()
	r.POST("/products", product.CreateProduct)
	r.GET("/products", product.GetAllProducts)
	r.PUT("/products/:id", product.UpdateProduct)
	r.DELETE("/products/:id", product.DeleteProduct)
	r.GET("/inventories/:id", inventaris.GetInventoryByProduct)
	r.POST("/inventories/:id", inventaris.UpdateInventory)
	r.POST("/orders", order.CreateOrder)
	r.GET("/orders/:id", order.GetOrder)

	r.Run(":8080")
}
