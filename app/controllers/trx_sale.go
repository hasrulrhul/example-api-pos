package controllers

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexTrxSale(c *gin.Context) {
	var trxSale []models.TrxSale
	config.DB.Preload("Customer").Preload("Detail").Preload("Outlet.Merchant").Find(&trxSale)
	c.JSON(http.StatusOK, service.Response(trxSale, c, "", 0))
}

func CreateTrxSale(c *gin.Context) {
	var form models.TrxSale
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parseCustomerID, _ := strconv.ParseInt(c.PostForm("customer_id"), 10, 64)
	parseOutletID, _ := strconv.ParseInt(c.PostForm("outlet_id"), 10, 64)

	var CartSale []models.CartSale
	config.DB.Where("customer_id = ?", parseCustomerID).Where("outlet_id = ?", parseOutletID).Find(&CartSale)

	var total int64
	config.DB.Raw("SELECT SUM(total_price) FROM cart_sales WHERE customer_id = ? and outlet_id = ?", parseCustomerID, parseOutletID).Scan(&total)

	// c.JSON(http.StatusOK, total)
	// return

	if total != 0 {
		data := models.TrxSale{
			CustomerID:   parseCustomerID,
			OutletID:     parseOutletID,
			TotalPayment: total,
		}
		if err := config.DB.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, "failed")
		} else {
			SaleID := strconv.FormatInt(int64(data.ID), 10)
			parseSaleID, _ := strconv.ParseInt(SaleID, 10, 64)

			for _, val := range CartSale {
				detail := models.TrxDetailSale{
					SaleID:     parseSaleID,
					OutletID:   val.OutletID,
					ProductID:  val.ProductID,
					Qty:        val.Qty,
					Price:      val.Price,
					TotalPrice: val.TotalPrice,
				}
				config.DB.Create(&detail)
			}
			config.DB.Delete(&CartSale)
			c.JSON(http.StatusOK, "checkout successfully")
		}
	}
}

func ShowTrxSale(c *gin.Context) {
	id := c.Params.ByName("id")
	var trxSale models.TrxSale
	err := config.DB.Preload("Detail").Preload("Outlet.Merchant").First(&trxSale, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not found")
	} else {
		c.JSON(http.StatusOK, trxSale)
	}
}

func UpdateTrxSale(c *gin.Context) {
	id := c.Params.ByName("id")
	var trxSale models.TrxSale
	data := config.DB.First(&trxSale, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, "data not found")
		return
	}

	input := models.TrxSale{
		Status: "paid",
	}

	if err := data.Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "payment successfully")
	}
}

func DeleteTrxSale(c *gin.Context) {
	id := c.Params.ByName("id")
	var trxSale models.TrxSale
	err := config.DB.First(&trxSale, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	if err := config.DB.Delete(&trxSale).Error; err != nil {
		c.JSON(http.StatusBadRequest, "failed")
	} else {
		c.JSON(http.StatusOK, "deleted data successfully")
	}
}

func ReportTrxSale(c *gin.Context) {
	startDate := c.PostForm("start_date")
	endDate := c.PostForm("end_date")

	var trxSale []models.TrxSale
	config.DB.Where("DATE(created_at) >= ?", startDate).Where("DATE(created_at) <= ?", endDate).Find(&trxSale)
	c.JSON(http.StatusOK, service.Response(trxSale, c, "", 0))
}

func ReportTrxSalePerProduct(c *gin.Context) {
	productID := c.Params.ByName("id")
	startDate := c.PostForm("start_date")
	endDate := c.PostForm("end_date")

	var TrxDetailSale []models.TrxDetailSale
	config.DB.Preload("Product").Where("product_id = ?", productID).Where("DATE(created_at) >= ?", startDate).Where("DATE(created_at) <= ?", endDate).Find(&TrxDetailSale)
	c.JSON(http.StatusOK, service.Response(TrxDetailSale, c, "", 0))
}
