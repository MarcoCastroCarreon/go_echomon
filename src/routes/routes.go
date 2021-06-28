package route

import (
	"github.com/MarcoCastroCarreon/go_echomon/src/handlers"
	"github.com/labstack/echo/v4"
)

func GetRoutes(server *echo.Echo) *echo.Echo {

	apiGroup := server.Group("/api")

	apiGroup.GET("/users", handlers.GetUsers)
	apiGroup.POST("/users", handlers.CreateUser)

	return server
}
