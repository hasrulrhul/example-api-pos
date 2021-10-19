package controllers

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexTrxPurchase(c *gin.Context) {
	var trxPurchase []models.TrxPurchase
	config.DB.Preload("Detail").Preload("Outlet.Merchant").Find(&trxPurchase)
	c.JSON(http.StatusOK, service.Response(trxPurchase, c, "", 0))
}

func CreateTrxPurchase(c *gin.Context) {
	var form models.TrxPurchase
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parseOutletID, _ := strconv.ParseInt(c.PostForm("outlet_id"), 10, 64)

	var cartPurchase []models.CartPurchase
	config.DB.Where("outlet_id = ?", parseOutletID).Find(&cartPurchase)

	var total int64
	config.DB.Raw("SELECT SUM(total_price) FROM cart_purchases WHERE outlet_id = ?", parseOutletID).Scan(&total)

	if total != 0 {
		data := models.TrxPurchase{
			OutletID:     parseOutletID,
			TotalPayment: total,
		}
		if err := config.DB.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, "failed")
		} else {
			PurchaseID := strconv.FormatInt(int64(data.ID), 10)
			parsePurchaseID, _ := strconv.ParseInt(PurchaseID, 10, 64)

			for _, val := range cartPurchase {
				detail := models.TrxDetailPurchase{
					PurchaseID: parsePurchaseID,
					OutletID:   val.OutletID,
					ProductID:  val.ProductID,
					Qty:        val.Qty,
					Price:      val.Price,
					TotalPrice: val.TotalPrice,
				}
				config.DB.Create(&detail)
			}
			config.DB.Delete(&cartPurchase)
			c.JSON(http.StatusOK, "checkout successfully")
		}
	}
}

func ShowTrxPurchase(c *gin.Context) {
	id := c.Params.ByName("id")
	var trxPurchase models.TrxPurchase
	err := config.DB.Preload("Detail").Preload("Outlet.Merchant").First(&trxPurchase, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
	} else {
		c.JSON(http.StatusOK, trxPurchase)
	}
}

func UpdateTrxPurchase(c *gin.Context) {
	id := c.Params.ByName("id")
	var trxPurchase models.TrxPurchase
	data := config.DB.First(&trxPurchase, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}

	input := models.TrxPurchase{
		Status: "paid",
	}

	if err := data.Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "payment successfully")
	}
}

func DeleteTrxPurchase(c *gin.Context) {
	id := c.Params.ByName("id")
	var trxPurchase models.TrxPurchase
	err := config.DB.First(&trxPurchase, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&trxPurchase).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "deleted data successfully")
	}
}

func ReportTrxPurchase(c *gin.Context) {
	id := c.Params.ByName("id")
	var trxPurchase models.TrxPurchase
	err := config.DB.First(&trxPurchase, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&trxPurchase).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "deleted data successfully")
	}
}

func ReportTrxPurchasePerProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var trxPurchase models.TrxPurchase
	err := config.DB.First(&trxPurchase, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&trxPurchase).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "deleted data successfully")
	}
}
