package controllers

import (
	"github.com/labstack/echo/v4"
)

type Handlers struct {
    UserController  *UserController
}

func (s *Handlers) CreateUser(ctx echo.Context) error {
    return s.UserController.CreateUser(ctx)
}

func (s *Handlers) GetHealth(ctx echo.Context) error {
    return s.UserController.GetHealth(ctx)
}
