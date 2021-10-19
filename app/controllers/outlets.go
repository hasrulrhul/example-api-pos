package controllers

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexOutlet(c *gin.Context) {
	var outlet []models.Outlet
	config.DB.Preload("Merchant").Find(&outlet)
	c.JSON(http.StatusOK, service.Response(outlet, c, "", 0))
}

func CreateOutlet(c *gin.Context) {
	var form models.Outlet
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parseMerchantID, _ := strconv.ParseInt(c.PostForm("merchant_id"), 10, 64)
	input := models.Outlet{
		MerchantID: parseMerchantID,
		OutletName: c.PostForm("outlet_name"),
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "created data successfully")
	}
}

func ShowOutlet(c *gin.Context) {
	id := c.Params.ByName("id")
	var outlet models.Outlet
	err := config.DB.Preload("Merchant").First(&outlet, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
	} else {
		c.JSON(http.StatusOK, outlet)
	}
}

func UpdateOutlet(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Outlet
	data := config.DB.First(&product, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}

	var form models.Outlet
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parseMerchantID, _ := strconv.ParseInt(c.PostForm("merchant_id"), 10, 64)
	input := models.Outlet{
		MerchantID: parseMerchantID,
		OutletName: c.PostForm("outlet_name"),
	}

	if err := data.Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "updated data successfully")
	}
}

func DeleteOutlet(c *gin.Context) {
	id := c.Params.ByName("id")
	var outlet models.Outlet
	err := config.DB.First(&outlet, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&outlet).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "deleted data successfully")
	}
}
