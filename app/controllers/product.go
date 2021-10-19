package controllers

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func IndexProduct(c *gin.Context) {
	var product []models.Product
	config.DB.Find(&product)
	c.JSON(http.StatusOK, service.Response(product, c, "", 0))
}

func CreateProduct(c *gin.Context) {
	var form models.Product
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	price := c.PostForm("price")
	parsePrice, _ := strconv.ParseInt(price, 10, 64)

	// upload file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	if err = c.SaveUploadedFile(file, "uploads/"+newFileName); err != nil {
		c.JSON(http.StatusBadRequest, "failed")
		return
	}

	input := models.Product{
		ProductName: c.PostForm("product_name"),
		Sku:         c.PostForm("sku"),
		Category:    c.PostForm("category"),
		Price:       parsePrice,
		Image:       newFileName,
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "created data successfully")
	}
}

func ShowProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product
	err := config.DB.First(&product, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product
	data := config.DB.First(&product, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}

	var form models.Product
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsePrice, _ := strconv.ParseInt(c.PostForm("price"), 10, 64)

	// upload file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	if err = c.SaveUploadedFile(file, "uploads/"+newFileName); err != nil {
		c.JSON(http.StatusBadRequest, "failed")
		return
	}

	if file != nil {
		os.Remove("uploads/" + product.Image)
	}

	input := models.Product{
		ProductName: c.PostForm("product_name"),
		Sku:         c.PostForm("sku"),
		Category:    c.PostForm("category"),
		Price:       parsePrice,
		Image:       newFileName,
	}

	if err := data.Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "updated data successfully")
	}
}

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product
	err := config.DB.First(&product, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}
	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		os.Remove("uploads/" + product.Image)
		c.JSON(http.StatusOK, "deleted data successfully")
	}
}
