package utils

import (
	"github.com/MarcoCastroCarreon/go_echomon/src/dtos"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func getStatusCodeMessage(statusCode int) string {
	switch statusCode {
	case 200:
		return "Ok"
	case 201:
		return "Created"
	case 202:
		return "Accepted"
	case 204:
		return "No Content"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 404:
		return "Not Found"
	case 409:
		return "Conflict"
	case 500:
		return "Internal Server Error"
	case 502:
		return "Bad Gateway"
	default:
		return "No Status Code"
	}
}

func GetResponse(ctx echo.Context, statusCode int, message string) error {
	response := new(dtos.Response)
	response.Code = getStatusCodeMessage(statusCode)
	response.Message = message
	response.StatusCode = statusCode
	return ctx.JSON(statusCode, response)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
