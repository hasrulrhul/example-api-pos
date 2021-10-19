package controllers

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexMerchant(c *gin.Context) {
	var mercant []models.Merchant
	config.DB.Find(&mercant)
	c.JSON(http.StatusOK, service.Response(mercant, c, "", 0))
}

func CreateMerchant(c *gin.Context) {
	var form models.Merchant
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := models.Merchant{
		MerchantName: c.PostForm("merchant_name"),
	}
	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "created data successfully")
	}
}

func ShowMerchant(c *gin.Context) {
	id := c.Params.ByName("id")
	var mercant models.Merchant
	err := config.DB.First(&mercant, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
	} else {
		c.JSON(http.StatusOK, mercant)
	}
}

func UpdateMerchant(c *gin.Context) {
	id := c.Params.ByName("id")
	var mercant models.Merchant
	data := config.DB.First(&mercant, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}
	var form models.Merchant
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := models.Merchant{
		MerchantName: c.PostForm("merchant_name"),
	}
	if err := data.Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "updated data successfully")
	}
}

func DeleteMerchant(c *gin.Context) {
	id := c.Params.ByName("id")
	var mercant models.Merchant
	err := config.DB.First(&mercant, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&mercant).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "deleted data successfully")
	}
}
