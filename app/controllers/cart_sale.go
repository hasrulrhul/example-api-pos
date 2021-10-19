package controllers

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexCartSale(c *gin.Context) {
	var cart []models.CartSale
	config.DB.Preload("Customer").Preload("Product").Preload("Outlet.Merchant").Find(&cart)
	c.JSON(http.StatusOK, service.Response(cart, c, "", 0))
}

func CreateCartSale(c *gin.Context) {
	var form models.CartSale
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parseCustomerID, _ := strconv.ParseInt(c.PostForm("customer_id"), 10, 64)
	parseProductID, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 64)
	parseOutletID, _ := strconv.ParseInt(c.PostForm("outlet_id"), 10, 64)
	parseQty, _ := strconv.ParseInt(c.PostForm("qty"), 10, 64)
	parsePrice, _ := strconv.ParseInt(c.PostForm("price"), 10, 64)

	data := models.CartSale{
		CustomerID: parseCustomerID,
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

func ShowCartSale(c *gin.Context) {
	id := c.Params.ByName("id")
	var cart models.CartSale
	err := config.DB.Preload("Customer").Preload("Product").Preload("Outlet.Merchant").First(&cart, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
	} else {
		c.JSON(http.StatusOK, cart)
	}
}

func UpdateCartSale(c *gin.Context) {
	id := c.Params.ByName("id")
	var cart models.CartSale
	data := config.DB.First(&cart, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}
	var form models.CartSale
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parseCustomerID, _ := strconv.ParseInt(c.PostForm("customer_id"), 10, 64)
	parseProductID, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 64)
	parseOutletID, _ := strconv.ParseInt(c.PostForm("outlet_id"), 10, 64)
	parseQty, _ := strconv.ParseInt(c.PostForm("qty"), 10, 64)
	parsePrice, _ := strconv.ParseInt(c.PostForm("price"), 10, 64)

	input := models.CartSale{
		CustomerID: parseCustomerID,
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

func DeleteCartSale(c *gin.Context) {
	id := c.Params.ByName("id")
	var cart models.CartSale
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
