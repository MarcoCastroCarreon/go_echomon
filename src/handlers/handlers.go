package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(ctx echo.Context) error {
	fmt.Println("GET_USERS")
	return ctx.String(http.StatusOK, "Hello")
}
