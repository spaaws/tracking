package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Order struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	TrackingCode string `json:"tracking_code" gorm:"unique"`
	UserID       uint   `json:"user_id"`
}

func CreateOrder(c *gin.Context, db *gorm.DB) {
	var order Order

	// Bind JSON data to the order variable
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if tracking code already exists
	var existingOrder Order
	if db.Where("tracking_code = ?", order.TrackingCode).First(&existingOrder).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tracking code already exists"})
		return
	}

	// Create the new order
	db.Create(&order)
	c.JSON(http.StatusOK, order)
}


func GetOrder(c *gin.Context, db *gorm.DB) {
	var order Order
	if err := db.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrder(c *gin.Context, db *gorm.DB) {
	var order Order
	if err := db.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&order)
	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context, db *gorm.DB) {
	var order Order
	if err := db.Delete(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}

func GetOrdersByUserID(c *gin.Context, db *gorm.DB) {
	var orders []Order
	if err := db.Where("user_id = ?", c.Param("id")).Find(&orders).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No orders found for this user"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
