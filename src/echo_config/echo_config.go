package echoconfig

import (
	"net/http"

	routes "github.com/MarcoCastroCarreon/go_echomon/src/routes"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

func InitServer() {
	server := echo.New()
	server.GET("", func(echoCtx echo.Context) error {
		return echoCtx.String(http.StatusOK, "Server Started!!")
	})

	server = routes.GetRoutes(server)
	color.Green("Initiating Server...")
	server.Logger.Fatal(server.Start(":3000"))
}
