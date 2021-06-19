package route

import (
	"github.com/MarcoCastroCarreon/go_echomon/src/handlers"
	"github.com/labstack/echo/v4"
)

func GetRoutes(server *echo.Echo) echo.Echo {
	server.Group("/api")
	{
		server.GET("/users", handlers.GetUsers)
	}

	return *server
}
