package handlers

import (
	"net/http"

	"github.com/MarcoCastroCarreon/go_echomon/src/dtos"
	"github.com/MarcoCastroCarreon/go_echomon/src/models"
	"github.com/MarcoCastroCarreon/go_echomon/src/utils"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
)

func GetUsers(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello")
}

func CreateUser(ctx echo.Context) error {
	user := new(dtos.CreateUserRequest)

	if err := ctx.Bind(user); err != nil {
		return utils.GetResponse(ctx, http.StatusBadRequest, "An error has ocurred trying to parse your request")
	}

	if errorOnRequest := ctx.Validate(user); errorOnRequest != nil {
		return utils.GetResponse(ctx, http.StatusBadRequest, "An error has ocurred trying to parse your request")
	}

	hashedPassword, hashError := utils.HashPassword((user.Password))

	if hashError != nil {
		return utils.GetResponse(ctx, http.StatusConflict, "Something went wrong, please try again later")
	}

	newUser := &models.UserModel{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}

	err := mgm.Coll(newUser).Create(newUser)

	if err != nil {
		return utils.GetResponse(ctx, http.StatusInternalServerError, "Something went wrong, please try again later")
	}

	return utils.GetResponse(ctx, http.StatusCreated, "User Created Successfully")
}
