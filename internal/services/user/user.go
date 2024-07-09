package services

import (
	"github.com/labstack/echo/v4"
)

type OrderService interface {
    CreateOrder(c echo.Context) error
    GetOrder(c echo.Context) error
    DeleteOrder(c echo.Context) error
	UpdateOrder(c echo.Context) error
}