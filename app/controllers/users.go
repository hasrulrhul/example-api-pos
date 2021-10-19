package controllers

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexUser(c *gin.Context) {
	var user []models.User
	config.DB.Find(&user)
	c.JSON(http.StatusOK, service.Response(user, c, "", 0))
}

func CreateUser(c *gin.Context) {
	var form models.User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := service.HashPassword(c.PostForm("password"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed to encription")
		return
	}

	input := models.User{
		Name:     c.PostForm("name"),
		Username: c.PostForm("username"),
		Email:    c.PostForm("email"),
		Password: hashedPassword,
		Role:     c.PostForm("role"),
		Type:     c.PostForm("type"),
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "created data successfully")
	}
}

func ShowUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := config.DB.First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	data := config.DB.First(&user, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}

	var form models.User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := service.HashPassword(c.PostForm("password"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed to encription")
		return
	}

	input := models.User{
		Name:     c.PostForm("name"),
		Username: c.PostForm("username"),
		Email:    c.PostForm("email"),
		Password: hashedPassword,
		Role:     c.PostForm("role"),
		Type:     c.PostForm("type"),
	}

	if err := data.Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "updated data successfully")
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := config.DB.First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "deleted data successfully")
	}
}

func UploadUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := config.DB.First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "successfull")
	}
}
