package router

import (
	"api_gateway/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Echo(e *echo.Echo, uc controllers.UserController) {
	users := e.Group("/users")
	{
		users.POST("/register", uc.Register)
		users.GET("/verify/:user_id/:token", uc.VerifyUser)
		users.GET("/verified", func(c echo.Context) error {
			return c.Render(http.StatusOK, "verified.html", nil)
		})
		users.POST("/login", uc.Login)
	}
}
