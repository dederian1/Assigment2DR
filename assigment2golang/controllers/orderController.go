package controllers

import (
	"assigment2golang/config"
	"assigment2golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetAllOrders mengambil semua pesanan
func GetAllOrders(c *gin.Context) {
	var orders []models.Order
	config.DB.Find(&orders)
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// CreateOrder membuat pesanan baru
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&order)
	c.JSON(http.StatusCreated, gin.H{"data": order})
}

// UpdateOrder memperbarui pesanan
func UpdateOrder(c *gin.Context) {
	// Dapatkan model jika ada
	var order models.Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validasi input
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perbarui model
	config.DB.Model(&order).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": order})
}

// DeleteOrder menghapus pesanan
func DeleteOrder(c *gin.Context) {
	// Dapatkan model jika ada
	var order models.Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&order)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

// HashPassword meng-hash password yang diberikan
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash membandingkan password dengan hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
