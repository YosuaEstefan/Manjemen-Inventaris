package routes

import (
	"manajemen-inventaris/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	productController := controllers.NewProductController(db)
	InventarisController := controllers.NewInventarisController(db)
	orderHandler := controllers.NewOrderHandler(db)
	r.POST("/products", productController.CreateProduct)
	r.GET("/products", productController.GetAllProducts)
	r.PUT("/products/:id", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)
	r.GET("/inventories/:id", InventarisController.GetInventoryByProduct)
	r.POST("/inventories/:id", InventarisController.UpdateInventory)
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)

	return r
}
