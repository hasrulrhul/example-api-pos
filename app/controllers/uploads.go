package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	// Upload the file to specific dst.
	if err = c.SaveUploadedFile(file, "uploads/"+file.Filename); err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
	// c.String(http.StatusOK, file.Filename)
}

func UploadFile2(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// Upload the file to specific dst.
	if err = c.SaveUploadedFile(file, "uploads/"+newFileName); err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func DeleteFile(c *gin.Context) {
	id := c.Params.ByName("id")
	var path = "uploads/" + id
	var err = os.Remove(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "success")
	}
}
