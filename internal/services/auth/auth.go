package services

import (
	"github.com/labstack/echo/v4"
)

type OrderService interface {
    AuthLogin(c echo.Context) error
}