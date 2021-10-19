package controllers

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexCartPurchase(c *gin.Context) {
	var cart []models.CartPurchase
	config.DB.Preload("Product").Preload("Outlet.Merchant").Find(&cart)
	c.JSON(http.StatusOK, service.Response(cart, c, "", 0))
}

func CreateCartPurchase(c *gin.Context) {
	var form models.CartPurchase
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parseProductID, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 64)
	parseOutletID, _ := strconv.ParseInt(c.PostForm("outlet_id"), 10, 64)
	parseQty, _ := strconv.ParseInt(c.PostForm("qty"), 10, 64)
	parsePrice, _ := strconv.ParseInt(c.PostForm("price"), 10, 64)

	data := models.CartPurchase{
		ProductID:  parseProductID,
		OutletID:   parseOutletID,
		Qty:        parseQty,
		Price:      parsePrice,
		TotalPrice: (parseQty * parsePrice),
	}
	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "created data successfully")
	}
}

func ShowCartPurchase(c *gin.Context) {
	id := c.Params.ByName("id")
	var cart models.CartPurchase
	err := config.DB.Preload("Product").Preload("Outlet.Merchant").First(&cart, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
	} else {
		c.JSON(http.StatusOK, cart)
	}
}

func UpdateCartPurchase(c *gin.Context) {
	id := c.Params.ByName("id")
	var cart models.CartPurchase
	data := config.DB.First(&cart, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}
	var form models.CartPurchase
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parseProductID, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 64)
	parseOutletID, _ := strconv.ParseInt(c.PostForm("outlet_id"), 10, 64)
	parseQty, _ := strconv.ParseInt(c.PostForm("qty"), 10, 64)
	parsePrice, _ := strconv.ParseInt(c.PostForm("price"), 10, 64)

	input := models.CartPurchase{
		ProductID:  parseProductID,
		OutletID:   parseOutletID,
		Qty:        parseQty,
		Price:      parsePrice,
		TotalPrice: (parseQty * parsePrice),
	}

	if err := data.Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "updated data successfully")
	}
}

func DeleteCartPurchase(c *gin.Context) {
	id := c.Params.ByName("id")
	var cart models.CartPurchase
	err := config.DB.First(&cart, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "deleted data successfully")
	}
}
