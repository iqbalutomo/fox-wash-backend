package main

import (
	config "api_gateway/configs"
	"api_gateway/controllers"
	router "api_gateway/routers"
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

	userController := controllers.NewUserController(userClient)

	router.Echo(e, *userController)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
