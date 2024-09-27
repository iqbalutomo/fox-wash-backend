package router

import (
	"api_gateway/controllers"
	"api_gateway/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Echo(e *echo.Echo, uc controllers.UserController, wc controllers.WashStationController, oc controllers.OrderController) {
	users := e.Group("/users")
	{
		register := users.Group("/register")
		{
			register.POST("/user", uc.UserRegister)
			register.POST("/washer", uc.WasherRegister)
			register.POST("/admin", uc.AdminRegister)
		}

		users.GET("/verify/:user_id/:token", uc.VerifyUser)
		users.GET("/verified", func(c echo.Context) error {
			return c.Render(http.StatusOK, "verified.html", nil)
		})

		users.POST("/login", uc.Login)
		users.POST("/logout", uc.Logout, middlewares.Auth)
	}

	admin := e.Group("/admins")
	admin.Use(middlewares.Auth)
	{
		admin.GET("/washer-activation/:email", uc.WasherActivation)
	}

	washstations := e.Group("/washstations")
	washstations.Use(middlewares.Auth)
	{
		washstations.POST("/wash-package", wc.CreateWashPackage)
		washstations.GET("/wash-package/all", wc.GetAllWashPackages)
		washstations.GET("/wash-package/:id", wc.GetWashPackageByID)
		washstations.PUT("/wash-package/:id", wc.UpdateWashPackage)
		washstations.DELETE("/wash-package/:id", wc.DeleteWashPackage)

		washstations.POST("/detailing-package", wc.CreateDetailingPackage)
		washstations.GET("/detailing-package/all", wc.GetAllDetailingPackages)
		washstations.GET("/detailing-package/:id", wc.GetDetailingPackageByID)
		washstations.PUT("/detailing-package/:id", wc.UpdateDetailingPackage)
		washstations.DELETE("/detailing-package/:id", wc.DeleteDetailingPackage)
	}

	orders := e.Group("")
	orders.Use(middlewares.Auth)
	{
		users := orders.Group("/users")
		{
			users.GET("/orders", oc.GetUserAllOrders)
			users.POST("/orders", oc.CreateOrder)
		}
		washers := orders.Group("/washers")
		{
			washers.GET("/orders", oc.GetWasherAllOrders)
			washers.GET("/orders/:id", oc.WasherGetOrderByID)
			washers.GET("/orders/ongoing", oc.WasherGetCurrentOrder)
			washers.PUT("/orders/status/:id", oc.UpdateWasherOrderStatus)
		}
	}

	payments := e.Group("/payments")
	{
		payments.POST("/invoice", oc.UpdatePaymentStatus)
	}
}
