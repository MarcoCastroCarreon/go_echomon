package echoconfig

import (
	"fmt"
	"net/http"

	routes "github.com/MarcoCastroCarreon/go_echomon/src/routes"
	"github.com/labstack/echo/v4"
)

func InitServer() {
	fmt.Println("Initiating Server...")
	server := echo.New()
	server.GET("", func(echoCtx echo.Context) error {
		return echoCtx.String(http.StatusOK, "Server Started!!")
	})

	routes.GetRoutes(server)

	server.Logger.Fatal(server.Start(":3000"))
}
