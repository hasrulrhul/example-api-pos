package controllers

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexCustomer(c *gin.Context) {
	var customer []models.Customer
	config.DB.Find(&customer)
	c.JSON(http.StatusOK, service.Response(customer, c, "", 0))
}

func CreateCustomer(c *gin.Context) {
	var form models.Customer
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := models.Customer{
		FirstName:   c.PostForm("first_name"),
		LastName:    c.PostForm("last_name"),
		Address:     c.PostForm("address"),
		Email:       c.PostForm("email"),
		PhoneNumber: c.PostForm("phone_number"),
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "created data successfully")
	}
}

func ShowCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer models.Customer
	err := config.DB.First(&customer, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

func UpdateCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer models.Customer
	data := config.DB.First(&customer, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}

	var form models.Customer
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := models.Customer{
		FirstName:   c.PostForm("first_name"),
		LastName:    c.PostForm("last_name"),
		Address:     c.PostForm("address"),
		Email:       c.PostForm("email"),
		PhoneNumber: c.PostForm("phone_number"),
	}

	if err := data.Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "updated data successfully")
	}
}

func DeleteCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer models.Customer
	err := config.DB.First(&customer, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}
	if err := config.DB.Delete(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "deleted data successfully")
	}
}

func UploadCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer models.Customer
	err := config.DB.First(&customer, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}
	if err := config.DB.Delete(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "successfull")
	}
}
