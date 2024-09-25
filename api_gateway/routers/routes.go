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
	}

	orders := e.Group("")
	orders.Use(middlewares.Auth)
	{
		users := orders.Group("/users")
		{
			users.POST("/orders", oc.CreateOrder)
		}
	}

	payments := e.Group("/payments")
	{
		payments.POST("/invoice", oc.UpdatePaymentStatus)
	}
}
