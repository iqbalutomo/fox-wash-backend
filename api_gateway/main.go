package main

import (
	config "api_gateway/configs"
	"api_gateway/controllers"
	router "api_gateway/routers"
	"api_gateway/services"
	"api_gateway/utils"
	"html/template"
	"io"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

// @title Fox Wash
// @version 0.0.1
// @description Online motorcycle washing service built with microservices that integrates user, washer, and admin.

// @contact.name The developer
// @contact.email muhlisiqbalutomoo@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host fox-wash-production-910972720279.asia-southeast2.run.app
// @BasePath /
func main() {
	e := echo.New()
	e.Validator = &utils.CustomValidator{NewValidator: validator.New()}
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}
	e.Static("/templates", "templates/")
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger(), middleware.Recover(), middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	userClientConn, userClient := config.InitUserServiceClient()
	defer userClientConn.Close()

	washstationClientConn, washstationClient := config.InitWashStationServiceClient()
	defer washstationClientConn.Close()

	orderClientConn, orderClient := config.InitOrderServiceClient()
	defer orderClientConn.Close()

	mapService := services.NewMapService()

	userController := controllers.NewUserController(userClient)
	washstationController := controllers.NewWashStationController(washstationClient)

	orderController := controllers.NewOrderController(orderClient, userClient, mapService)

	router.Echo(e, *userController, *washstationController, *orderController)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
