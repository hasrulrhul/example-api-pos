package controllers

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexProductOutlet(c *gin.Context) {
	var product_outlet []models.ProductOutlet
	config.DB.Preload("Product").Preload("Outlet.Merchant").Find(&product_outlet)
	c.JSON(http.StatusOK, service.Response(product_outlet, c, "", 0))
}

func CreateProductOutlet(c *gin.Context) {
	var form models.ProductOutlet
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parseProductID, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 64)
	parseOutletID, _ := strconv.ParseInt(c.PostForm("outlet_id"), 10, 64)
	parseSellingPrice, _ := strconv.ParseInt(c.PostForm("selling_price"), 10, 64)
	parseQty, _ := strconv.ParseInt(c.PostForm("qty"), 10, 64)

	input := models.ProductOutlet{
		ProductID:    parseProductID,
		OutletID:     parseOutletID,
		Qty:          parseQty,
		SellingPrice: parseSellingPrice,
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "created data successfully")
	}
}

func ShowProductOutlet(c *gin.Context) {
	id := c.Params.ByName("id")
	var product_outlet models.ProductOutlet
	err := config.DB.Preload("Product").Preload("Outlet.Merchant").First(&product_outlet, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
	} else {
		c.JSON(http.StatusOK, product_outlet)
	}
}

func UpdateProductOutlet(c *gin.Context) {
	id := c.Params.ByName("id")
	var product_outlet models.ProductOutlet
	data := config.DB.First(&product_outlet, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}

	var form models.ProductOutlet
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parseProductID, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 64)
	parseOutletID, _ := strconv.ParseInt(c.PostForm("outlet_id"), 10, 64)
	parseSellingPrice, _ := strconv.ParseInt(c.PostForm("selling_price"), 10, 64)
	parseQty, _ := strconv.ParseInt(c.PostForm("qty"), 10, 64)

	input := models.ProductOutlet{
		ProductID:    parseProductID,
		OutletID:     parseOutletID,
		Qty:          parseQty,
		SellingPrice: parseSellingPrice,
	}

	if err := data.Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "updated data successfully")
	}
}

func DeleteProductOutlet(c *gin.Context) {
	id := c.Params.ByName("id")
	var product_outlet models.ProductOutlet
	err := config.DB.First(&product_outlet, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}
	if err := config.DB.Delete(&product_outlet).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "deleted data successfully")
	}
}

func UploadProductOutlet(c *gin.Context) {
	id := c.Params.ByName("id")
	var product_outlet models.ProductOutlet
	err := config.DB.First(&product_outlet, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}
	if err := config.DB.Delete(&product_outlet).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "successfull")
	}
}
