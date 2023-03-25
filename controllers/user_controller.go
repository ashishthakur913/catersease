package controllers

import (
	"github.com/ashishthakur913/CaterEase/models"
	"github.com/ashishthakur913/CaterEase/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx echo.Context) error {
	var user models.User

	err := ctx.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	err = c.userService.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetHealth(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}
