package controllers

import (
	"crud_api/database"
	"crud_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProductInWork(c *gin.Context) {
	var productInWork models.ProductInWork
	if err := c.ShouldBindJSON(&productInWork); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&productInWork)
	c.JSON(http.StatusOK, productInWork)
}
func GetAllProductsInWork(c *gin.Context) {
	var piw []models.ProductInWork
	database.DB.Preload("User").Preload("Product").Find(&piw)
	c.JSON(http.StatusOK, piw)
}

func GetProductInWorkByID(c *gin.Context) {
	id := c.Param("id")
	var piw models.ProductInWork
	if err := database.DB.Preload("User").Preload("Product").First(&piw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}
	c.JSON(http.StatusOK, piw)
}

// Обновление связи
func UpdateProductInWork(c *gin.Context) {
	id := c.Param("id")
	var piw models.ProductInWork
	if err := database.DB.First(&piw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}
	var input models.ProductInWork
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	piw.UserID = input.UserID
	piw.ProductID = input.ProductID
	database.DB.Save(&piw)
	c.JSON(http.StatusOK, piw)
}

// Удаление связи
func DeleteProductInWork(c *gin.Context) {
	id := c.Param("id")
	var piw models.ProductInWork
	if err := database.DB.First(&piw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}
	database.DB.Delete(&piw)
	c.JSON(http.StatusOK, gin.H{"message": "Entry deleted"})
}
