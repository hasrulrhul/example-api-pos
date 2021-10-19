package main

import (
	"api-pos/app/models"
	"api-pos/config"
	"api-pos/route"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Merchant{})
	config.DB.AutoMigrate(&models.Outlet{})
	config.DB.AutoMigrate(&models.Product{})
	config.DB.AutoMigrate(&models.Customer{})
	config.DB.AutoMigrate(&models.ProductOutlet{})
	config.DB.AutoMigrate(&models.CartPurchase{})
	config.DB.AutoMigrate(&models.CartSale{})
	config.DB.AutoMigrate(&models.TrxPurchase{})
	config.DB.AutoMigrate(&models.TrxDetailPurchase{})
	config.DB.AutoMigrate(&models.TrxSale{})
	config.DB.AutoMigrate(&models.TrxDetailSale{})
}

func main() {

	defer config.CloseDatabaseConnection(config.DB)

	r := route.SetupRouter()
	// f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()

	// log.SetOutput(f)
	// log.Println(config.DB)
	r.Run(":" + os.Getenv("APP_PORT"))

}
