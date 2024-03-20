package controllers

import (
	"assigment2golang/config"
	"assigment2golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateItem membuat item baru
func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&item)
	c.JSON(http.StatusCreated, gin.H{"data": item})
}

// GetItems mengambil semua item
func GetItems(c *gin.Context) {
	var items []models.Item
	config.DB.Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

// UpdateItem memperbarui sebuah item
func UpdateItem(c *gin.Context) {
	var item models.Item
	if err := config.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updateInput models.Item
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&item).Updates(updateInput)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

// DeleteItem menghapus sebuah item
func DeleteItem(c *gin.Context) {
	var item models.Item
	if err := config.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
