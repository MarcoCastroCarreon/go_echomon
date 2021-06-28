package echoconfig

import (
	"net/http"

	routes "github.com/MarcoCastroCarreon/go_echomon/src/routes"
	"github.com/MarcoCastroCarreon/go_echomon/src/utils"
	"github.com/fatih/color"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer() {
	server := echo.New()
	server.Use(middleware.CORS())
	server.Use(middleware.Logger())
	server.GET("", func(echoCtx echo.Context) error {
		return echoCtx.String(http.StatusOK, "Server Started!!")
	})
	server.Validator = &utils.CustomValidator{Validator: validator.New()}

	server = routes.GetRoutes(server)
	color.Green("Initiating Server...")
	server.Logger.Fatal(server.Start(":3000"))
}
