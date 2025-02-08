package controllers

import (
	"manajemen-inventaris/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pc.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	var products []models.Product
	pc.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func (h *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if result := h.DB.First(&product, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func (h *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	h.DB.Delete(&models.Product{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Produk dihapus"})
}
