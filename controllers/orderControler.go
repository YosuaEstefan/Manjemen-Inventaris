package controllers

import (
	"manajemen-inventaris/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler struct {
	DB *gorm.DB
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{DB: db}
}

type OrderCreate struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var orderData OrderCreate
	if err := c.ShouldBindJSON(&orderData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if result := h.DB.First(&product, orderData.ProductID); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	order := models.Order{
		ProductID: orderData.ProductID,
		Quantity:  orderData.Quantity,
		OrderDate: time.Now(),
	}
	h.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if result := h.DB.First(&order, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesanan tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, order)
}
