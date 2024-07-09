package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"LastProject/internal/database"
	"LastProject/internal/localtime"
	"LastProject/internal/model"
)
type OrderLogic struct {
	DB *gorm.DB
}

// @Summary Create a new order
// @Description Create a new order in the system
// @Tags orders
// @Accept  json
// @Produce  json
// @Param   OrderId query int true "OrderId"
// @Param   Name query string true "Order name"
// @Success 201 {object} model.Order
// @Failure 400 "enter correct UserId and OrderId"
// @Failure 401 "Unauthorized"
// @Failure 404 "order was not created"
// @Router /orders [post]
// @Security BearerAuth
func (ol OrderLogic) CreateOrder(c echo.Context) error {
	name := c.QueryParam("Name")

	OrderId, err := strconv.Atoi(c.QueryParam("OrderId"))
	if err != nil {
		return echo.ErrBadRequest
	}

	VerifiedUser := c.Get("user").(*jwt.Token)
	claims := VerifiedUser.Claims.(jwt.MapClaims)
	idFloat := claims["UserId"].(float64)
	id := uint(idFloat)

	order := model.Order{Name: name, UserId: id, OrderId: uint(OrderId), CreatedAt: localtime.LocalTime(), Status: "Accepted", FinishedAt: time.Time{}}

	database.AddOrderToDataBase(ol.DB, &order)
	return c.JSON(http.StatusCreated, order)
}

// @Summary Get an order by ID
// @Description Get details of a order by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param   OrderId path int true "Order ID"
// @Success 200 {object} model.Order
// @Failure 401 "Unauthorized"
// @Failure 404 "order not found"
// @Router /orders/{OrderId} [get]
// @Security BearerAuth
func (ol OrderLogic) GetOrder(c echo.Context) error {
	id := c.Param("OrderId")

	var order model.Order
	
	if result := ol.DB.First(&order, id); result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "order not found")
	}

	return c.JSON(http.StatusOK, order)
}

// @Summary Delete an order by ID
// @Description Deleting order from db according to ID
// @Tags orders
// @Accept json
// @Param OrderId path int true "Order ID"
// @Success 200 "Order deleted"
// @Failure 401 "Unauthorized"
// @Failure 404 "Order not found"
// @Router /orders/{OrderId} [delete]
// @Security BearerAuth
func (ol OrderLogic) DeleteOrder(c echo.Context) error {
	id := c.Param("OrderId")
	if err := ol.DB.Delete(&model.Order{}, id).Error; err != nil {
        return err
    }

	return nil
}

// @Summary Updates status of order
// @Description Can update status of order by id
// @Tags orders
// @Accept  json
// @Produce  json
// @Param   OrderId path int true "Order ID"
// @Param   Status query string true "Status"
// @Success 200 {object} model.Order
// @Failure 401 "Unauthorized"
// @Failure 404 "order not found"
// @Failure 500 "internal server error"
// @Router /orders/{OrderId} [put]
// @Security BearerAuth
func (ol OrderLogic) UpdateOrder(c echo.Context) error {
    orderId := c.Param("OrderId")
	status := c.QueryParam("Status")
    var order model.Order

    if err := ol.DB.First(&order, orderId).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "order not found")
    }

	order.Status = status
	if strings.ToLower(status) == "finished" {
		order.FinishedAt = localtime.LocalTime()
	}else {
		order.FinishedAt = time.Time{}
	}

    if err := ol.DB.Save(&order).Error; err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, order)
}