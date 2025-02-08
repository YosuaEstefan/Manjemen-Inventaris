package controllers

import (
	"manajemen-inventaris/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InventarisController struct {
	DB *gorm.DB
}

func NewInventarisController(db *gorm.DB) *InventarisController {
	return &InventarisController{DB: db}
}

func (h *InventarisController) GetInventoryByProduct(c *gin.Context) {
	productID := c.Param("id")
	var inventory []models.Inventory
	h.DB.Where("product_id = ?", productID).Find(&inventory)
	c.JSON(http.StatusOK, inventory)
}

type InventoryUpdate struct {
	Location string `json:"location"`
	Delta    int    `json:"delta"`
}

func (h *InventarisController) UpdateInventory(c *gin.Context) {
	productID := c.Param("id")
	var update InventoryUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var inv models.Inventory
	result := h.DB.Where("product_id = ? AND location = ?", productID, update.Location).First(&inv)
	if result.Error == gorm.ErrRecordNotFound {
		inv = models.Inventory{
			ProductID: parseUint(productID),
			Location:  update.Location,
			Quantity:  update.Delta,
		}
		h.DB.Create(&inv)
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	} else {
		inv.Quantity += update.Delta
		h.DB.Save(&inv)
	}
	c.JSON(http.StatusOK, inv)
}

func parseUint(s string) uint {
	i, _ := strconv.ParseUint(s, 10, 32)
	return uint(i)
}
