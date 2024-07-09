package services

import (
	"github.com/labstack/echo/v4"
)

type UserService interface {
    CreateUser(c echo.Context) error
    GetUser(c echo.Context) error
    DeleteUser(c echo.Context) error
}