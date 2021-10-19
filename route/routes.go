package route

import (
	"api-pos/app/controllers"
	"api-pos/middleware"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	name     = "name"
	username = "username"
	email    = "email"
	role     = "role"
)

func SetupRouter() *gin.Engine {
	// r := gin.Default()
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	r.MaxMultipartMemory = 8 << 20

	Route := r.Group("/api")
	{
		Route.GET("hello", controllers.Index)
		Route.POST("/login", controllers.Login)
		Route.POST("/register", controllers.Register)
		Route.POST("/logout", controllers.Logout)

		Route.GET("/ping", func(c *gin.Context) {
			c.String(200, "No Authorization JWT")
		})

		auth := Route.Group("/")
		auth.Use(middleware.Authz())
		auth.Use(AuthRequired)
		{
			auth.GET("/status", controllers.Status)

			user := auth.Group("/users")
			{
				user.GET("", controllers.IndexUser)
				user.POST("", controllers.CreateUser)
				user.GET("/:id", controllers.ShowUser)
				user.PUT("/:id", controllers.UpdateUser)
				user.DELETE("/:id", controllers.DeleteUser)
			}

			role := auth.Group("/merchants")
			{
				role.GET("", controllers.IndexMerchant)
				role.POST("", controllers.CreateMerchant)
				role.GET("/:id", controllers.ShowMerchant)
				role.PUT("/:id", controllers.UpdateMerchant)
				role.DELETE("/:id", controllers.DeleteMerchant)
			}

			menu := auth.Group("/outlets")
			{
				menu.GET("", controllers.IndexOutlet)
				menu.POST("", controllers.CreateOutlet)
				menu.GET("/:id", controllers.ShowOutlet)
				menu.PUT("/:id", controllers.UpdateOutlet)
				menu.DELETE("/:id", controllers.DeleteOutlet)
			}

			product := auth.Group("/products")
			{
				product.GET("", controllers.IndexProduct)
				product.POST("", controllers.CreateProduct)
				product.GET("/:id", controllers.ShowProduct)
				product.PUT("/:id", controllers.UpdateProduct)
				product.DELETE("/:id", controllers.DeleteProduct)
			}

			customer := auth.Group("/customers")
			{
				customer.GET("", controllers.IndexCustomer)
				customer.POST("", controllers.CreateCustomer)
				customer.GET("/:id", controllers.ShowCustomer)
				customer.PUT("/:id", controllers.UpdateCustomer)
				customer.DELETE("/:id", controllers.DeleteCustomer)
			}

			productoutlet := auth.Group("/product-outlets")
			{
				productoutlet.GET("", controllers.IndexProductOutlet)
				productoutlet.POST("", controllers.CreateProductOutlet)
				productoutlet.GET("/:id", controllers.ShowProductOutlet)
				productoutlet.PUT("/:id", controllers.UpdateProductOutlet)
				productoutlet.DELETE("/:id", controllers.DeleteProductOutlet)
			}

			cart := auth.Group("/cart-purchases")
			{
				cart.GET("", controllers.IndexCartPurchase)
				cart.POST("", controllers.CreateCartPurchase)
				cart.GET("/:id", controllers.ShowCartPurchase)
				cart.PUT("/:id", controllers.UpdateCartPurchase)
				cart.DELETE("/:id", controllers.DeleteCartPurchase)
			}

			trxpurchase := auth.Group("/transaction-purchases")
			{
				trxpurchase.GET("", controllers.IndexTrxPurchase)
				trxpurchase.POST("/checkout", controllers.CreateTrxPurchase)
				trxpurchase.GET("/:id", controllers.ShowTrxPurchase)
				trxpurchase.PUT("/payment/:id", controllers.UpdateTrxPurchase)
				trxpurchase.DELETE("/:id", controllers.DeleteTrxPurchase)
				// trxpurchase.GET("/report", controllers.ReportTrxPurchase)
				// trxpurchase.GET("/report/:producID", controllers.ReportTrxPurchasePerProduct)

			}

			carts := auth.Group("/cart-sales")
			{
				carts.GET("", controllers.IndexCartSale)
				carts.POST("", controllers.CreateCartSale)
				carts.GET("/:id", controllers.ShowCartSale)
				carts.PUT("/:id", controllers.UpdateCartSale)
				carts.DELETE("/:id", controllers.DeleteCartSale)
			}

			trxsale := auth.Group("/transaction-sales")
			{
				trxsale.GET("", controllers.IndexTrxSale)
				trxsale.POST("/checkout", controllers.CreateTrxSale)
				trxsale.GET("/:id", controllers.ShowTrxSale)
				trxsale.PUT("/payment/:id", controllers.UpdateTrxSale)
				trxsale.DELETE("/:id", controllers.DeleteTrxSale)
				// trxpurchase.GET("/report", controllers.ReportTrxSale)
				// trxpurchase.GET("/report/:producID", controllers.ReportTrxSalePerProduct)
			}

		}

	}

	return r
}

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(name)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}
